package main

import (
  "fmt"
  "flag"
  "os"
  "log"
  "strings"
  "encoding/json"
  "io/ioutil"
  "math/rand"
  "time"
)

// Config represents a structure for the config.json file
type Config struct {
  RootDirectory string `json:"RootDirectory"`
  BackupDirectory string `json:"BackupDirectory"`
  ProjectName string `json:"ProjectName"`
  ProjectVersion string `json:"ProjectVersion"`
  Author string `json:"Author"`
  License string `json:"License"`
}

// Log represents a structure for the log.json file
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

// Commit creates a commit within the log.json file
// containing a message inputed by the user
func Commit() string {
  // Ask the user for a message using the "-m" flag
  message := flag.String("m", "", "Message for the commit inside log.json")
  flag.Parse()

  if len(*message) <= 0 {
    log.Fatal("Error: please input a message in order to create a new commit...")
  }

  // Getting the old log.json file
  jsonFile, err := os.Open("log.json")
  if err != nil {
    log.Fatal(err)
  }
  defer jsonFile.Close()

  byteValue, _ := ioutil.ReadAll(jsonFile) 

  // Appending the new commit to the log.json file
  var logs []Log
  err = json.Unmarshal([]byte(byteValue), &logs)
  logs = append(logs, Log{ID: RandomString(), Created: time.Now().String(), Message: *message})

  // Finding the root and the backup directory
	rootDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
 
  // Save the log.json file
  file, _ := json.MarshalIndent(logs, "", " ")
	_ = ioutil.WriteFile(rootDir + string(os.PathSeparator) + "log.json", file, 0644)

	file, _ = json.MarshalIndent(logs, "", " ")
	_ = ioutil.WriteFile(GetBupDir() + string(os.PathSeparator) + "log.json", file, 0644)

  fmt.Println(*message)

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
  Commit()
}
