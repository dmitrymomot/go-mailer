package mailer

type (
	// WelcomeEmailOptions struct
	WelcomeEmailOptions struct {
	}

	// WelcomeEmail struct
	WelcomeEmail struct {
	}
)

// NewWelcomeEmail factory
func NewWelcomeEmail(opt WelcomeEmailOptions) WelcomeEmail {
	return WelcomeEmail{}
}
