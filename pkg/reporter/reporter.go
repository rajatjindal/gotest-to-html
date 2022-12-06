package reporter

import (
	"bytes"
	"embed"
	"encoding/json"
	"html/template"

	"github.com/rajatjindal/gotest-to-html/pkg/parser"
)

func ToJson(data *TestDataWithMeta) ([]byte, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return b, nil
}

//go:embed templates/**/*
//go:embed templates/*
var templates embed.FS

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

	t, err := template.New("index.html.gtpl").ParseFS(templates, "templates/components/icons/*.gtpl", "templates/components/*.gtpl", "templates/index.html.gtpl")
	if err != nil {
		return nil, err
	}

	err = t.ExecuteTemplate(&buf, "index.html.gtpl", data)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
