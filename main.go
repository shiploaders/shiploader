package main

import (
	"fmt"
	"log"
	"os"
	command "shiploader/cmd"
)

func main() {
	fmt.Println("Shiploader CLI v0.01")
	app := command.Generate()
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
