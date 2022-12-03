package reporter

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"html/template"

	"github.com/rajatjindal/junit-to-html/pkg/parser"
)

func ToJson(tests []*parser.Test) (string, error) {
	b, err := json.Marshal(tests)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

//go:embed index.html.gtpl
var tmpl string

type TestDataWithMeta struct {
	TitlePrimary   string
	TitleSecondary string
	Tags           []Tag
	Tests          []*parser.Test
}
type Tag struct {
	Key   string
	Value string
}

func ToHTML(data *TestDataWithMeta) (string, error) {
	var buf bytes.Buffer

	t, err := template.New("report").Parse(tmpl)
	if err != nil {
		return "", err
	}

	err = t.ExecuteTemplate(&buf, "report", data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
