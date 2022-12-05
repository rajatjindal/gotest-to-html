package parser

import (
	"encoding/json"
	"io"
	"os"
)

func IngestFile(filename string) ([]*Test, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	p := &eventProcessor{
		allTests:  map[string]*Test{},
		TestsTree: []*Test{},
	}

	dec := json.NewDecoder(file)
	for {
		var event Event

		err := dec.Decode(&event)
		if err == io.EOF {
			// all done
			break
		}

		if err != nil {
			return nil, err
		}

		p.processEvent(event)
	}

	return p.TestsTree, nil
}
