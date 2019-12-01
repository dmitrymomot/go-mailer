package builder

import (
	"github.com/pkg/errors"
)

type (
	// EmailTemplate struct
	EmailTemplate struct {
		html      string
		plainText string
		opt       EmailBuilderOptions
	}

	// EmailBuilderOptions struct
	EmailBuilderOptions struct {
		ProductName string
		ProductLink string
		Logo        string
		Parts       []Part
	}
)

// New builder
func New(opt EmailBuilderOptions) EmailTemplate {
	return EmailTemplate{
		opt: opt,
	}
}

// HTMLTemplate returns rendered html template
func (t EmailTemplate) HTMLTemplate() string {
	return t.html
}

// PlainTextTemplate returns rendered plain text template
func (t EmailTemplate) PlainTextTemplate() string {
	return t.plainText
}

// Build email template
func (t EmailTemplate) Build(parts ...Part) (htmlStr, plainText string, err error) {
	t.opt.Parts = parts

	htmlStr, err = parseTemplate(emailHTMLTemplate, t.opt)
	if err != nil {
		return "", "", errors.Wrap(err, "parse html template")
	}

	plainText, err = parseTemplate(emailPlainTextTemplate, t.opt)
	if err != nil {
		return "", "", errors.Wrap(err, "parse plain text template")
	}

	return htmlStr, plainText, nil
}

