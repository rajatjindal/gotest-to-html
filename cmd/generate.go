package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/rajatjindal/gotest-to-html/pkg/parser"
	"github.com/rajatjindal/gotest-to-html/pkg/reporter"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(actionCmd)
}

// actionCmd is the github action command
var actionCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate report for single execution",
	Run: func(cmd *cobra.Command, args []string) {
		err := generate()
		if err != nil {
			fmt.Println("ERROR ", err.Error())
			os.Exit(1)
		}
	},
}

func generate() error {
	tests, err := parser.IngestFile(getInputForAction("gotest_output_file"))
	if err != nil {
		return err
	}

	data := &reporter.TestDataWithMeta{
		TitlePrimary:   getInputForAction("title_primary"),
		TitleSecondary: getInputForAction("title_secondary"),
		Tests:          tests,
		Tags:           getTagsFromInput(getInputForAction("tags")),
	}

	out, err := reporter.ToHTML(data)
	if err != nil {
		return err
	}

	err = os.WriteFile(htmlOutputFile(), out, 0644)
	if err != nil {
		return err
	}

	if getInputForAction("json_output_file") != "" {
		out, err := reporter.ToJson(data)
		if err != nil {
			return err
		}

		err = os.WriteFile(jsonOutputFile(), out, 0644)
		if err != nil {
			return err
		}
	}

	return nil
}

func getInputForAction(key string) string {
	return os.Getenv(fmt.Sprintf("INPUT_%s", strings.ToUpper(key)))
}

func getTagsFromInput(s string) []reporter.Tag {
	tags := []reporter.Tag{}
	if len(s) == 0 {
		return tags
	}

	for _, pair := range strings.Split(s, ";") {
		parts := strings.Split(pair, "=")

		if len(parts) != 2 {
			continue
		}

		tags = append(tags, reporter.Tag{
			Key:   strings.TrimSpace(parts[0]),
			Value: strings.TrimSpace(parts[1]),
		})
	}

	return tags
}

func htmlOutputFile() string {
	file := "report.html"
	if getInputForAction("html_output_file") != "" {
		file = getInputForAction("html_output_file")
	}

	return filepath.Join(os.Getenv("GITHUB_WORKSPACE"), file)
}

func jsonOutputFile() string {
	return filepath.Join(os.Getenv("GITHUB_WORKSPACE"), getInputForAction("json_output_file"))
}
