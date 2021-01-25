package main

import (
	"os"
	"testing"
	"io/ioutil"
	"encoding/json"
)

// TempFiles represents a structure for the dummy temp_files.json
type TestTempFiles struct {
	Filepaths []string
	Created string
}

// Dummy temp_files.json
var tempFilesJSON = `
	{
		"Filepaths": [
			"/home/user/Documents/example_dir",
			"/home/user/Documents/example_file.txt",
			"/home/user/Documents/example_subfolder",
			"/home/user/Documents/example_subfolder/example_subfile.txt",
		],
		"Created": "2021-01-25 13:49:46"
	}
`

// Test to check if there are files
// inside the directory inputed by the user
func TestIsDirEmpty(t *testing.T) {
	// Getting the working directory
	rootDir, err := os.Getwd()
	if err != nil {
		expected := true
		got := false
	
		if got != expected {
			t.Errorf("Expected: %v, got: %v", expected, got)
		}
	}

	// Scaning the directory
	files, err := ioutil.ReadDir(rootDir)
	
	if err != nil {
		expected := true
		got := false

		if got != expected {
			t.Errorf("Expected: %v, got: %v, files: %v", expected, got, files)
		}		
	}
}

// Test to check the dummy temp_files.json file
func TestTempFilesJSON(t *testing.T) {
	var testTempFiles TestTempFiles
	json.Unmarshal([]byte(tempFilesJSON), &testTempFiles)

	expected := testTempFiles.Filepaths
	got := testTempFiles.Filepaths
	if got != nil {
		t.Errorf("Expected: %v, got: %v", expected, got)
	}
}
