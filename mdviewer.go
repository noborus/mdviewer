package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/glamour"
	"github.com/noborus/ov/oviewer"
)

func MDViewer(only bool, args []string) (*oviewer.Root, error) {
	r, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(0),
	)
	if err != nil {
		return nil, err
	}

	var documents []*oviewer.Document
	for _, arg := range args {
		m, err := mdDocuments(r, only, arg)
		if err != nil {
			log.Println(err)
			continue
		}
		documents = append(documents, m...)
	}

	if len(documents) == 0 {
		return nil, fmt.Errorf("no file to display")
	}

	return oviewer.NewOviewer(documents...)

}

// mdDocuments returns markDown documents (rendered and original).
func mdDocuments(r *glamour.TermRenderer, only bool, fileName string) ([]*oviewer.Document, error) {
	source, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	// Rendered document
	renderDoc, err := oviewer.NewDocument()
	if err != nil {
		return nil, err
	}
	renderDoc.FileName = fileName + "(Render)"

	render, err := r.RenderBytes(source)
	if err != nil {
		return nil, err
	}
	if err := renderDoc.ControlReader(bytes.NewBuffer(render), nil); err != nil {
		return nil, err
	}

	if only {
		return []*oviewer.Document{renderDoc}, nil
	}

	// Original document
	originalDoc, err := oviewer.NewDocument()
	if err != nil {
		return nil, err
	}
	originalDoc.FileName = fileName

	if err := originalDoc.ControlReader(bytes.NewBuffer(source), nil); err != nil {
		return nil, err
	}

	return []*oviewer.Document{renderDoc, originalDoc}, nil
}
