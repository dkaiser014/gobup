package main

import (
	"fmt"
	"os"
	"strings"

	"bup/cmd/commit"
	"bup/cmd/push"
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

	if strings.Compare(arg, "commit") == 0 {
		commit.Commit()
	}

	if strings.Compare(arg, "push") == 0 {
		push.Push()
	}
}
