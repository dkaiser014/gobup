package main

import (
  "fmt"
  "flag"
  "os"
  "log"
  "encoding/json"
  "io/ioutil"
)

/*
*
* - TODO
* > Take the name of the new branch with a flag
* > Check if the folder already exists, if not, create it, if yes, exit
* > Print a success message and exit
*
*/

// Config represents a structure for the config.json file
type Config struct {
  RootDirectory string `json:"RootDirectory"`
  BackupDirectory string `json:"BackupDirectory"`
  ProjectName string `json:"ProjectName"`
  ProjectVersion string `json:"ProjectVersion"`
  Author string `json:"Author"`
  License string `json:"License"`
}

// GetBupDir returns the directory where the backup
// is stored
func GetBupDir() string {
  jsonFile, err := os.Open("config.json")
  if err != nil {
    log.Fatal(err)
  }
  defer jsonFile.Close()

  byteValue, _ := ioutil.ReadAll(jsonFile)
  var config Config
  json.Unmarshal(byteValue, &config)

  bupDir := config.BackupDirectory
  return bupDir
}

func CreateBranch() string {
  // Ask the user for a folder name using the "-f" flag
  name := flag.String("f", "", "Name of the folder to be created")
  flag.Parse()

  if len(*name) <=0 {
    log.Fatal("Error: please input a folder name in order to create a new folder...")
  } 

  // Appending the name of the new folder to bupDir
  newDir := GetBupDir() + string(os.PathSeparator) + *name
  fmt.Println(newDir)

  // Checking if folder exists, if yes, exit, if not, create it
  if _, err := os.Stat(newDir); os.IsExist(err) {
    log.Fatal(err)
  }

  if _, err := os.Stat(newDir); os.IsNotExist(err) {
    // Create the directory
    err := os.Mkdir(newDir, 0755)
    if err != nil {
      log.Fatal(err)
    }
  }

  fmt.Println("Successfully created a new branch in", newDir)

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
  CreateBranch()
}
