package entity

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Usecases interface {
	CreateChallange(ctx context.Context, input CreateChallangeInput) (Challange, error)
	ListChallanges(ctx context.Context, filter ListChallangeFilter) ([]Challange, error)
	DeleteChallange(ctx context.Context, id uuid.UUID) error
	PutChallange(ctx context.Context, input PutChallangeInput) (Challange, error)
}

type CreateChallangeInput struct {
	Name        string
	Description string
	Owner       Email
	MaxPoints   int
	StartsAt    time.Time
	EndsAt      time.Time
	Challangers []Email
}

type ListChallangeFilter struct {
	Email
}

type PutChallangeInput struct {
	Name        *string
	Description *string
	Owner       *Email
	MaxPoints   *int
	StartsAt    *time.Time
	EndsAt      *time.Time
	Challangers []Email
}
