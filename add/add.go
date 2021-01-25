package main

import (
	"fmt"
	"os"
	"log"
	"path/filepath"
	"encoding/json"
	"io/ioutil"
	"time"
)

/*
*
* TODO
* - Take console arguments as input and check them
* - If the args equal "." store all the filepaths within the folder in a .json file
* - If the args equal to "test.txt" verify that the file exists and store the filepath
* -- in a .json file
*
*/

// Files represents a structure for a .json file
type Files struct {
	Filepaths []string
	Created string
}

// FindRootDir finds the working directory
// where the program is being run
func FindRootDir() string {
	rootDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return rootDir
}

// WalkDir lists all files and directories using
// the path provided by the user
func WalkDir(rootDir string) ([]string, error) {
	var files []string
	err := filepath.Walk(FindRootDir(), func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	return files, err
}

// Add adds all the filepath(s) within 
// the folder to a temporal .json
func Add() string {
	// Read all the file(s) within the directory
	// and sub-directories, if it has
	files, err := WalkDir(FindRootDir())
	if err != nil {
		log.Fatal(err)
	}

	var filePaths []string
	for _, file := range files {
		// List all the files that are going 
		// to get saved to the .json file
		fmt.Printf("Added: %v \n", file)

		// Save the details of the file(s)
		// in a temporal .json file
		filePaths = append(filePaths, []string{file}...)
		files := Files {
			Filepaths: filePaths,
			Created: time.Now().String(),
		}

		file, _ := json.MarshalIndent(files, "", " ")
		_ = ioutil.WriteFile(FindRootDir() + string(os.PathSeparator) + "temp_files.json", file, 0644)
	}

	// Terminate execution
	os.Exit(3)

	return ""
}

func main() {
	fmt.Println(`
________ _____  __________ 
___  __ )__  / / /___  __ \
__  __  |_  / / / __  /_/ /
_  /_/ / / /_/ /  _  ____/ 
/_____/  \____/   /_/     
	`)
	Add()
}
