package school

// Person is a simple interface that all peoply types should implement.
type Person interface {
	ContactInfo() Contact
}

// Contact stores standard contact information.
type Contact struct {
	Firstname    string
	Lastname     string
	EmailAddress string
	PhoneNumber  string
}
