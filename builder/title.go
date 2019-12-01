package builder

import (
	"log"

	"github.com/pkg/errors"
)

// H1 ...
func H1(title string) Part {
	htmlTemplate := `<h1 style="margin-top: 0; color: #2D3748; font-size: 22px; font-weight: bold; text-align: left;"
	align="left">{{.title}}</h1>`
	htmlStr, err := parseTemplate(htmlTemplate, map[string]interface{}{
		"title": title,
	})
	if err != nil {
		log.Fatal(errors.Wrap(err, "h1 html template"))
	}

	plainTextTemplate := `
*****
{{.title}}
*****`
	plainText, err := parseTemplate(plainTextTemplate, map[string]interface{}{
		"title": title,
	})
	if err != nil {
		log.Fatal(errors.Wrap(err, "h1 plain text template"))
	}

	return newPart(htmlStr, plainText)
}

// H2 ...
func H2(title string) Part {
	htmlTemplate := `<h2 style="margin-top: 0; color: #2D3748; font-size: 16px; font-weight: bold; text-align: left;"
	align="left">{{title}}</h2>`
	htmlStr, err := parseTemplate(htmlTemplate, map[string]interface{}{
		"title": title,
	})
	if err != nil {
		log.Fatal(errors.Wrap(err, "h2 html template"))
	}

	plainTextTemplate := `
--------
{{.title}}
--------`
	plainText, err := parseTemplate(plainTextTemplate, map[string]interface{}{
		"title": title,
	})
	if err != nil {
		log.Fatal(errors.Wrap(err, "h2 plain text template"))
	}

	return newPart(htmlStr, plainText)
}
