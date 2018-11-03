package schoold

import (
	"github.com/google/uuid"
)

type guid = uuid.UUID

// GPS for maps.
type GPS struct {
	Latitude  float64
	Longitude float64
}

// MailAddress is a phiycal location
type MailAddress struct {
	Street     string
	Appartment string
	City       string
	State      string
	Zip        string
}

// LetterHead formats the address for a label.
func (a MailAddress) LetterHead() string {
	return a.Street + a.Appartment + ", " + a.City + ", " + a.State + ", " + a.Zip
}

// Location is phyical and mapable.
// locations are cool.
type Location struct {
	ID   guid
	Name string
	MailAddress
	GPS
}

/*
func main() {

	var c Campus

	c.ID = uuid.New()
	c.Name = "Cool School"
	c.Street = "123 cool dr"
	c.Appartment = "456"

	fmt.Println(c)
	fmt.Println(c.LetterHead())

}
*/
