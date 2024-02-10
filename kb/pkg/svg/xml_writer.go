package svg

import (
	"bytes"
	"os"

	"github.com/beevik/etree"
)

type xmlWriterSettings struct {
	Indent              bool
	IndentChars         string
	NewLineOnAttributes bool
}

type xmlWriter struct {
	settings *xmlWriterSettings
	doc      *etree.Document
	root     *etree.Element
	curr     *etree.Element
}

func CreateXMLWriter(settings *xmlWriterSettings) *xmlWriter {
	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)

	return &xmlWriter{
		settings: settings,
		doc:      doc,
		root:     nil,
		curr:     nil,
	}
}

func (w *xmlWriter) writeStartElement(tag string) {
	if w.curr == nil {
		w.curr = w.doc.CreateElement(tag)
	} else {
		w.curr = w.curr.CreateElement(tag)
	}
}

func (w *xmlWriter) writeEndElement() {
	w.curr = w.curr.Parent()
}

func (w *xmlWriter) writeAttributeString(name string, value string) {
	w.curr.CreateAttr(name, value)
}

func (w *xmlWriter) writeText(text string) {
	w.curr.CreateText(text)
}

func (w *xmlWriter) writeToFile(outFile string) {
	w.doc.Indent(4)
	b := &bytes.Buffer{}
	w.doc.WriteTo(b)
	os.WriteFile(outFile, b.Bytes(), 0644)
}
