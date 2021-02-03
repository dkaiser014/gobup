package add

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"bup/pkg/existsfile"
	"bup/pkg/getrootdir"
	"bup/pkg/walkdir"
)

// Files represents a structure for the temp_files.json file.
type Files struct {
	Filepaths []string
}

// Add adds all the filepath(s) within
// the folder to a temporal json file.
func Add() string {
	if existsfile.ExistsFile(getrootdir.GetRootDir()+string(os.PathSeparator)+"config.json") == true {
		files, err := walkdir.WalkDir(getrootdir.GetRootDir())
		if err != nil {
			log.Fatal("Error:", err)
		}

		var filepaths []string
		for _, file := range files {
			fmt.Printf("Added: %v \n", file)

			filepaths = append(filepaths, []string{file}...)
			files := Files{
				Filepaths: filepaths,
			}

			file, _ := json.MarshalIndent(files, "", " ")
			_ = ioutil.WriteFile(getrootdir.GetRootDir()+string(os.PathSeparator)+"temp_files.json", file, 0644)
		}
	}

	os.Exit(3)
	return ""
}
