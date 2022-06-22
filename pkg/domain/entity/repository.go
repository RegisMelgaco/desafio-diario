package entity

import (
	"context"

	"github.com/google/uuid"
)

type ChallangeReposity interface {
	Create(context.Context, CreateChallangeInput) (Challange, error)
	List(context.Context, ListChallangeFilter) ([]Challange, error)
	Put(context.Context, PutChallangeInput) (Challange, error)
	Delete(context.Context, uuid.UUID) error
}
