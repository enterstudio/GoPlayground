package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/boseji/goboseji"
	"github.com/satori/go.uuid"
	"github.com/valyala/fasthttp"
	"time"
)

type edata struct {
	Id    uuid.UUID `json:"id"`
	D     string    `json:"value,omitempty"`
	T     time.Time `json:"time"`
	Valid bool      `json:"-"`
}

var dstore []edata

const sizeDstore = 10

func retError(ctx *fasthttp.RequestCtx, er string, status int) {
	ctx.Logger().Printf("%s", er)
	ctx.Error(er, status)
}

func dataEpController(ctx *fasthttp.RequestCtx) {

	contentType := string(ctx.Request.Header.ContentType())

	switch {

	case ctx.IsGet():

		// Load the Small-ify and JSON options
		small := ctx.QueryArgs().Has("small") || ctx.QueryArgs().Has("Small")
		jsonEn := ctx.QueryArgs().Has("json") || ctx.QueryArgs().Has("JSON")

		// Log the Request
		ctx.Logger().Printf("Correct Request Small=%v JSON=%v", small, jsonEn)

		// Fall out if there is no Data Init
		if len(dstore) == 0 {
			ctx.Error("Empty", fasthttp.StatusNoContent)
			break
		}

		outbuf := fasthttp.AcquireByteBuffer()

		// process the D-Store
		for i := 0; i < len(dstore); i++ {
			// If the entry is valid
			if dstore[i].Valid {

				if !jsonEn {
					// Non-Json Process
					outbuf.WriteString(fmt.Sprintf("[%d]\n", i+1))
					if !small {

						outbuf.WriteString(dstore[i].Id.String() + "\n")
						outbuf.WriteString(dstore[i].T.Format("02-01-2006 15:04:05.000 MST") + "\n")
					}
					outbuf.WriteString(dstore[i].D + "\n")
				} else {
					// Process JSON
					bj, _ := json.Marshal(dstore[i])
					outbuf.B = append(outbuf.B, bj...)
					outbuf.WriteString(",") // Add a Coma For Tail end processing
				}

			}
		}

		if jsonEn {
			// Eliminate the Last Coma in JSON
			outbuf.B = bytes.TrimRight(outbuf.B, ",")
			// Use brackets only when multiple elements are available in JSON array
			if bytes.Count(outbuf.B, []byte("{")) > 1 {
				outbuf.B = bytes.Join([][]byte{[]byte("["), outbuf.B, []byte("]")}, []byte{})
			}
		}

		// Finally Print the Output
		ctx.Write(outbuf.B)
		// Release the Buffer
		fasthttp.ReleaseByteBuffer(outbuf)
		break

	case ctx.IsPost():
		if contentType != "application/x-www-form-urlencoded" {
			retError(ctx, "Error Unknown Content Type "+contentType, fasthttp.StatusBadRequest)
			break
		}

		// Create D Store if it does not Exists
		if len(dstore) == 0 {
			dstore = make([]edata, sizeDstore)
		}

		// Find the First Empty Location
		index := 0
		for ; index < len(dstore); index++ {
			if !dstore[index].Valid {
				break
			}
		}

		// If list full create an Empty location
		if index == len(dstore) {
			for i := 1; i < len(dstore); i++ {
				dstore[i-1] = dstore[i]
			}
			dstore[len(dstore)-1].Valid = false
			index = len(dstore) - 1
		}

		// Insert Data
		dstore[index].D = string(ctx.FormValue("data"))
		dstore[index].Id = uuid.NewV4()
		dstore[index].Valid = true
		dstore[index].T = goboseji.TimeIST(time.Now())
		ctx.WriteString(fmt.Sprintf("%d", index))
		ctx.Logger().Printf("Correct Request index=%d", index)
		break

	default:
		retError(ctx, "Error Unknown Protocol", fasthttp.StatusBadRequest)
		break
	}
}
