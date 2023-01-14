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
	inp := getInputForAction("gotest_output_file")
	files := []string{}

	if stat, _ := os.Stat(inp); stat != nil && stat.IsDir() {
		list, err := os.ReadDir(inp)
		if err != nil {
			return err
		}

		for _, l := range list {
			files = append(files, l.Name())
		}
	} else {
		for _, file := range strings.Split(inp, ",") {
			_, err := os.Stat(file)
			if os.IsNotExist(err) {
				continue
			}

			files = append(files, file)
		}
	}

	for _, file := range files {
		generateOne(file)
	}

	return nil
}

func generateOne(file string) error {
	tests, err := parser.IngestFile(file)
	if err != nil {
		return err
	}

	title := filepath.Base(file)
	title = strings.ReplaceAll(title, ".json", "")
	title = strings.ReplaceAll(title, "-", " ")

	data := &reporter.TestDataWithMeta{
		TitlePrimary:   title,
		TitleSecondary: getInputForAction("title_secondary"),
		Tests:          tests,
		Tags:           getTagsFromInput(getInputForAction("tags")),
	}

	out, err := reporter.ToHTML(data)
	if err != nil {
		return err
	}

	err = os.WriteFile(strings.ReplaceAll(file, ".json", "-report.html"), out, 0644)
	if err != nil {
		return err
	}

	out, err = reporter.ToJson(data)
	if err != nil {
		return err
	}

	err = os.WriteFile(strings.ReplaceAll(file, ".json", "-report.json"), out, 0644)
	if err != nil {
		return err
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
