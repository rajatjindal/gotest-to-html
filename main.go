package main

import (
	"fmt"

	"github.com/rajatjindal/junit-to-html/pkg/parser"
	"github.com/rajatjindal/junit-to-html/pkg/reporter"
	"github.com/sirupsen/logrus"
)

func main() {
	tests, err := parser.IngestFile("tests.json")
	if err != nil {
		logrus.Fatal(err)
	}

	out, err := reporter.ToHTML(tests)
	if err != nil {
		logrus.Fatal(err)
	}

	fmt.Println(out)
}
