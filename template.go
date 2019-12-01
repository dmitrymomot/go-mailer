package mailer

import (
	"bytes"
	"html/template"

	rice "github.com/GeertJohan/go.rice"
	"github.com/pkg/errors"
)

type (
	// EmailTemplate interface
	EmailTemplate interface {
		Subject() string
		HTMLTemplate() string
		PlainTextTemplate() string
	}
)

func getTemplateStr(tplName string) (string, error) {
	templateBox, err := rice.FindBox("templates")
	if err != nil {
		return "", errors.Wrap(err, "get template from box")
	}
	templateString, err := templateBox.String(tplName)
	if err != nil {
		return "", errors.Wrap(err, "get template content as string")
	}
	return templateString, nil
}

func parseTemplate(templateString string, opt map[string]interface{}) (string, error) {
	t, err := template.New("email").Parse(templateString)
	if err != nil {
		return "", errors.Wrap(err, "parse template")
	}
	var tpl bytes.Buffer
	if err := t.Execute(&tpl, opt); err != nil {
		return "", errors.Wrap(err, "execute template")
	}
	return tpl.String(), nil
}
