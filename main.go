package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/rajatjindal/junit-to-html/pkg/parser"
	"github.com/rajatjindal/junit-to-html/pkg/reporter"
)

func main() {
	tests, err := parser.IngestFile(getInputForAction("gotest_output_file"))
	if err != nil {
		fmt.Println("ERROR ", err.Error())
		os.Exit(1)
	}

	data := &reporter.TestDataWithMeta{
		TitlePrimary:   getInputForAction("title_primary"),
		TitleSecondary: getInputForAction("title_secondary"),
		Tests:          tests,
		Tags:           getTagsFromInput(getInputForAction("tags")),
	}

	out, err := reporter.ToHTML(data)
	if err != nil {
		fmt.Println("ERROR ", err.Error())
		os.Exit(1)
	}

	err = os.WriteFile(getInputForAction("html_report_file"), out, 0644)
	if err != nil {
		fmt.Println("ERROR ", err.Error())
		os.Exit(1)
	}
}

func getInputForAction(key string) string {
	return os.Getenv(fmt.Sprintf("INPUT_%s", strings.ToUpper(key)))
}

func getTagsFromInput(s string) []reporter.Tag {
	tags := []reporter.Tag{}
	if len(s) == 0 {
		return tags
	}

	for _, pair := range strings.Split(s, " ") {
		parts := strings.Split(pair, "=")

		if len(parts) != 2 {
			continue
		}

		tags = append(tags, reporter.Tag{
			Key:   parts[0],
			Value: parts[1],
		})
	}

	return tags
}
