package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"

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

type OutputItem struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
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

	sort.Slice(valCurs.Valute, func(i, j int) bool {
		return valCurs.Valute[i].Value < valCurs.Valute[j].Value
	})

	var outputItems []OutputItem
	for _, val := range valCurs.Valute {
		outputItems = append(outputItems, OutputItem{
			NumCode:  val.NumCode,
			CharCode: val.CharCode,
			Value:    val.Value,
		})
	}

	jsonData, err := json.MarshalIndent(outputItems, "", "  ")
	if err != nil {
		panic(fmt.Sprintf("Error marshalling json: %s", err))
	}

	outputDir := filepath.Dir(config.OutputFile)
	err = os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		panic(fmt.Sprintf("Error creating output directory: %s", err))
	}

	err = os.WriteFile(config.OutputFile, jsonData, os.ModePerm)
	if err != nil {
		panic(fmt.Sprintf("Error creating output file: %s", err))
	}

}
