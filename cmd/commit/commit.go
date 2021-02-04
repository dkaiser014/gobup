package commit

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"bup/pkg/existsfile"
	"bup/pkg/getbupdir"
	"bup/pkg/getrootdir"
	"bup/pkg/randomstring"
)

// Log represents a structure for the log.json file.
type Log struct {
	ID        string `json:"ID"`
	Message   string `json:"Message"`
	CreatedAt string `json:"CreatedAt"`
}

// Commit creates a commit within the log.json file
// containing a message inputed by the user.
func Commit() string {
	if existsfile.ExistsFile(getrootdir.GetRootDir() + string(os.PathSeparator) + "temp_files.json") {
		fmt.Print("> Commit message: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		message := scanner.Text()

		if len(message) <= 0 {
			log.Fatal("Error: please input a message in order to create a new commit...")
		}

		jsonFile, err := os.Open(getrootdir.GetRootDir() + string(os.PathSeparator) + "log.json")

		if err != nil {
			log.Fatal("Error:", err)
		}

		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)

		// Appending to log.json
		var logs []Log
		err = json.Unmarshal([]byte(byteValue), &logs)
		logs = append(logs, Log{ID: randomstring.RandomString(), Message: message, CreatedAt: time.Now().String()})

		file, _ := json.MarshalIndent(logs, "", " ")
		_ = ioutil.WriteFile(getrootdir.GetRootDir()+string(os.PathSeparator)+"log.json", file, 0644)
		_ = ioutil.WriteFile(getbupdir.GetBupDir()+string(os.PathSeparator)+"log.json", file, 0644)

		fmt.Printf("> New commit ID:%v successfully added to log.json", randomstring.RandomString())
	}

	os.Exit(3)
	return ""
}
