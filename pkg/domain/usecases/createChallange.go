package usecases

import (
	"context"
	"local/desafio-diario/pkg/domain/entity"
	"local/desafio-diario/pkg/domain/erring"
)

func (u usecases) CreateChallange(ctx context.Context, input entity.CreateChallangeInput) (entity.Challange, error) {
	dErr := erring.New("Usecases.CreateChallange")

	c, err := u.challangeRepo.Create(ctx, input)
	if err != nil {
		return c, dErr.Wrap(err).Err()
	}

	return c, nil
}
