package clay

import "github.com/clipperhouse/typewriter"

var templates = typewriter.TemplateSlice{
	set,
}

var set = &typewriter.Template{
	Name: "ModelGenerator",
	Text: `
// Retrieves an object by id
func {{.Name}}ByIdentifier({{.PrimaryKey}} int64) {
}
`,
	TypeConstraint: typewriter.Constraint{Comparable: true},
}
