package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"storage" // For App-engine compile
	//_ "../storage" // For Editor
	"errors"
)

func GetKey(db storage.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key, err, code := processGet(r)
		if err != nil {
			http.Error(w, err.Error(), code)
			return
		}
		val, err := db.Get(key)
		if err == storage.ErrNotFound {
			http.Error(w, "not found", http.StatusNotFound)
			return
		} else if err != nil {
			http.Error(w, fmt.Sprintf("error from db: %s", err), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(val)
	}
}

func SetKey(db storage.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut && r.Method != http.MethodPost {
			http.Error(w, "error unsupported request method use PUT / POST", http.StatusBadRequest)
			return
		}
		var key string
		var err error
		var val []byte
		switch r.Method {
		case http.MethodPut:
			key = r.URL.Query().Get("key")
			if key == "" {
				http.Error(w, "Missing key Name in query string", http.StatusBadRequest)
				return
			}
			defer r.Body.Close()
			val, err = ioutil.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "error reading PUT body", http.StatusBadRequest)
				return
			}
		case http.MethodPost:
			if r.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
				http.Error(w, "Content type missing or incorrect only support form-urlencode", http.StatusNotAcceptable)
				return
			}
			key = r.PostFormValue("key")
			if key == "" {
				http.Error(w, "Missing key Name in POST body", http.StatusBadRequest)
				return
			}
			val = []byte(r.PostFormValue("value"))
			if len(val) == 0 {
				http.Error(w, "Missing value in POST body", http.StatusBadRequest)
				return
			}
		}

		if err := db.Set(key, val); err != nil {
			http.Error(w, "Error in setting DB", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func processGet(r *http.Request) (string, error, int) {
	if r.Method != http.MethodGet {
		return "", errors.New("error unsupported request method use GET"), http.StatusBadRequest
	}
	key := r.URL.Query().Get("key")
	if key == "" {
		return "", errors.New("Missing key Name in query string"), http.StatusBadRequest
	}
	return key, nil, http.StatusOK
}
