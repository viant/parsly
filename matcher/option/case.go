package option

// Case represents case options
type Case struct {
	Sensitive bool
	Embed     bool
}

// NewCase creates a case option
func NewCase(sensitive bool) *Case {
	return &Case{Sensitive: sensitive}
}
