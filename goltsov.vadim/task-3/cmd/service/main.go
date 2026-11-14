package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Valute  []Valute `xml:"Valute"`
}

type Valute struct {
	NumCode  int     `xml:"NumCode"`
	CharCode string  `xml:"CharCode"`
	Value    float64 `xml:"Value"`
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

	xmlData, err := os.ReadFile(config.InputFile)
	if err != nil {
		panic(fmt.Sprintf("Error reading input file: %s", err))
	}

	var valCurs ValCurs
	err = xml.Unmarshal(xmlData, &valCurs)
	if err != nil {
		panic(fmt.Sprintf("Error parsing input file: %s", err))
	}

}
