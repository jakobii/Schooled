package school

import (
	"github.com/google/uuid"
)

type Class struct {
	ID       uuid.UUID
	Teachers []Teacher
}
