package gogen

import (
	"bytes"
	"errors"
	"io/ioutil"

	"text/template"

	"github.com/swiftcarrot/genny"
)

var TemplateHelpers = map[string]interface{}{}

// TemplateTransformer will run any file that has a ".tmpl" extension through text/template
func TemplateTransformer(data interface{}, helpers map[string]interface{}) genny.Transformer {
	if helpers == nil {
		helpers = TemplateHelpers
	}
	t := genny.NewTransformer(".tmpl", func(f genny.File) (genny.File, error) {
		return renderWithTemplate(f, data, helpers)
	})
	t.StripExt = true
	return t
}

func renderWithTemplate(f genny.File, data interface{}, helpers template.FuncMap) (genny.File, error) {
	if f == nil {
		return f, errors.New("file was nil")
	}
	path := f.Name()
	t := template.New(path)
	if helpers != nil {
		t = t.Funcs(helpers)
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return f, err
	}
	t, err = t.Parse(string(b))
	if err != nil {
		return f, err
	}

	var bb bytes.Buffer
	if err = t.Execute(&bb, data); err != nil {
		err = err
		return f, err
	}
	return genny.StripExt(genny.NewFile(path, &bb), ".tmpl"), nil
}
