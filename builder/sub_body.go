package builder

import (
	"log"

	"github.com/pkg/errors"
)

// SubBody ...
func SubBody(items []string) Part {
	htmlTemplate := `<table class="body-sub" role="presentation"
	style="margin-top: 25px; padding-top: 25px; border-top-width: 1px; border-top-color: #EDF2F7; border-top-style: solid;">
	<tr>
	  <td
		style="word-break: break-word; font-family: &quot;Nunito Sans&quot;, Helvetica, Arial, sans-serif; font-size: 16px;">
		{{range $item := .items}}
		<p class="f-fallback sub"
		  style="font-size: 13px; line-height: 1.625; color: #2D3748; margin: .4em 0 1.1875em;">{{$item}}</p>
		{{end}}
	  </td>
	</tr>
  </table>`
	htmlStr, err := parseTemplate(htmlTemplate, map[string]interface{}{
		"items": items,
	})
	if err != nil {
		log.Fatal(errors.Wrap(err, "sub body html template"))
	}

	plainTextTemplate := `{{range $item := .items}}
{{$item}}{{end}}`
	plainText, err := parseTemplate(plainTextTemplate, map[string]interface{}{
		"items": items,
	})
	if err != nil {
		log.Fatal(errors.Wrap(err, "sub body plain text template"))
	}

	return newPart(htmlStr, plainText)
}
