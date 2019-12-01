package builder

import (
	"log"

	"github.com/pkg/errors"
	"github.com/valyala/fasttemplate"
)

// DefaultButton ...
func DefaultButton(text, link string) Part {
	htmlTemplate := `<table class="body-action" align="center" width="100%" cellpadding="0" cellspacing="0"
	role="presentation"
	style="width: 100%; -premailer-width: 100%; -premailer-cellpadding: 0; -premailer-cellspacing: 0; text-align: center; margin: 30px auto; padding: 0;">
	<tr>
	  <td align="center"
		style="word-break: break-word; font-family: &quot;Nunito Sans&quot;, Helvetica, Arial, sans-serif; font-size: 16px;">
		<table width="100%" border="0" cellspacing="0" cellpadding="0" role="presentation">
		  <tr>
			<td align="center"
			  style="word-break: break-word; font-family: &quot;Nunito Sans&quot;, Helvetica, Arial, sans-serif; font-size: 16px;">
			  <a href="{{.link}}" class="f-fallback button" target="_blank"
				style="color: #FFF; border-color: #4299E1; border-style: solid; border-width: 10px 18px; background-color: #4299E1; display: inline-block; text-decoration: none; border-radius: 3px; box-shadow: 0 2px 3px rgba(0, 0, 0, 0.16); -webkit-text-size-adjust: none; box-sizing: border-box;">{{.text}}</a>
			</td>
		  </tr>
		</table>
	  </td>
	</tr>
  </table>`
	htmlStr, err := parseTemplate(htmlTemplate, map[string]interface{}{
		"text": text,
		"link": link,
	})
	if err != nil {
		log.Fatal(errors.Wrap(err, "button html template"))
	}

	plainTextTemplate := `
{{.text}} ( {{.link}} )
	`
	plainText, err := parseTemplate(plainTextTemplate, map[string]interface{}{
		"text": text,
		"link": link,
	})
	if err != nil {
		log.Fatal(errors.Wrap(err, "button plain text template"))
	}

	return newPart(htmlStr, plainText)
}

// DangerButton ...
func DangerButton(text, link string) Part {
	htmlTemplate := `<table class="body-action" align="center" width="100%" cellpadding="0" cellspacing="0"
	role="presentation"
	style="width: 100%; -premailer-width: 100%; -premailer-cellpadding: 0; -premailer-cellspacing: 0; text-align: center; margin: 30px auto; padding: 0;">
	<tr>
	  <td align="center"
		style="word-break: break-word; font-family: &quot;Nunito Sans&quot;, Helvetica, Arial, sans-serif; font-size: 16px;">
		<table width="100%" border="0" cellspacing="0" cellpadding="0" role="presentation">
		  <tr>
			<td align="center"
			  style="word-break: break-word; font-family: &quot;Nunito Sans&quot;, Helvetica, Arial, sans-serif; font-size: 16px;">
			  <a href="{{link}}" class="f-fallback button button--red" target="_blank"
				style="color: #FFF; border-color: #F56565; border-style: solid; border-width: 10px 18px; background-color: #F56565; display: inline-block; text-decoration: none; border-radius: 3px; box-shadow: 0 2px 3px rgba(0, 0, 0, 0.16); -webkit-text-size-adjust: none; box-sizing: border-box;">{{text}}</a>
			</td>
		  </tr>
		</table>
	  </td>
	</tr>
  </table>`
	th := fasttemplate.New(htmlTemplate, "{{", "}}")
	htmlStr := th.ExecuteString(map[string]interface{}{
		"text": text,
		"link": link,
	})

	plainTextTemplate := `
	{{text}} ( {{link}} )
	`
	tp := fasttemplate.New(plainTextTemplate, "{{", "}}")
	plainText := tp.ExecuteString(map[string]interface{}{
		"text": text,
		"link": link,
	})

	return newPart(htmlStr, plainText)
}

// SuccessButton ...
func SuccessButton(text, link string) Part {
	htmlTemplate := `<table class="body-action" align="center" width="100%" cellpadding="0" cellspacing="0"
	role="presentation"
	style="width: 100%; -premailer-width: 100%; -premailer-cellpadding: 0; -premailer-cellspacing: 0; text-align: center; margin: 30px auto; padding: 0;">
	<tr>
	  <td align="center"
		style="word-break: break-word; font-family: &quot;Nunito Sans&quot;, Helvetica, Arial, sans-serif; font-size: 16px;">
		<table width="100%" border="0" cellspacing="0" cellpadding="0" role="presentation">
		  <tr>
			<td align="center"
			  style="word-break: break-word; font-family: &quot;Nunito Sans&quot;, Helvetica, Arial, sans-serif; font-size: 16px;">
			  <a href="{{link}}" class="f-fallback button button--green" target="_blank"
				style="color: #FFF; border-color: #48BB78; border-style: solid; border-width: 10px 18px; background-color: #48BB78; display: inline-block; text-decoration: none; border-radius: 3px; box-shadow: 0 2px 3px rgba(0, 0, 0, 0.16); -webkit-text-size-adjust: none; box-sizing: border-box;">{{text}}</a>
			</td>
		  </tr>
		</table>
	  </td>
	</tr>
  </table>`
	th := fasttemplate.New(htmlTemplate, "{{", "}}")
	htmlStr := th.ExecuteString(map[string]interface{}{
		"text": text,
		"link": link,
	})

	plainTextTemplate := `
{{text}} ( {{link}} )
	`
	tp := fasttemplate.New(plainTextTemplate, "{{", "}}")
	plainText := tp.ExecuteString(map[string]interface{}{
		"text": text,
		"link": link,
	})

	return newPart(htmlStr, plainText)
}
