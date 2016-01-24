package main

import (
	"fmt"
	"os"

	"github.com/integralist/Buskin/homebrew"
	"github.com/integralist/buskin/config"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No config file path provided")
		os.Exit(1)
	}

	configPath := os.Args[1]
	configFile := config.Get(configPath)

	c, err := config.Parse(configFile)
	if err != nil {
		fmt.Printf("Problem parsing the yaml: %s\n", err.Error())
	}

	fmt.Printf("%+v\n\n", *c)

	homebrew.Setup()

	// go run buskin.go ./config.yaml
	// getting back exit status 1 for running ruby script?
}

/*
github ssh key:
  - check directory ~/.ssh exists
	- generate key with random passcode (print code)
	- run ssh agent
	- print public key to stdout with notice to user to add to GitHub gui
	  - be nice if this could be done programmatically with an API
	- set global config for username/email

homebrew:
  - check for homebrew
	- install homebrew
	- install packages if missing (report those already present)
	- put check in for neovim

git:
  - clone dotfiles repo
	- move dotfiles into specified directory
	- clone vundle repo and run vim plugin installation

shell:
  - change terminal theme programmatically?

golang:
  - go get all the specified packages

Indicate to user other software they might need:
  - Java JDK
	- Vagrant
	- VirtualBox
	- Docker
*/
