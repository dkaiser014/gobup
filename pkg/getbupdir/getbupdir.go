package getbupdir

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"bup/pkg/getrootdir"
)

// Config represents a structure for the config.json file
type Config struct {
	ID              string `json:"ID"`
	IsInit          string `json:"IsInit"`
	CreatedAt       string `json:"CreatedAt"`
	BackupDirectory string `json:"BackupDirectory"`
	RootDirectory   string `json:"RootDirectory"`
	ProjectName     string `json:"ProjectName"`
	ProjectVersion  string `json:"ProjectVersion"`
	Author          string `json:"Author"`
	License         string `json:"License"`
}

// GetBupDir returns the backup_directory
// from the config.json file
func GetBupDir() string {
	jsonFile, err := os.Open(getrootdir.GetRootDir() + string(os.PathSeparator) + "config.json")

	if err != nil {
		log.Fatal("Error:", err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var config Config
	json.Unmarshal(byteValue, &config)

	bupDir := config.BackupDirectory
	return bupDir
}
