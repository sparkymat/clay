package clay

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/clipperhouse/typewriter"
)

func init() {
	err := typewriter.Register(NewModelGenerator())
	if err != nil {
		panic(err)
	}
}

type ModelGenerator struct{}

func NewModelGenerator() *ModelGenerator {
	return &ModelGenerator{}
}

func (g *ModelGenerator) Name() string {
	return "model"
}

func (g *ModelGenerator) Imports(t typewriter.Type) []typewriter.ImportSpec {
	return []typewriter.ImportSpec{}
}

func (g *ModelGenerator) Write(writer io.Writer, t typewriter.Type) error {
	tag, found := t.FindTag(g)

	if !found {
		return nil
	}

	// Fetch core template
	coreTemplates := templates.Where(func(t *typewriter.Template) bool {
		if t != nil && t.Name == "core" {
			return true
		}

		return false
	})
	if len(coreTemplates) == 0 {
		return errors.New("Unable to find 'core' template")
	}
	coreTemplate, err := coreTemplates[0].Parse()
	if err != nil {
		return errors.New("Unable to parse 'core' template")
	}

	// Fetch relation template
	relationTemplates := templates.Where(func(t *typewriter.Template) bool {
		if t != nil && t.Name == "relation" {
			return true
		}

		return false
	})
	if len(relationTemplates) == 0 {
		return errors.New("Unable to find 'relation' template")
	}
	relationTemplate, err := relationTemplates[0].Parse()
	if err != nil {
		return errors.New("Unable to parse 'relation' template")
	}

	cValues := coreValues{ModelClassName: t.Name}
	if err := coreTemplate.Execute(writer, cValues); err != nil {
		return err
	}

	for _, relation := range tag.Values {
		templateValues := relationValues{}

		templateValues.ModelClassName = t.Name
		templateValues.RelationIdentifierName = fmt.Sprintf("%vIdentifier", strings.Title(relation.Name))
		templateValues.RelationName = relation.Name
		templateValues.RelationClassName = strings.Title(relation.Name)

		if err := relationTemplate.Execute(writer, templateValues); err != nil {
			return err
		}
	}

	return nil
}
