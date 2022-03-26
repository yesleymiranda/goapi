package user

import (
	"context"
	"fmt"

	"github.com/yesleymiranda/go-toolkit/logger"
)

type Repository struct {
	DB string
}

func NewRepository(db string) *Repository {
	return &Repository{
		DB: db,
	}
}

func (r *Repository) GetByID(ctx context.Context, id int64) (*User, error) {
	logger.Info(fmt.Sprintf("user repository id:%v", id))
	return &User{ID: id, Name: "name"}, nil
}
