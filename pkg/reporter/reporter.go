package reporter

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"html/template"

	"github.com/rajatjindal/junit-to-html/pkg/parser"
)

func ToJson(data *TestDataWithMeta) ([]byte, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return b, nil
}

//go:embed index.html.gtpl
var tmpl string

type TestDataWithMeta struct {
	TitlePrimary   string         `json:"titlePrimary"`
	TitleSecondary string         `json:"titleSecondary"`
	Tags           []Tag          `json:"tags"`
	Tests          []*parser.Test `json:"tests"`
}

type Tag struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func ToHTML(data *TestDataWithMeta) ([]byte, error) {
	var buf bytes.Buffer

	t, err := template.New("report").Parse(tmpl)
	if err != nil {
		return nil, err
	}

	err = t.ExecuteTemplate(&buf, "report", data)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
