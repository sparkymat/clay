package clay

import (
	"io"

	"github.com/clipperhouse/typewriter"
)

type ModelGenerator struct{}

func NewModelGenerator() *ModelGenerator {
	return &ModelGenerator{}
}

func (g *ModelGenerator) Name() string {
	return "modelgen"
}

func (g *ModelGenerator) Imports(t typewriter.Type) []typewriter.ImportSpec {
	return []typewriter.ImportSpec{}
}

func (g *ModelGenerator) Write(writer io.Writer, t typewriter.Type) error {
	tag, found := t.FindTag(g)

	if !found {
		return nil
	}

	tmpl, err := templates.ByTag(t, tag)

	if err != nil {
		return err
	}

	if err := tmpl.Execute(writer, t); err != nil {
		return err
	}

	return nil
}
