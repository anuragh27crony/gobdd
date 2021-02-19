package cucumber

type Feature struct {
	Elements    []Scenario `json:"elements"`
	Uri         string     `json:"uri"`
	Id          string     `json:"id"`
	Keyword     string     `json:"keyword"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Linenumber  int        `json:"line"`
}

type Scenario struct {
	Steps       []Step `json:"steps"`
	Tags        []Tag  `json:"tags"`
	Id          string `json:"id"`
	Keyword     string `json:"keyword"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

type Tag struct {
	Name       string `json:"name"`
	Linenumber int    `json:"line"`
}

type Step struct {
	StepResult Stepresult   `json:"result"`
	match      Filelocation `json:"match"`
	Keyword    string       `json:"keyword"`
	Name       string       `json:"name"`
	Line       int          `json:"line"`
}

type Stepresult struct {
	ErrorMsg      string `json:"error_message"`
	RunStatus     string `json:"status"`
	ExecutionTime int64  `json:"duration"`
}

type Filelocation struct {
	Location string `json:"location"`
}

func GenerateScenario() Scenario {
	return Scenario{}
}

func GenerateFeature(name string, id string, description string, line int) Feature {
	return Feature{
		Elements:    nil,
		Uri:         name,
		Id:          id,
		Keyword:     "Feature",
		Name:        name,
		Description: description,
		Linenumber:  line,
	}
}

func (f *Feature) AddScenario(sc Scenario) {
	f.Elements = append(f.Elements, sc)
}

func GenerateStep(keyword string, name string, line int, location string) Step {
	newstep := Step{
		StepResult: Stepresult{},
		match:      Filelocation{Location: location},
		Keyword:    keyword,
		Name:       name,
		Line:       line,
	}
	return newstep
}

func (sc *Scenario) AddStepObj(step Step) {
	sc.Steps = append(sc.Steps, step)
}

func (sc *Scenario) AddStep(keyword string, name string, line int, location string, result string) {
	newstep := Step{
		StepResult: Stepresult{
			RunStatus: result,
		},
		match:   Filelocation{Location: location},
		Keyword: keyword,
		Name:    name,
		Line:    line,
	}
	sc.Steps = append(sc.Steps, newstep)
}

func (s *Step) UpdateResult(status string, duration int64) {
	s.StepResult = Stepresult{
		RunStatus:     status,
		ExecutionTime: duration,
	}
}
