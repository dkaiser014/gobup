package readjson

import (
	"io/ioutil"
	"log"
	"os"

	"bup/pkg/getrootdir"
)

// ReadJSON returns the content of a .json file
func ReadJSON(file string) []byte {
	jsonFile, err := os.Open(getrootdir.GetRootDir() + string(os.PathSeparator) + "log.json")

	if err != nil {
		log.Fatal("Error:", err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	return byteValue
}
