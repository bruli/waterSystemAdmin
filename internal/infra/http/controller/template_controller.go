package controller

import "github.com/flosch/pongo2/v6"

type TemplateController map[string]any

func (c TemplateController) Add(key string, value any) {
	c[key] = value
}

func (c TemplateController) AddError(msg string) {
	c["error_msg"] = msg
}

func (c TemplateController) toPongoContext() pongo2.Context {
	return pongo2.Context(c)
}

func newTemplateController() TemplateController {
	return make(TemplateController)
}
