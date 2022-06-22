package usecases

import "local/desafio-diario/pkg/domain/entity"

type usecases struct {
	challangeRepo entity.ChallangeReposity
}

func New(challangeRepo entity.ChallangeReposity) entity.Usecases {
	return usecases{challangeRepo}
}
