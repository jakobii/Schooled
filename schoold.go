package school

// Person is a simple interface that all peoply types should implement.
type Personal interface {
	ContactInfo() Contact
}

// Contact stores standard contact information.
type Contact struct {
	FirstName    string
	LastName     string
	MiddleName   string
	EmailAddress string
	PhoneNumber  string
}

type Person struct {
	FirstName    string
	LastName     string
	MiddleName   string
	EmailAddress []string
	PhoneNumber  []string
}
