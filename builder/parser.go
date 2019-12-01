package builder

import (
	"bytes"
	"html/template"

	"github.com/pkg/errors"
)

func parseTemplate(templateString string, data interface{}) (string, error) {
	t, err := template.New("email").Funcs(template.FuncMap{
		"inc": func(i int) int { return i + 1 },
	}).Parse(templateString)
	if err != nil {
		return "", errors.Wrap(err, "parse template")
	}
	var tpl bytes.Buffer
	if err := t.Execute(&tpl, data); err != nil {
		return "", errors.Wrap(err, "execute template")
	}
	return tpl.String(), nil
}
