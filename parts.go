package mailer

type (
	// Part of email body interface
	Part interface {
		HTMLTemplate() string
		PlainTextTemplate() string
	}
)
