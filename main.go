package main

import (
	"fmt"
	"os"

	"github.com/rajatjindal/junit-to-html/pkg/parser"
	"github.com/rajatjindal/junit-to-html/pkg/reporter"
	"github.com/sirupsen/logrus"
)

func main() {
	tests, err := parser.IngestFile(os.Args[1])
	if err != nil {
		logrus.Fatal(err)
	}

	data := &reporter.TestDataWithMeta{
		TitlePrimary:   os.Getenv("title-primary"),
		TitleSecondary: os.Getenv("title-secondary"),
		Tests:          tests,
		Tags: []reporter.Tag{
			{
				Key:   "golang",
				Value: "1.17",
			},
		},
	}
	out, err := reporter.ToHTML(data)
	if err != nil {
		logrus.Fatal(err)
	}

	fmt.Println(out)
}
