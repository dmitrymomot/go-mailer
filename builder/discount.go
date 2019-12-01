package builder

import (
	"log"

	"github.com/pkg/errors"
)

// Discount ...
func Discount(title, text, btnLink, btnTitle string) Part {
	htmlTemplate := `<table class="discount" align="center" width="100%" cellpadding="0" cellspacing="0"
	role="presentation"
	style="width: 100%; margin: 0; padding: 24px; border: 2px dashed #CBD5E0; -premailer-width: 100%; -premailer-cellpadding: 0; -premailer-cellspacing: 0; background-color: #EDF2F7;"
	bgcolor="#EDF2F7">
	<tr>
	  <td align="center"
		style="word-break: break-word; font-family: &quot;Nunito Sans&quot;, Helvetica, Arial, sans-serif; font-size: 16px;">
		<h1 class="f-fallback discount_heading"
		  style="margin-top: 0; color: #2D3748; font-size: 22px; font-weight: bold; text-align: center;"
		  align="center">{{.title}}</h1>
		<p class="f-fallback discount_body"
		  style="font-size: 15px; line-height: 1.625; text-align: center; color: #2D3748; margin: .4em 0 1.1875em;"
		  align="center">{{.text}}</p>
		<!-- Border based button
https://litmus.com/blog/a-guide-to-bulletproof-buttons-in-email-design -->
		<table width="100%" border="0" cellspacing="0" cellpadding="0" role="presentation">
		  <tr>
			<td align="center"
			  style="word-break: break-word; font-family: &quot;Nunito Sans&quot;, Helvetica, Arial, sans-serif; font-size: 16px;">
			  <a href="{{.btnLink}}" class="f-fallback button button--green" target="_blank"
				style="color: #FFF; border-color: #48BB78; border-style: solid; border-width: 10px 18px; background-color: #48BB78; display: inline-block; text-decoration: none; border-radius: 3px; box-shadow: 0 2px 3px rgba(0, 0, 0, 0.16); -webkit-text-size-adjust: none; box-sizing: border-box;">{{.btnTitle}}</a>
			</td>
		  </tr>
		</table>
	  </td>
	</tr>
  </table>`
	htmlStr, err := parseTemplate(htmlTemplate, map[string]interface{}{
		"title":    title,
		"text":     text,
		"btnLink":  btnLink,
		"btnTitle": btnTitle,
	})
	if err != nil {
		log.Fatal(errors.Wrap(err, "discount html template"))
	}

	plainTextTemplate := `
***************************
{{.title}}
***************************

{{.text}}

{{.btnTitle}} ( {{.btnLink}} )`
	plainText, err := parseTemplate(plainTextTemplate, map[string]interface{}{
		"title":    title,
		"text":     text,
		"btnLink":  btnLink,
		"btnTitle": btnTitle,
	})
	if err != nil {
		log.Fatal(errors.Wrap(err, "discount plain text template"))
	}

	return newPart(htmlStr, plainText)
}
