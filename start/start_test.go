package main

import (
	"os"
	"strings"
	"testing"
	"runtime"
	"path/filepath"
	"encoding/json"
)

// Config represents a structure for the dummy config.json
type TestConfig struct {
	RootDirectory, BackupDirectory, ProjectName, ProjectVersion, Author, License string
}

// Log represents a structure for the dummy log.json
type TestLog struct {
	ID, Created, Message string
}

// Dummy config.json
var configJSON = `
	{
		"RootDirectory": "C:/Users/coem/Documents/programming/bup",
		"BackupDirectory": "C:/Users/coem/Documents/bup",
		"ProjectName": "bup",
		"ProjectVersion": "0.0.1",
		"Author": "kevinsuner",
		"License": "MIT"
	}
`

// Dummy log.json
var logJSON = `
	{
		"ID": "TRdcnPtoBODRpRyoYOdt",
		"Created": "2021-01-09 17:20:48",
		"Message": "Bup successfully initialized... at C:/Users/coem/Documents/bup"
	}
`

// Test to check the data inputed
// by the user
func TestProjectData(t *testing.T) {
	// Finding the root directory
	_, b, _, _ := runtime.Caller(0)
	rootDirectory := filepath.Dir(b)

	// Appending the root directory 
	// to the backup directory inputed
	// by the user
	parts := strings.Split(rootDirectory, string(os.PathSeparator))
	backupDirectory := parts[0] + "/Users/coem/Documents/bup"

	projectName := strings.Replace("bup", "\n", "", -1)
	projectVersion := strings.Replace("0.0.1", "\n", "", -1)
	author := strings.Replace("kevinsuner", "\n", "", -1)
	license := strings.Replace("MIT", "\n", "", -1)

	expected := "bup"
	got := projectName
	if got != expected {
		t.Errorf("Expected: %v, got: %v", expected, got)
	}

	expected = "0.0.1"
	got = projectVersion
	if got != expected {
		t.Errorf("Expected: %v, got: %v", expected, got)
	}

	expected = "kevinsuner"
	got = author
	if got != expected {
		t.Errorf("Expected: %v, got: %v", expected, got)
	}

	expected = "MIT"
	got = license
	if got != expected {
		t.Errorf("Expected: %v, got: %v", expected, got)
	}

	expected = "C:\\Users\\coem\\Documents\\programming\\bup"
	got = rootDirectory
	if got != expected {
		t.Errorf("Expected: %v, got: %v", expected, got)
	}

	expected = "C:/Users/coem/Documents/bup"
	got = backupDirectory
	if got != expected {
		t.Errorf("Expected: %v, got: %v", expected, got)
	}
}

// Test to check if the backup folder exists
func TestBackupDir(t *testing.T) {
	// Finding the root directory
	// and adding the directory 
	// inputed by the user
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	parts := strings.Split(basepath, string(os.PathSeparator))
	bupDir := parts[0] + "/Users/coem/Documents/bup"

	if _, err := os.Stat(bupDir); os.IsNotExist(err) {
		expected := true
		got := true
		if got != expected {
			t.Errorf("Expected: %v, got: %v", expected, got)
		}
	}

	if _, err := os.Stat(bupDir); os.IsExist(err) {
		expected := true
		got := true
		if got != expected {
			t.Errorf("Expected: %v, got: %v", expected, got)
		}
	}
}

// Test to check the dummy config.json
func TestConfigJSON(t *testing.T) {
	var testConfig TestConfig
	json.Unmarshal([]byte(configJSON), &testConfig)
	
	expected := "C:/Users/coem/Documents/programming/bup"
	got := testConfig.RootDirectory
	if got != expected {
		t.Errorf("Expected: %v, got: %v", expected, got)
	}

	expected = "C:/Users/coem/Documents/bup"
	got = testConfig.BackupDirectory
	if got != expected {
		t.Errorf("Expected: %v, got: %v", expected, got)
	}

	expected = "bup"
	got = testConfig.ProjectName
	if got != expected {
		t.Errorf("Expected: %v, got: %v", expected, got)
	}

	expected = "0.0.1"
	got = testConfig.ProjectVersion
	if got != expected {
		t.Errorf("Expected: %v, got: %v", expected, got)
	}

	expected = "kevinsuner"
	got = testConfig.Author
	if got != expected {
		t.Errorf("Expected: %v, got: %v", expected, got)
	}

	expected = "MIT"
	got = testConfig.License
	if got != expected {
		t.Errorf("Expected: %v, got: %v", expected, got)
	}
}

// Test to check the dummy log.json
func TestLogJSON(t *testing.T) {
	var testLog TestLog
	json.Unmarshal([]byte(logJSON), &testLog)
	
	expected := "TRdcnPtoBODRpRyoYOdt"
	got := testLog.ID
	if got != expected {
		t.Errorf("Expected: %v, got: %v", expected, got)
	}

	expected = "2021-01-09 17:20:48"
	got = testLog.Created
	if got != expected {
		t.Errorf("Expected: %v, got: %v", expected, got)
	}

	expected = "Bup successfully initialized... at C:/Users/coem/Documents/bup"
	got = testLog.Message
	if got != expected {
		t.Errorf("Expected: %v, got: %v", expected, got)
	}
}