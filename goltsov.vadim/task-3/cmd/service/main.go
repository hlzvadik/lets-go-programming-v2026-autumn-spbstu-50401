package main

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func main() {
	configPath := flag.String("config", "config.yaml", "Path to config file")
	flag.Parse()

	configData, err := os.ReadFile(*configPath)
	if err != nil {
		panic(fmt.Sprintf("Error reading config file: %s", err))
	}

	var config Config
	err = yaml.Unmarshal(configData, &config)
	if err != nil {
		panic(fmt.Sprintf("Error parsing config file: %s", err))
	}

}
