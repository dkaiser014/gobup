package main

import (
  "fmt"
  "flag"
  //"os"
  "log"
  "strings"
  //"path/filepath"
  //"encoding/json"
  //"io/ioutil"
  "math/rand"
  "time"
)

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

// Commit creates a commit within the log.json file
// containing a message inputed by the user
func Commit(message string) string {
  fmt.Println(message)
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

  // Ask the user for a message using the "-m" flag
  message := flag.String("m", "", "Message for the commit inside log.json")
  flag.Parse()

  if len(*message) <= 0 {
    log.Fatal("Error: please input a message in order to create a new commit...")
  }

	Commit(*message)
}
