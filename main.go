package main

import (
	"fmt"
	"os"
	"strings"

	"bup/cmd/add"
	"bup/cmd/start"
)

func main() {
	fmt.Println(`
________ _____  __________ 
___  __ )__  / / /___  __ \
__  __  |_  / / / __  /_/ /
_  /_/ / / /_/ /  _  ____/ 
/_____/  \____/   /_/     
	`)

	arg := string(os.Args[1])

	if strings.Compare(arg, "start") == 0 {
		start.Start()
	}

	if strings.Compare(arg, "add") == 0 {
		add.Add()
	}
}
