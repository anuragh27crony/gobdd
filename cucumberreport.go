package gobdd

import (
	"bufio"
	"fmt"
	"github.com/cucumber/gherkin-go/v13"
	msgs "github.com/cucumber/messages-go/v12"
	"os"
	"path/filepath"
)

type stepresult struct {
	status   string
	duration int64
}

type filelocation struct {
	location string
}

type step struct {
	result  stepresult
	match   filelocation
	keyword string
	name    string
	line    int32
}

func newstep(keyword string, name string, line int32, location string) step {
	return step{
		result:  stepresult{},
		match:   filelocation{location: location},
		keyword: keyword,
		name:    name,
		line:    line,
	}
}

func (s *step) updateStepResult(status string, duration int64) {
	s.result = stepresult{
		status:   status,
		duration: duration,
	}
}

type tag struct {
	name string
	line int32
}

type scenario struct {
	steps       []step
	tags        []tag
	id          string
	keyword     string
	name        string
	description string
	bddtype     string
}

func newscenario(steps []step, tags []tag, name string, id string, description string) scenario {
	return scenario{
		steps:       steps,
		tags:        tags,
		id:          id,
		keyword:     "Scenario",
		name:        name,
		description: description,
		bddtype:     "scenario",
	}
}
func (sc *scenario) addNewStep(keyword string, name string, line int32, location string) {
	newstep := step{
		result:  stepresult{},
		match:   filelocation{location: location},
		keyword: keyword,
		name:    name,
		line:    line,
	}
	sc.steps = append(sc.steps, newstep)
}

type feature struct {
	elements    []scenario
	uri         string
	id          string
	keyword     string
	name        string
	description string
	line        int32
}

type cucumberreport struct {
	features []feature
}

func newfeature(scenarios []scenario, name string, id string, description string, line int32) feature {
	return feature{
		elements:    scenarios,
		uri:         name,
		id:          id,
		keyword:     "Feature",
		name:        name,
		description: description,
		line:        0,
	}
}

func (f *feature) addNewScenario(steps []step, tags []tag, name string, id string, description string) {
	f.elements = append(f.elements, scenario{
		steps:       steps,
		tags:        tags,
		id:          id,
		keyword:     "Scenario",
		name:        name,
		description: description,
		bddtype:     "scenario",
	})
}

func main() {
	files, err := filepath.Glob("../features")
	if err != nil {
		fmt.Println("Error Occured")
	}

	for _, file := range files {
		f, err := os.Open(file)
		defer f.Close()
		fileIO := bufio.NewReader(f)

		doc, err := gherkin.ParseGherkinDocument(fileIO, (&msgs.Incrementing{}).NewId)
		if err != nil {
			fmt.Println("error while loading document: %s\n", err)
		}

		if doc.Feature == nil {
			fmt.Println("Empty Feature files")
		}

	}

}
//
//func format(feature *msgs.GherkinDocument_Feature) {
//
//}
