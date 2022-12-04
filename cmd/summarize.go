package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/rajatjindal/gotest-to-html/pkg/reporter"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(summarizeCmd)
}

// actionCmd is the github action command
var summarizeCmd = &cobra.Command{
	Use:   "summarize",
	Short: "generate historical report",
	Run: func(cmd *cobra.Command, args []string) {
		err := summarize()
		if err != nil {
			fmt.Println("ERROR ", err.Error())
			os.Exit(1)
		}
	},
}

type Summary struct {
	Total      int
	Passed     int
	Failed     int
	Skipped    int
	Duration   float64
	Executions int
}

func summarize() error {
	datadir := "data/executions"
	dirs, err := os.ReadDir(datadir)
	if err != nil {
		return err
	}

	limitdirs := dirs
	limit := 10
	if len(dirs) > limit {
		limitdirs = dirs[len(dirs)-limit:]
	}

	groupBy := "os,spin"
	if getInputForAction("group_by") != "" {
		groupBy = getInputForAction("group_by")
	}

	groupByTags := strings.Split(groupBy, ",")

	groupedData := map[string][]reporter.TestDataWithMeta{}
	for _, dir := range limitdirs {
		rawfile := filepath.Join(datadir, dir.Name(), "raw.json")
		raw, err := os.ReadFile(rawfile)
		if err != nil {
			fmt.Printf("skipping file %s due to err: %s\n", rawfile, err.Error())
			continue
		}

		var data reporter.TestDataWithMeta
		err = json.Unmarshal(raw, &data)
		if err != nil {
			fmt.Printf("skipping file %s due to err: %s\n", rawfile, err.Error())
			continue
		}

		groupkey := getGroupKey(data, groupByTags)

		groupedData[groupkey] = append(groupedData[groupkey], data)
	}

	summary := []Summary{}
	for key, value := range groupedData {
		summary = append(summary, getSummary(key, value))
	}

	fmt.Printf("%#v\n", summary)

	return nil
}

func getSummary(key string, data []reporter.TestDataWithMeta) Summary {
	summary := Summary{}
	for _, d := range data {
		summary.Executions++
		for _, test := range d.Tests {
			summary.Total++
			summary.Duration += test.Duration

			switch test.Result {
			case "pass":
				summary.Passed++
			case "fail":
				summary.Failed++
			case "skip":
				summary.Skipped++
			}
		}
	}

	return summary
}

func getGroupKey(data reporter.TestDataWithMeta, groupByTags []string) string {
	groupkey := ""
	for _, key := range groupByTags {
		for _, tk := range data.Tags {
			if tk.Key != key {
				continue
			}

			groupkey += key + ":" + tk.Value + "|"
		}
	}

	return groupkey
}
