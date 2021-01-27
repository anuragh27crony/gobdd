package report

import (
	"github.com/anuragh27crony/gobdd/formatter/cucumber"
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/cucumber/gherkin-go/v13"
	msgs "github.com/cucumber/messages-go/v12"
	"io/ioutil"
	"os"
	"path/filepath"
)

func parseFeatures() []cucumber.Feature {
	files, err := filepath.Glob("features/*.feature")
	if err != nil {
		fmt.Println("Error Occured")
	}
	var features []cucumber.Feature

	for _, file := range files {
		f, err := os.Open(file)
		defer f.Close()
		fileIO := bufio.NewReader(f)

		doc, err := gherkin.ParseGherkinDocument(fileIO, (&msgs.Incrementing{}).NewId)
		if err != nil {
			fmt.Printf("error while loading document: %s\n", err)
		}

		if doc.Feature == nil {
			fmt.Println("Empty Feature FIles")
		}

		features = append(features, cucumber.FormatFeature(doc.Feature))
	}
	return features
}

func GenerateJson() {
	writeJsonFile("suite.json", parseFeatures())
}

func writeJsonFile(jsonFilePath string, data interface{}) {
	b, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ioutil.WriteFile(jsonFilePath, b, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
}
