package main

import (
	"github.com/satori/go.uuid"
	"io/ioutil"
)

// File Name for the Fixed reference generation
const fixedUUIDFileName = ".uuid-fix"

// Directly generate and give back the UUID string
func GetUUID() string {
	return uuid.NewV4().String()
}

// Directly generate and give back the UUID RAW Byte Array
func GetUUIDbytes() []byte {
	return uuid.NewV4().Bytes()
}

func DerivedUUID(gen string) (string, error) {
	// Check if file exists and read it
	fixeduid, err := ioutil.ReadFile(fixedUUIDFileName)
	if err != nil {
		// If fine does not exist then generate an ID and Save it file
		fixeduid = []byte(uuid.NewV4().String())
		err = ioutil.WriteFile(fixedUUIDFileName, fixeduid, 0644)
	}

	// Generate the new UUID from the previously fixed UID and generator string
	tempid, _ := uuid.FromString(string(fixeduid))
	uid := uuid.NewV5(tempid, gen)
	return uid.String(), err
}
