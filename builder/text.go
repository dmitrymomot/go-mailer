package builder

import (
	"log"

	"github.com/pkg/errors"
)

// Text ...
func Text(text string) Part {
	htmlTemplate := `<p style="font-size: 16px; line-height: 1.625; color: #2D3748; margin: .4em 0 1.1875em;">{{.text}}</p>`
	htmlStr, err := parseTemplate(htmlTemplate, map[string]interface{}{
		"text": text,
	})
	if err != nil {
		log.Fatal(errors.Wrap(err, "text html template"))
	}

	plainTextTemplate := `
{{.text}}`
	plainText, err := parseTemplate(plainTextTemplate, map[string]interface{}{
		"text": text,
	})
	if err != nil {
		log.Fatal(errors.Wrap(err, "text plain text template"))
	}

	return newPart(htmlStr, plainText)
}
