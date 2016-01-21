package clay

import "github.com/clipperhouse/typewriter"

var templates = typewriter.TemplateSlice{
	set,
}

var set = &typewriter.Template{
	Name: "ModelGenerator",
	Text: `
// Clear clears the entire set to be the empty set.
func (set *{{.Name}}Set) Clear() {
	*set = make({{.Name}}Set)
}
`,
	TypeConstraint: typewriter.Constraint{Comparable: true},
}
