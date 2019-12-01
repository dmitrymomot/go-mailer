package builder

import (
	"log"

	"github.com/pkg/errors"
)

// Dictionary ...
func Dictionary(d map[string]interface{}) Part {
	htmlTemplate := `<table class="attributes" width="100%" cellpadding="0" cellspacing="0" role="presentation"
	style="margin: 0 0 21px;">
		<tr>
			<td class="attributes_container"
			style="word-break: break-word; font-family: &quot;Nunito Sans&quot;, Helvetica, Arial, sans-serif; font-size: 16px;">
			<table width="100%" cellpadding="0" cellspacing="0" role="presentation">
				{{range $key, $value := .dict}}
				<tr>
					<td class="attributes_item"
					style="word-break: break-word; font-family: &quot;Nunito Sans&quot;, Helvetica, Arial, sans-serif; font-size: 16px; padding: 0;">
						<strong>{{$key}}:</strong> {{$value}}
					</td>
				</tr>
				{{end}}
			</table>
			</td>
		</tr>
	</table>`
	htmlStr, err := parseTemplate(htmlTemplate, map[string]interface{}{
		"dict": d,
	})
	if err != nil {
		log.Fatal(errors.Wrap(err, "dictionary html template"))
	}

	plainTextTemplate := `
{{range $key, $value := .dict}}
{{$key}}: {{$value}}{{end}}
`
	plainText, err := parseTemplate(plainTextTemplate, map[string]interface{}{
		"dict": d,
	})
	if err != nil {
		log.Fatal(errors.Wrap(err, "dictionary plain text template"))
	}

	return newPart(htmlStr, plainText)
}
