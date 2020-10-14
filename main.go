package main

import (
	"bytes"
	"io/ioutil"
	"os"

	"github.com/charmbracelet/glamour"
	"github.com/noborus/ov/oviewer"
)

func main() {
	fileName := os.Args[1]
	source, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	result, err := glamour.Render(string(source), "dark")
	if err != nil {
		panic(err)
	}
	buff := bytes.NewBuffer([]byte(result))
	render, err := oviewer.NewDocument()
	if err != nil {
		panic(err)
	}
	render.FileName = fileName + "(Render)"
	err = render.ReadAll(ioutil.NopCloser(buff))
	if err != nil {
		panic(err)
	}

	original, err := oviewer.NewDocument()
	if err != nil {
		panic(err)
	}
	original.FileName = fileName
	err = original.ReadAll(ioutil.NopCloser(bytes.NewBuffer(source)))
	if err != nil {
		panic(err)
	}

	root, err := oviewer.NewOviewer(render, original)
	if err != nil {
		panic(err)
	}

	err = root.Run()
	if err != nil {
		panic(err)
	}
}
