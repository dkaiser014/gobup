package start

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"bup/pkg/existsdir"
	"bup/pkg/getrootdir"
	"bup/pkg/randomstring"
)

// Config represents a structure for the config.json file.
type Config struct {
	ID              string `json:"ID"`
	IsInit          bool   `json:"IsInit"`
	CreatedAt       string `json:"CreatedAt"`
	BackupDirectory string `json:"BackupDirectory"`
	RootDirectory   string `json:"RootDirectory"`
	ProjectName     string `json:"ProjectName"`
	ProjectVersion  string `json:"ProjectVersion"`
	Author          string `json:"Author"`
	License         string `json:"License"`
}

// Start asks the user for information about the project
// and initializes the utility.
func Start() string {
	var bupDir, projectName, projectVersion, author, license string

	fmt.Print("> Backup directory: ")
	_, err := fmt.Scanln(&bupDir)
	if err != nil {
		log.Fatal("Error:", err)
	}

	fmt.Print("> Project name: ")
	_, err = fmt.Scanln(&projectName)
	if err != nil {
		log.Fatal("Error:", err)
	}

	fmt.Print("> Project version: ")
	_, err = fmt.Scanln(&projectVersion)
	if err != nil {
		log.Fatal("Error:", err)
	}

	fmt.Print("> Author: ")
	_, err = fmt.Scanln(&author)
	if err != nil {
		log.Fatal("Error:", err)
	}

	fmt.Print("> License: ")
	_, err = fmt.Scanln(&license)
	if err != nil {
		log.Fatal("Error:", err)
	}

	bupDir = strings.Replace(bupDir, "/", string(os.PathSeparator), -1)
	if existsdir.ExistsDir(bupDir) == false {
		config := Config{
			ID:              randomstring.RandomString(),
			IsInit:          true,
			CreatedAt:       time.Now().String(),
			BackupDirectory: bupDir,
			RootDirectory:   getrootdir.GetRootDir(),
			ProjectName:     projectName,
			ProjectVersion:  projectVersion,
			Author:          author,
			License:         license,
		}

		file, _ := json.MarshalIndent(config, "", " ")
		_ = ioutil.WriteFile(getrootdir.GetRootDir()+string(os.PathSeparator)+"config.json", file, 0644)
		_ = ioutil.WriteFile(bupDir+string(os.PathSeparator)+"config.json", file, 0644)

		_ = ioutil.WriteFile(getrootdir.GetRootDir()+string(os.PathSeparator)+"log.json", []byte(`{"Message": "Run ~bup commit~ to initialize log.json"}`), 0644)
		_ = ioutil.WriteFile(bupDir+string(os.PathSeparator)+"log.json", []byte(`{"Message": "Run ~bup commit~ to initialize log.json"}`), 0644)
	}

	return ""
}
