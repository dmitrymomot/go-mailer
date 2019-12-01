package builder

import (
	"log"

	"github.com/pkg/errors"
)

// Ul unordered list
func Ul(items []string) Part {
	htmlTemplate := `
	{{if .items}}
	<ul style="font-size: 16px; line-height: 1.625; margin: .4em 0 1.1875em;">
		{{range $item := .items}}
		<li>{{$item}}</li>
		{{end}}
	</ul>
	{{end}}`
	htmlStr, err := parseTemplate(htmlTemplate, map[string]interface{}{
		"items": items,
	})
	if err != nil {
		log.Fatal(errors.Wrap(err, "unordered list html template"))
	}

	plainTextTemplate := `
{{if .items}}{{range $item := .items}}
* {{$item}}{{end}}{{end}}`
	plainText, err := parseTemplate(plainTextTemplate, map[string]interface{}{
		"items": items,
	})
	if err != nil {
		log.Fatal(errors.Wrap(err, "unordered list plain text template"))
	}

	return newPart(htmlStr, plainText)
}

// Ol ordered list
func Ol(items []string) Part {
	htmlTemplate := `
	{{if .items}}
	<ul style="font-size: 16px; line-height: 1.625; margin: .4em 0 1.1875em;">
		{{range $item := .items}}
		<li>{{$item}}</li>
		{{end}}
	</ul>
	{{end}}`
	htmlStr, err := parseTemplate(htmlTemplate, map[string]interface{}{
		"items": items,
	})
	if err != nil {
		log.Fatal(errors.Wrap(err, "ordered list html template"))
	}

	plainTextTemplate := `
{{if .items}}{{range $k, $item := .items}}
{{inc $k}}. {{$item}}{{end}}{{end}}`
	plainText, err := parseTemplate(plainTextTemplate, map[string]interface{}{
		"items": items,
	})
	if err != nil {
		log.Fatal(errors.Wrap(err, "ordered list plain text template"))
	}

	return newPart(htmlStr, plainText)
}
