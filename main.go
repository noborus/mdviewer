package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"

	"github.com/charmbracelet/glamour"
	"github.com/noborus/ov/oviewer"
)

func main() {
	fileName := os.Args[1]

	r, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(0),
	)
	if err != nil {
		log.Fatal(err)
	}

	source, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	result, err := r.Render(string(source))
	if err != nil {
		log.Fatal(err)
	}
	buff := bytes.NewBuffer([]byte(result))
	render, err := oviewer.NewDocument()
	if err != nil {
		log.Fatal(err)
	}
	render.FileName = fileName + "(Render)"
	err = render.ReadAll(buff)
	if err != nil {
		log.Fatal(err)
	}

	original, err := oviewer.NewDocument()
	if err != nil {
		log.Fatal(err)
	}
	original.FileName = fileName
	err = original.ReadAll(bytes.NewBuffer(source))
	if err != nil {
		log.Fatal(err)
	}

	root, err := oviewer.NewOviewer(render, original)
	if err != nil {
		log.Fatal(err)
	}

	err = root.Run()
	if err != nil {
		log.Fatal(err)
	}
}