var emailHTMLTemplate = `
<!DOCTYPE html
  PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">

<head>
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
  <meta name="x-apple-disable-message-reformatting" />
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
  <title></title>
  <style type="text/css" rel="stylesheet" media="all">
    /* Base ------------------------------ */

    @import url("https://fonts.googleapis.com/css?family=Nunito+Sans:400,700&amp;display=swap");

    body {
      width: 100% !important;
      height: 100%;
      margin: 0;
      -webkit-text-size-adjust: none;
    }

    a {
      color: #4299E1;
    }

    a img {
      border: none;
    }

    td {
      word-break: break-word;
    }

    .preheader {
      display: none !important;
      visibility: hidden;
      mso-hide: all;
      font-size: 1px;
      line-height: 1px;
      max-height: 0;
      max-width: 0;
      opacity: 0;
      overflow: hidden;
    }

    /* Type ------------------------------ */

    body,
    td,
    th {
      font-family: "Nunito Sans", Helvetica, Arial, sans-serif;
    }

    h1 {
      margin-top: 0;
      color: #2D3748;
      font-size: 22px;
      font-weight: bold;
      text-align: left;
    }

    h2 {
      margin-top: 0;
      color: #2D3748;
      font-size: 16px;
      font-weight: bold;
      text-align: left;
    }

    h3 {
      margin-top: 0;
      color: #2D3748;
      font-size: 14px;
      font-weight: bold;
      text-align: left;
    }

    td,
    th {
      font-size: 16px;
    }

    p,
    ul,
    ol,
    blockquote {
      margin: .4em 0 1.1875em;
      font-size: 16px;
      line-height: 1.625;
    }

    p.sub {
      font-size: 13px;
    }

    /* Utilities ------------------------------ */

    .align-right {
      text-align: right;
    }

    .align-left {
      text-align: left;
    }

    .align-center {
      text-align: center;
    }

    /* Buttons ------------------------------ */

    .button {
      background-color: #4299E1;
      border-top: 10px solid #4299E1;
      border-right: 18px solid #4299E1;
      border-bottom: 10px solid #4299E1;
      border-left: 18px solid #4299E1;
      display: inline-block;
      color: #FFF;
      text-decoration: none;
      border-radius: 3px;
      box-shadow: 0 2px 3px rgba(0, 0, 0, 0.16);
      -webkit-text-size-adjust: none;
      box-sizing: border-box;
    }

    .button--green {
      background-color: #48BB78;
      border-top: 10px solid #48BB78;
      border-right: 18px solid #48BB78;
      border-bottom: 10px solid #48BB78;
      border-left: 18px solid #48BB78;
    }

    .button--red {
      background-color: #F56565;
      border-top: 10px solid #F56565;
      border-right: 18px solid #F56565;
      border-bottom: 10px solid #F56565;
      border-left: 18px solid #F56565;
    }

    @media only screen and (max-width: 500px) {
      .button {
        width: 100% !important;
        text-align: center !important;
      }
    }

    /* Attribute list ------------------------------ */

    .attributes {
      margin: 0 0 21px;
    }

    .attributes_content {
      background-color: #EDF2F7;
      padding: 16px;
    }

    .attributes_item {
      padding: 0;
    }

    /* Related Items ------------------------------ */

    .related {
      width: 100%;
      margin: 0;
      padding: 25px 0 0 0;
      -premailer-width: 100%;
      -premailer-cellpadding: 0;
      -premailer-cellspacing: 0;
    }

    .related_item {
      padding: 10px 0;
      color: #CBD5E0;
      font-size: 15px;
      line-height: 18px;
    }

    .related_item-title {
      display: block;
      margin: .5em 0 0;
    }

    .related_item-thumb {
      display: block;
      padding-bottom: 10px;
    }

    .related_heading {
      border-top: 1px solid #CBD5E0;
      text-align: center;
      padding: 25px 0 10px;
    }

    /* Discount Code ------------------------------ */

    .discount {
      width: 100%;
      margin: 0;
      padding: 24px;
      -premailer-width: 100%;
      -premailer-cellpadding: 0;
      -premailer-cellspacing: 0;
      background-color: #EDF2F7;
      border: 2px dashed #CBD5E0;
    }

    .discount_heading {
      text-align: center;
    }

    .discount_body {
      text-align: center;
      font-size: 15px;
    }

    /* Social Icons ------------------------------ */

    .social {
      width: auto;
    }

    .social td {
      padding: 0;
      width: auto;
    }

    .social_icon {
      height: 20px;
      margin: 0 8px 10px 8px;
      padding: 0;
    }

    /* Data table ------------------------------ */

    .purchase {
      width: 100%;
      margin: 0;
      padding: 35px 0;
      -premailer-width: 100%;
      -premailer-cellpadding: 0;
      -premailer-cellspacing: 0;
    }

    .purchase_content {
      width: 100%;
      margin: 0;
      padding: 25px 0 0 0;
      -premailer-width: 100%;
      -premailer-cellpadding: 0;
      -premailer-cellspacing: 0;
    }

    .purchase_item {
      padding: 10px 0;
      color: #2D3748;
      font-size: 15px;
      line-height: 18px;
    }

    .purchase_heading {
      padding-bottom: 8px;
      border-bottom: 1px solid #EDF2F7;
    }

    .purchase_heading p {
      margin: 0;
      color: #718096;
      font-size: 12px;
    }

    .purchase_footer {
      padding-top: 15px;
      border-top: 1px solid #EDF2F7;
    }

    .purchase_total {
      margin: 0;
      text-align: right;
      font-weight: bold;
      color: #2D3748;
    }

    .purchase_total--label {
      padding: 0 15px 0 0;
    }

    body {
      background-color: #EDF2F7;
      color: #2D3748;
    }

    p {
      color: #2D3748;
    }

    .email-wrapper {
      width: 100%;
      margin: 0;
      padding: 0;
      -premailer-width: 100%;
      -premailer-cellpadding: 0;
      -premailer-cellspacing: 0;
      background-color: #EDF2F7;
    }

    .email-content {
      width: 100%;
      margin: 0;
      padding: 0;
      -premailer-width: 100%;
      -premailer-cellpadding: 0;
      -premailer-cellspacing: 0;
    }

    /* Masthead ----------------------- */

    .email-masthead {
      padding: 25px 0;
      text-align: center;
    }

    .email-masthead_logo {
      width: 94px;
    }

    .email-masthead_name {
      font-size: 16px;
      font-weight: bold;
      color: #A0AEC0;
      text-decoration: none;
      text-shadow: 0 1px 0 white;
    }

    /* Body ------------------------------ */

    .email-body {
      width: 100%;
      margin: 0;
      padding: 0;
      -premailer-width: 100%;
      -premailer-cellpadding: 0;
      -premailer-cellspacing: 0;
    }

    .email-body_inner {
      width: 570px;
      margin: 0 auto;
      padding: 0;
      -premailer-width: 570px;
      -premailer-cellpadding: 0;
      -premailer-cellspacing: 0;
      background-color: #FFFFFF;
    }

    .email-footer {
      width: 570px;
      margin: 0 auto;
      padding: 0;
      -premailer-width: 570px;
      -premailer-cellpadding: 0;
      -premailer-cellspacing: 0;
      text-align: center;
    }

    .email-footer p {
      color: #A0AEC0;
    }

    .body-action {
      width: 100%;
      margin: 30px auto;
      padding: 0;
      -premailer-width: 100%;
      -premailer-cellpadding: 0;
      -premailer-cellspacing: 0;
      text-align: center;
    }

    .body-sub {
      margin-top: 25px;
      padding-top: 25px;
      border-top: 1px solid #EDF2F7;
    }

    .content-cell {
      padding: 45px;
    }

    /*Media Queries ------------------------------ */

    @media only screen and (max-width: 600px) {

      .email-body_inner,
      .email-footer {
        width: 100% !important;
      }
    }

    @media (prefers-color-scheme: dark) {

      body,
      .email-body,
      .email-body_inner,
      .email-content,
      .email-wrapper,
      .email-masthead,
      .email-footer {
        background-color: #2D3748 !important;
        color: #FFF !important;
      }

      p,
      ul,
      ol,
      blockquote,
      h1,
      h2,
      h3 {
        color: #FFF !important;
      }

      .attributes_content,
      .discount {
        background-color: #1A202C !important;
      }

      .email-masthead_name {
        text-shadow: none !important;
      }
    }
  </style>
  <!--[if mso]>
    <style type="text/css">
      .f-fallback  {
        font-family: Arial, sans-serif;
      }
    </style>
  <![endif]-->
  <style type="text/css" rel="stylesheet" media="all">
    body {
      width: 100% !important;
      height: 100%;
      margin: 0;
      -webkit-text-size-adjust: none;
    }

    body {
      font-family: "Nunito Sans", Helvetica, Arial, sans-serif;
    }

    body {
      background-color: #EDF2F7;
      color: #2D3748;
    }
  </style>
</head>

<body
  style="width: 100% !important; height: 100%; -webkit-text-size-adjust: none; font-family: &quot;Nunito Sans&quot;, Helvetica, Arial, sans-serif; background-color: #EDF2F7; color: #2D3748; margin: 0;"
  bgcolor="#EDF2F7">
  <span class="preheader"
    style="display: none !important; visibility: hidden; mso-hide: all; font-size: 1px; line-height: 1px; max-height: 0; max-width: 0; opacity: 0; overflow: hidden;">This
    is example text for the preheader set via the YAML front-matter for each email.</span>
  <table class="email-wrapper" width="100%" cellpadding="0" cellspacing="0" role="presentation"
    style="width: 100%; -premailer-width: 100%; -premailer-cellpadding: 0; -premailer-cellspacing: 0; background-color: #EDF2F7; margin: 0; padding: 0;"
    bgcolor="#EDF2F7">
    <tr>
      <td align="center"
        style="word-break: break-word; font-family: &quot;Nunito Sans&quot;, Helvetica, Arial, sans-serif; font-size: 16px;">
        <table class="email-content" width="100%" cellpadding="0" cellspacing="0" role="presentation"
          style="width: 100%; -premailer-width: 100%; -premailer-cellpadding: 0; -premailer-cellspacing: 0; margin: 0; padding: 0;">
          <tr>
            <td class="email-masthead"
              style="word-break: break-word; font-family: &quot;Nunito Sans&quot;, Helvetica, Arial, sans-serif; font-size: 16px; text-align: center; padding: 25px 0;"
              align="center">
              <a href="https://example.com" class="f-fallback email-masthead_name"
                style="color: #A0AEC0; font-size: 16px; font-weight: bold; text-decoration: none; text-shadow: 0 1px 0 white;">
                [Product Name]
              </a>
            </td>
          </tr>
          <!-- Email Body -->
          <tr>
            <td class="email-body" width="570" cellpadding="0" cellspacing="0"
              style="word-break: break-word; margin: 0; padding: 0; font-family: &quot;Nunito Sans&quot;, Helvetica, Arial, sans-serif; font-size: 16px; width: 100%; -premailer-width: 100%; -premailer-cellpadding: 0; -premailer-cellspacing: 0;">
              <table class="email-body_inner" align="center" width="570" cellpadding="0" cellspacing="0"
                role="presentation"
                style="width: 570px; -premailer-width: 570px; -premailer-cellpadding: 0; -premailer-cellspacing: 0; background-color: #FFFFFF; margin: 0 auto; padding: 0;"
                bgcolor="#FFFFFF">
                <!-- Body content -->
                <tr>
                  <td class="content-cell"
                    style="word-break: break-word; font-family: &quot;Nunito Sans&quot;, Helvetica, Arial, sans-serif; font-size: 16px; padding: 45px;">
                    <div class="f-fallback">
					            {{range $part := .Parts}}{{$part.HTMLTemplate}}{{end}}
                    </div>
                  </td>
                </tr>
              </table>
            </td>
          </tr>
          <tr>
            <td
              style="word-break: break-word; font-family: &quot;Nunito Sans&quot;, Helvetica, Arial, sans-serif; font-size: 16px;">
              <table class="email-footer" align="center" width="570" cellpadding="0" cellspacing="0" role="presentation"
                style="width: 570px; -premailer-width: 570px; -premailer-cellpadding: 0; -premailer-cellspacing: 0; text-align: center; margin: 0 auto; padding: 0;">
                <tr>
                  <td class="content-cell" align="center"
                    style="word-break: break-word; font-family: &quot;Nunito Sans&quot;, Helvetica, Arial, sans-serif; font-size: 16px; padding: 45px;">
                    <p class="f-fallback sub align-center"
                      style="font-size: 13px; line-height: 1.625; text-align: center; color: #A0AEC0; margin: .4em 0 1.1875em;"
                      align="center">© 2019 [Product Name]. All rights reserved.</p>
                    <p class="f-fallback sub align-center"
                      style="font-size: 13px; line-height: 1.625; text-align: center; color: #A0AEC0; margin: .4em 0 1.1875em;"
                      align="center">
                      [Company Name, LLC]
                      <br />1234 Street Rd.
                      <br />Suite 1234
                    </p>
                  </td>
                </tr>
              </table>
            </td>
          </tr>
        </table>
      </td>
    </tr>
  </table>
</body>

</html>
`

var emailPlainTextTemplate = `
Thanks for trying out [Product Name]. We’ve pulled together some information and resources to help you get started.

[Product Name] ( https://example.com )

{{range $part := .Parts}}{{$part.PlainTextTemplate}}{{end}}

© 2019 [Product Name]. All rights reserved.

[Company Name, LLC]

1234 Street Rd.

Suite 1234
`
