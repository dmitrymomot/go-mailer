package builder

type (
	// Part of template interface
	Part interface {
		HTMLTemplate() string
		PlainTextTemplate() string
	}

	// Part of template
	part struct {
		html      string
		plainText string
	}
)

// NewPart factory
func newPart(html, plainText string) Part {
	return part{html: html, plainText: plainText}
}

// String ...
func (p part) String() string {
	return p.html
}

// HTMLTemplate ...
func (p part) HTMLTemplate() string {
	return p.html
}

// PlainTextTemplate ...
func (p part) PlainTextTemplate() string {
	return p.plainText
}
