package cucumber

import (
	msgs "github.com/cucumber/messages-go/v12"
)

func FormatFeatureWithScenario(gherkinFeature *msgs.GherkinDocument_Feature) Feature {
	ft := GenerateFeature(gherkinFeature.GetName(), gherkinFeature.GetName(), gherkinFeature.GetDescription(), int(gherkinFeature.Location.GetLine()))

	for _, child := range gherkinFeature.Children {
		ft.AddScenario(FormatScenarioWithSteps(child.GetScenario(), ""))
	}
	return ft
}

func FormatFeature(gherkinFeature *msgs.GherkinDocument_Feature) Feature {
	ft := GenerateFeature(gherkinFeature.GetName(), gherkinFeature.GetName(), gherkinFeature.GetDescription(), int(gherkinFeature.Location.GetLine()))
	return ft
}

func FormatScenario(gherkinScenario *msgs.GherkinDocument_Feature_Scenario) Scenario {
	sc := Scenario{
		Steps: nil, Tags: FormatTags(gherkinScenario.GetTags()),
		Id:          gherkinScenario.GetId(),
		Keyword:     gherkinScenario.GetKeyword(),
		Name:        gherkinScenario.GetName(),
		Description: gherkinScenario.GetDescription(),
		Type:        gherkinScenario.GetKeyword(),
	}
	return sc
}

func FormatScenarioWithSteps(gherkinScenario *msgs.GherkinDocument_Feature_Scenario, stepstatus string) Scenario {
	sc := Scenario{
		Steps: nil, Tags: FormatTags(gherkinScenario.GetTags()),
		Id:          gherkinScenario.GetId(),
		Keyword:     gherkinScenario.GetKeyword(),
		Name:        gherkinScenario.GetName(),
		Description: gherkinScenario.GetDescription(),
		Type:        gherkinScenario.GetKeyword(),
	}

	for _, steps := range gherkinScenario.Steps {
		sc.AddStep(steps.GetKeyword(), steps.GetText(), int(steps.Location.GetLine()), "", stepstatus)
	}
	return sc
}

func FormatTags(cucumberTags []*msgs.GherkinDocument_Feature_Tag) []Tag {
	var tags []Tag

	for _, cucumbertag := range cucumberTags {
		tags = append(tags, Tag{Name: cucumbertag.Name, Linenumber: int(cucumbertag.Location.Line)})
	}

	return tags
}
