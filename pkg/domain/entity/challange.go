package entity

import (
	"time"

	"github.com/google/uuid"
)

type Challange struct {
	ExternalID  uuid.UUID
	ID          int
	Name        string
	Description string
	Owner       Email
	MaxPoints   int
	StartsAt    time.Time
	EndsAt      time.Time
	Challangers []Email
}
