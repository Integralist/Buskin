package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Structure contains the provided configuration
type Structure struct {
	Github struct {
		Username string `yaml:"username"`
		Email    string `yaml:"email"`
	} `yaml:"github"`

	Shell struct {
		Type string `yaml:"type"`
	} `yaml:"shell"`

	Dotfiles struct {
		Path string `yaml:"path"`
	} `yaml:"dotfiles"`

	Homebrew []string `yaml:"homebrew"`

	Gopackages []string `yaml:"go"`
}

// Get retrieves the specified configuration file
func Get(configPath string) []byte {
	config, err := ioutil.ReadFile(configPath)

	if err != nil {
		fmt.Printf("Problem reading the config file: %s\n", err)
		panic(err)
	}

	return config
}

// Parse unmarshals the yaml content into a struct
func Parse(configFile []byte) (*Structure, error) {
	var c Structure

	err := yaml.Unmarshal(configFile, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
