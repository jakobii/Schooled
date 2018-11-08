package school

// Person is a simple interface that all peoply types should implement.
type Contacter interface {
	Contact() TMI
}


// TMI (To Much Information) stores standard contact information.
type TMI struct {
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
