package builder

import (
	"log"

	"github.com/pkg/errors"
)

// OptionsItem struct
type OptionsItem struct {
	Title string
	Link  string
	Text  string
}

// Options ...
func Options(items []OptionsItem) Part {
	htmlTemplate := `<table class="attributes" width="100%" cellpadding="0" cellspacing="0" role="presentation"
	style="margin: 0 0 21px;">
	<tr>
	  <td class="attributes_content"
		style="word-break: break-word; font-family: &quot;Nunito Sans&quot;, Helvetica, Arial, sans-serif; font-size: 16px; background-color: #EDF2F7; padding: 16px;"
		bgcolor="#EDF2F7">
		<table width="100%" cellpadding="0" cellspacing="0" role="presentation">
			{{range $item := .items}}
			<tr>
				<td class="attributes_item"
				style="word-break: break-word; font-family: &quot;Nunito Sans&quot;, Helvetica, Arial, sans-serif; font-size: 16px; padding: 0;">
				<div class="f-fallback">
					<strong><a href="{{$item.Link}}" style="color: #4299E1;">{{$item.Title}}</a></strong> - {{$item.Text}}
					<br />
					<br />
				</div>
				</td>
			</tr>
			{{end}}
		</table>
	  </td>
	</tr>
  </table>`
	htmlStr, err := parseTemplate(htmlTemplate, map[string]interface{}{
		"items": items,
	})
	if err != nil {
		log.Fatal(errors.Wrap(err, "options html template"))
	}

	plainTextTemplate := `
{{range $item := .items}}
{{$item.Title}} ( {{$item.Link}} ) - {{$item.Text}}
{{end}}`
	plainText, err := parseTemplate(plainTextTemplate, map[string]interface{}{
		"items": items,
	})
	if err != nil {
		log.Fatal(errors.Wrap(err, "options plain text template"))
	}

	return newPart(htmlStr, plainText)
}
