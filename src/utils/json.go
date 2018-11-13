package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func DecodeJson(body io.Reader, target interface{}) error {
	val, err := ioutil.ReadAll(body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%+v\n", val)
	decodeErr := json.NewDecoder(body).Decode(&target)

	return decodeErr
}

func PrintObject(object interface{}) {
	jsonData, err := json.MarshalIndent(object, "", "\t")
	if err != nil {
		log.Fatalln(err)
	}

	os.Stdout.Write(jsonData)
}
