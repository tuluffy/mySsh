package main

import (
	"fmt"
	"os"

	"os/user"

	"flag"

	"./core"
)

func main() {
	who, err := user.Current()
	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}
	prefix := who.HomeDir
	flag.Parse()
	app := core.App{ServerPath: prefix + "/.mySsh"}
	app.Exec()
}
