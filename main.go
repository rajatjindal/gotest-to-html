package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
	"text/template"
	"time"

	junit "github.com/joshdk/go-junit"
	"github.com/sirupsen/logrus"
)

var allTests = []*Test{}

type Event struct {
	Test    string
	Package string
	Action  string
	Time    time.Time
	Elapsed float64
	Output  string
}

//go:embed tests.json
var raw string

//go:embed index.html.tmpl
var htmlTemplate string

func main() {
	dec := json.NewDecoder(strings.NewReader(raw))
	for {
		var event Event

		err := dec.Decode(&event)
		if err == io.EOF {
			// all done
			break
		}
		if err != nil {
			log.Fatal("here ", err)
		}

		processEvent(event)
	}

	// for _, test := range allTests {
	// 	if test.Parent != nil {
	// 		continue
	// 	}

	// 	printChildren(test)
	// }

	t, err := template.New("report").Funcs(map[string]interface{}{
		"getHumanReadableName": getHumanReadableName,
		"getIcon":              getIcon,
	}).Parse(htmlTemplate)
	if err != nil {
		logrus.Fatal(err)
	}

	onlyParents := []*Test{}
	for _, test := range allTests {
		if test.Parent == nil {
			onlyParents = append(onlyParents, test)
		}
	}

	data := struct {
		Tests []*Test
	}{
		Tests: onlyParents,
	}

	err = t.ExecuteTemplate(os.Stdout, "report", data)
	if err != nil {
		logrus.Fatal(err)
	}
}

type Test struct {
	PrimaryKey string
	Name       string
	Parent     *Test
	Children   []*Test
	Result     string
	Duration   float64
	Systemout  string
}

func printChildren(test *Test) {
	fmt.Println(test.Name, " - ", test.Result, " - ", test.Duration)
	for _, child := range test.Children {
		if len(child.Children) > 0 {
			printChildren(child)
		} else {
			fmt.Println(child.Name, " - ", child.Result, " - ", child.Duration)
		}
	}
}
func processEvent(event Event) {
	if event.Action == "run" {
		parent := findParentTest(event)
		current := &Test{Name: event.Test, Parent: parent, PrimaryKey: event.Package + "|" + event.Test}
		if parent != nil {
			parent.Children = append(parent.Children, current)
		}

		allTests = append(allTests, current)
	} else if event.Action == "output" {
		test := findTest(event)
		test.Systemout = event.Output
	} else if event.Action == "pass" {
		test := findTest(event)
		test.Result = event.Action
		test.Duration = event.Elapsed
	} else if event.Action == "fail" {
		test := findTest(event)
		test.Result = event.Action
		test.Duration = event.Elapsed
	} else if event.Action == "skip" {
		test := findTest(event)
		test.Result = event.Action
		test.Duration = event.Elapsed
	} else {
		fmt.Println(event.Action)
	}
}

func findParentTest(event Event) *Test {
	if len(allTests) == 0 {
		return nil
	}

	if !strings.Contains(event.Test, "/") {
		return nil
	}

	name := event.Test
	parentName := name

	parts := strings.Split(name, "/")
	parentName = strings.Join(parts[:len(parts)-1], "/")

	parentPrimaryKey := event.Package + "|" + parentName
	for _, test := range allTests {
		if test.PrimaryKey == parentPrimaryKey {
			return test
		}
	}

	return nil
}

func findTest(event Event) *Test {
	if len(allTests) == 0 {
		return &Test{
			Name:       event.Test,
			Children:   []*Test{},
			PrimaryKey: event.Package + "|" + event.Test,
		}
	}

	primaryKey := event.Package + "|" + event.Test
	for _, test := range allTests {
		if test.PrimaryKey == primaryKey {
			return test
		}
	}

	return &Test{
		Name:       event.Test,
		Children:   []*Test{},
		PrimaryKey: event.Package + "|" + event.Test,
	}
}

func getIcon(status junit.Status) string {
	switch status {
	case junit.StatusPassed:
		return ":white_check_mark:"
	case junit.StatusFailed:
		return ":x:"
	case junit.StatusError:
		return ":question:"
	case junit.StatusSkipped:
		return ":grey_question:"
	}

	return ":thinking:"
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
