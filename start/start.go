package main

import (
	"fmt"
	"os"
	"strings"
	"log"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"time"
)

/*
*
* > TODO
*
*/

// Config represents a structure for a .json file 
type Config struct {
	RootDirectory, BackupDirectory, ProjectName, ProjectVersion, Author, License string
}

// Log represents a structure for a .json file
type Log struct {
	ID, Created, Message string
}

// RandomString generates a random string both with
// Lowercase and Uppercase characters to be later used
// in the ID inside log.json
func RandomString() string {
	rand.Seed(time.Now().Unix())

	// Lowercase and Uppercase
	var output strings.Builder
	charSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	length := 20
	for i := 0; i < length; i++ {
		random := rand.Intn(len(charSet))
		randomChar := charSet[random]
		output.WriteString(string(randomChar))
	}

	return output.String()
}

// Start initializes bup in a specific folder
func Start() string {
	// Getting the project data
	var bupDir string
	fmt.Print("> Backup destination (required): ")
	_, err := fmt.Scanln(&bupDir)
	if err != nil {
		log.Fatal(err)
	}

	var projectName string
	fmt.Print("> Project name (required): ")
	_, err = fmt.Scanln(&projectName)
	if err != nil {
		log.Fatal(err)
	}

	var projectVersion string = "0.0.1"
	fmt.Print("> Project version (0.0.1) by default: ")
	fmt.Scanln(&projectVersion)

	var author string = "John Doe"
	fmt.Print("> Author (John Doe) by default: ")
	fmt.Scanln(&author)

	var license string = "MIT"
	fmt.Print("> License (MIT) by default: ")
	fmt.Scanln(&license)

	// Finding the root directory
	rootDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Spliting the directory into parts
	parts := strings.Split(rootDir, string(os.PathSeparator))

	// And adding the directory
	// inputed by the user
	bupDir = parts[0] + strings.Replace(bupDir, "/", string(os.PathSeparator), -1)

	// Checking if folder exists, if yes, exit, if not, create it
	if _, err := os.Stat(bupDir); os.IsNotExist(err) {
		// Creating the directory
		err := os.Mkdir(bupDir, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

	if _, err := os.Stat(parts[0] + bupDir); os.IsExist(err) {
		// Stopping execution
		log.Fatal(err)
	}

	// Saving data to a .json file in the old and new folder
	config := Config {
		RootDirectory: rootDir,
		BackupDirectory: bupDir,
		ProjectName: projectName,
		ProjectVersion: projectVersion,
		Author: author,
		License: license,
	}

	file, _ := json.MarshalIndent(config, "", " ")
	_ = ioutil.WriteFile(rootDir + string(os.PathSeparator) + "config.json", file, 0644)

	file, _ = json.MarshalIndent(config, "", " ")
	_ = ioutil.WriteFile(bupDir + string(os.PathSeparator) + "config.json", file, 0644)

	// Initializing bup and creating a .json log file in the old and new folder
	log := Log {
		ID: RandomString(),
		Created: time.Now().String(),
		Message: "Bup successfully initialized... at " + bupDir,
	}

	file, _ = json.MarshalIndent(log, "", " ")
	_ = ioutil.WriteFile(rootDir + string(os.PathSeparator) + "log.json", file, 0644)

	file, _ = json.MarshalIndent(log, "", " ")
	_ = ioutil.WriteFile(bupDir + string(os.PathSeparator) + "log.json", file, 0644)

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
	Start()
}
