package parser

import (
	"regexp"
	"strings"
	"time"
)

type Event struct {
	Test    string
	Package string
	Action  string
	Time    time.Time
	Elapsed float64
	Output  string
}

type eventProcessor struct {
	allTests  []*Test
	TestsTree []*Test
}

func (p *eventProcessor) processEvent(event Event) {
	switch event.Action {
	case "run":
		parent := p.findParentTest(event)
		current := &Test{Name: getHumanReadableName(event.Test), Parent: parent != nil, Id: event.Package + "|" + event.Test}
		if parent != nil {
			parent.Children = append(parent.Children, current)
		}

		if parent == nil {
			p.TestsTree = append(p.TestsTree, current)
		}

		p.allTests = append(p.allTests, current)
	case "output":
		test := p.findTest(event)
		test.Logs = append(test.Logs, event.Output)
	case "pass":
		test := p.findTest(event)
		test.Result = event.Action
		test.Duration = event.Elapsed
	case "fail":
		test := p.findTest(event)
		test.Result = event.Action
		test.Duration = event.Elapsed
	case "skip":
		test := p.findTest(event)
		test.Result = event.Action
		test.Duration = event.Elapsed
	}
}

// TODO: findParentTest and findTest should be optimized for faster find
func (p *eventProcessor) findParentTest(event Event) *Test {
	if len(p.allTests) == 0 {
		return nil
	}

	if !strings.Contains(event.Test, "/") {
		return nil
	}

	name := event.Test
	parentName := name

	parts := strings.Split(name, "/")
	parentName = strings.Join(parts[:len(parts)-1], "/")

	parentId := event.Package + "|" + parentName
	for _, test := range p.allTests {
		if test.Id == parentId {
			return test
		}
	}

	return nil
}

func (p *eventProcessor) findTest(event Event) *Test {
	if len(p.allTests) == 0 {
		return &Test{
			Name:     getHumanReadableName(event.Test),
			Children: []*Test{},
			Id:       event.Package + "|" + event.Test,
		}
	}

	primaryKey := event.Package + "|" + event.Test
	for _, test := range p.allTests {
		if test.Id == primaryKey {
			return test
		}
	}

	return &Test{
		Name:     getHumanReadableName(event.Test),
		Children: []*Test{},
		Id:       event.Package + "|" + event.Test,
	}
}

func getHumanReadableName(name string) string {
	if !strings.Contains(name, "/") {
		re := regexp.MustCompile(`[A-Z][^A-Z]*`)
		tokens := re.FindAllString(name, -1)
		return strings.Join(tokens, " ")
	}

	if len(strings.Split(name, "/")) == 2 {
		namex := strings.Split(name, "/")[1]
		return strings.Join(strings.Split(namex, "_"), " ")
	}

	if len(strings.Split(name, "/")) > 2 {
		namex := strings.Split(name, "/")[2]
		return strings.Join(strings.Split(namex, "_"), " ")
	}

	return name
}
