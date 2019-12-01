package mailer

type (
	// Email struct
	Email struct {
		Product Product
		Body    Body
	}

	// Product struct
	Product struct {
		Name        string
		Link        string
		Logo        string
		Copyright   string
		TroubleText string
	}

	// Body struct
	Body struct {
	}
)
