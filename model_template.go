package clay

import "github.com/clipperhouse/typewriter"

var templates = typewriter.TemplateSlice{
	core,
	relation,
}

var core = &typewriter.Template{
	Name: "core",
	Text: `
// Retrieves an object by id
func {{.ModelClassName}}ByIdentifier(identifier int64) (*{{.ModelClassName}}, error) {
  return nil, errors.New("Unable to load object")
}
`,
}

var relation = &typewriter.Template{
	Name: "relation",
	Text: `
func (m *{{.ModelClassName}}) {{.RelationClassName}}() (*{{.RelationClassName}}, error) {
  if m.{{.RelationName}} != nil {
    return m.{{.RelationName}}, nil
  }

  if m.{{.RelationIdentifierName}} == nil {
    return nil, errors.New("Foreign key is nil")
  }

  {{.RelationName}}, err := {{.RelationClassName}}ByIdentifier(m.{{.RelationIdentifierName}})
  if err != nil {
    return nil, errors.New("Unable to load object")
  }

  m.{{.RelationName}} = {{.RelationName}}

  return m.{{.RelationName}}, nil
}
`,
}
