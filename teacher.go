package school

import (
	"github.com/google/uuid"
	"time"
)

type Teacher struct {
	Person
	uuid     uuid.UUID
	id       string // any type of id
	HireDate time.Time
}
