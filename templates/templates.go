package templates

import "html/template"

// Template define and take directory where the templates are stored
func Template() *template.Template {
	return template.Must(template.ParseGlob("views/*"))
}
