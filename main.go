package main

import (
	"fmt"
	"os"

	"github.com/rajatjindal/junit-to-html/pkg/parser"
	"github.com/rajatjindal/junit-to-html/pkg/reporter"
)

func main() {
	tests, err := parser.IngestFile(os.Args[1])
	if err != nil {
		fmt.Println("ERROR ", err.Error())
		os.Exit(1)
	}

	data := &reporter.TestDataWithMeta{
		TitlePrimary:   "Spin",
		TitleSecondary: "e2e test report",
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
		fmt.Println("ERROR ", err.Error())
		os.Exit(1)
	}

	fmt.Println(out)
}
