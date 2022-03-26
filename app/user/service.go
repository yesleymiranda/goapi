package user

import (
	"context"

	"github.com/yesleymiranda/go-toolkit/logger"
)

type Service struct {
	respository *Repository
}

func NewService(respository *Repository) *Service {
	return &Service{
		respository: respository,
	}
}

func (r *Service) GetByID(ctx context.Context, id int64) (*User, error) {
	res, err := r.respository.GetByID(ctx, id)
	if err != nil {
		logger.Error("error on user service")
		return nil, err
	}
	return res, nil
}
