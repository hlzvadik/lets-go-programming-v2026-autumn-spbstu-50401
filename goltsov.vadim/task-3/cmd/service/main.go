package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/charmap"
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
	NumCode  int    `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
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

	xmlFile, err := os.Open(config.InputFile)
	if err != nil {
		panic(fmt.Sprintf("Error opening input file: %s", err))
	}
	defer xmlFile.Close()

	var valCurs ValCurs
	decoder := xml.NewDecoder(xmlFile)
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		if charset == "windows-1251" {
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		}
		return nil, fmt.Errorf("unknown charset: %s", charset)
	}

	err = decoder.Decode(&valCurs)
	if err != nil {
		panic(fmt.Sprintf("Error parsing input file: %s", err))
	}

	sort.Slice(valCurs.Valute, func(i, j int) bool {
		goodStringFloat1 := strings.ReplaceAll(valCurs.Valute[i].Value, ",", ".")
		goodFloat1, err := strconv.ParseFloat(goodStringFloat1, 64)
		if err != nil {
			panic(fmt.Sprintf("Error parsing input file: %s", err))
		}
		goodStringFloat2 := strings.ReplaceAll(valCurs.Valute[j].Value, ",", ".")
		goodFloat2, err := strconv.ParseFloat(goodStringFloat2, 64)
		if err != nil {
			panic(fmt.Sprintf("Error parsing input file: %s", err))
		}
		return goodFloat1 > goodFloat2
	})

	var outputItems []OutputItem
	for _, val := range valCurs.Valute {
		goodStringFloat := strings.ReplaceAll(val.Value, ",", ".")
		goodFloat, err := strconv.ParseFloat(goodStringFloat, 64)
		if err != nil {
			panic(fmt.Sprintf("Error parsing input file: %s", err))
		}
		outputItems = append(outputItems, OutputItem{
			NumCode:  val.NumCode,
			CharCode: val.CharCode,
			Value:    goodFloat,
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
