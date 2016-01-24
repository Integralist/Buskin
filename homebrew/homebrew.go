package homebrew

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
)

// Setup downloads Homebrew and runs it
func Setup() {
	resp, err := http.Get("https://raw.githubusercontent.com/Homebrew/install/master/install")
	if err != nil {
		fmt.Printf("Problem downloading Homebrew Ruby script: %s\n", err.Error())
	}

	defer resp.Body.Close()

	homebrew, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Problem reading Homebrew Ruby script: %s\n", err.Error())
	}

	homebrewPath := "/tmp/homebrew.rb"

	err = ioutil.WriteFile(homebrewPath, homebrew, 0777)
	if err != nil {
		fmt.Printf("Problem creating Homebrew Ruby script: %s\n", err.Error())
	}

	cmd := exec.Command("ruby", homebrewPath)
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Problem running Homebrew Ruby script: %s\n", err.Error())
		fmt.Println("Check that Homebrew isn't already installed")
	}
	// fmt.Printf("\n\n%q\n", out.String())
}
