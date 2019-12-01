package builder

import (
	"log"

	"github.com/pkg/errors"
)

type (
	// PurchaseTableOptions struct
	PurchaseTableOptions struct {
		ID          string
		Date        string
		Items       map[string]string
		TotalAmount string
		Lang        *PurchaseTableLangOptions
	}

	// PurchaseTableLangOptions struct
	PurchaseTableLangOptions struct {
		Description string
		Amount      string
		Total       string
	}
)

var defaultLangOptions = &PurchaseTableLangOptions{
	Description: "Description",
	Amount:      "Amount",
	Total:       "Total",
}

// PurchaseTable ...
func PurchaseTable(opt PurchaseTableOptions) Part {
	if opt.Lang == nil {
		opt.Lang = defaultLangOptions
	}

	htmlTemplate := `<table class="purchase" width="100%" cellpadding="0" cellspacing="0" role="presentation"
	style="width: 100%; -premailer-width: 100%; -premailer-cellpadding: 0; -premailer-cellspacing: 0; margin: 0; padding: 35px 0;">
	<tr>
	  <td
		style="word-break: break-word; font-family: &quot;Nunito Sans&quot;, Helvetica, Arial, sans-serif; font-size: 16px;">
		<h3
		  style="margin-top: 0; color: #2D3748; font-size: 14px; font-weight: bold; text-align: left;"
		  align="left">{{.opt.ID}}</h3>
	  </td>
	  <td
		style="word-break: break-word; font-family: &quot;Nunito Sans&quot;, Helvetica, Arial, sans-serif; font-size: 16px;">
		<h3 class="align-right"
		  style="margin-top: 0; color: #2D3748; font-size: 14px; font-weight: bold; text-align: right;"
		  align="right">{{.opt.Date}}</h3>
	  </td>
	</tr>
	<tr>
	  <td colspan="2"
		style="word-break: break-word; font-family: &quot;Nunito Sans&quot;, Helvetica, Arial, sans-serif; font-size: 16px;">
		<table class="purchase_content" width="100%" cellpadding="0" cellspacing="0"
		  style="width: 100%; -premailer-width: 100%; -premailer-cellpadding: 0; -premailer-cellspacing: 0; margin: 0; padding: 25px 0 0;">
		  <tr>
			<th class="purchase_heading" align="left"
			  style="font-family: &quot;Nunito Sans&quot;, Helvetica, Arial, sans-serif; font-size: 16px; padding-bottom: 8px; border-bottom-width: 1px; border-bottom-color: #EDF2F7; border-bottom-style: solid;">
			  <p class="f-fallback"
				style="font-size: 12px; line-height: 1.625; color: #718096; margin: 0;">{{.opt.Lang.Description}}
			  </p>
			</th>
			<th class="purchase_heading" align="right"
			  style="font-family: &quot;Nunito Sans&quot;, Helvetica, Arial, sans-serif; font-size: 16px; padding-bottom: 8px; border-bottom-width: 1px; border-bottom-color: #EDF2F7; border-bottom-style: solid;">
			  <p class="f-fallback"
				style="font-size: 12px; line-height: 1.625; color: #718096; margin: 0;">{{.opt.Lang.Amount}}</p>
			</th>
		  </tr>
		  {{range $description, $amount := .opt.Items}}
		  <tr>
			<td width="80%" class="purchase_item"
			  style="word-break: break-word; font-family: &quot;Nunito Sans&quot;, Helvetica, Arial, sans-serif; font-size: 15px; color: #2D3748; line-height: 18px; padding: 10px 0;">
			  <span class="f-fallback">{{$description}}</span></td>
			<td class="align-right" width="20%"
			  style="word-break: break-word; font-family: &quot;Nunito Sans&quot;, Helvetica, Arial, sans-serif; font-size: 16px; text-align: right;"
			  align="right"><span class="f-fallback">{{$amount}}</span></td>
		  </tr>
		  {{end}}
		  <tr>
			<td width="80%" class="purchase_footer" valign="middle"
			  style="word-break: break-word; font-family: &quot;Nunito Sans&quot;, Helvetica, Arial, sans-serif; font-size: 16px; padding-top: 15px; border-top-width: 1px; border-top-color: #EDF2F7; border-top-style: solid;">
			  <p class="f-fallback purchase_total purchase_total--label"
				style="font-size: 16px; line-height: 1.625; text-align: right; font-weight: bold; color: #2D3748; margin: 0; padding: 0 15px 0 0;"
				align="right">{{.opt.Lang.Total}}</p>
			</td>
			<td width="20%" class="purchase_footer" valign="middle"
			  style="word-break: break-word; font-family: &quot;Nunito Sans&quot;, Helvetica, Arial, sans-serif; font-size: 16px; padding-top: 15px; border-top-width: 1px; border-top-color: #EDF2F7; border-top-style: solid;">
			  <p class="f-fallback purchase_total"
				style="font-size: 16px; line-height: 1.625; text-align: right; font-weight: bold; color: #2D3748; margin: 0;"
				align="right">{{.opt.TotalAmount}}</p>
			</td>
		  </tr>
		</table>
	  </td>
	</tr>
  </table>`
	htmlStr, err := parseTemplate(htmlTemplate, map[string]interface{}{
		"opt": opt,
	})
	if err != nil {
		log.Fatal(errors.Wrap(err, "purchase table html template"))
	}

	plainTextTemplate := `
{{.opt.ID}}
--------------

{{.opt.Date}}
--------

{{.opt.Lang.Description}} –– {{.opt.Lang.Amount}}
{{range $description, $amount := .opt.Items}}
* {{$description}} –– {{$amount}}{{end}}

{{.opt.Lang.Total}} –– {{.opt.TotalAmount}}`
	plainText, err := parseTemplate(plainTextTemplate, map[string]interface{}{
		"opt": opt,
	})
	if err != nil {
		log.Fatal(errors.Wrap(err, "purchase table plain text template"))
	}

	return newPart(htmlStr, plainText)
}
