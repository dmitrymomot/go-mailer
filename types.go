package mailer

// Predefined email types
const (
	WelcomeEmail EmailType = iota + 1
	ConfirmEmail
	ResetPasswordEmail
	InvoiceEmail
	ReceiptEmail
	TrialExpiredEmail
	TrialExpiringEmail
	UserInvitationEmail
	CommentNotificationEmail
)

type (
	// EmailType ...
	EmailType int
)
