package app

import (
	"github.com/lazzytchik/wish/internal/model"
)

type Usecaser interface {
	SaveWish(wish model.Wish) error
	GetWishes() ([]model.Wish, error)
}

type Usecases struct {
}

func (u *Usecases) SaveWish(wish model.Wish) error {
	return nil
}

func (u *Usecases) GetWishes() ([]model.Wish, error) {
	return []model.Wish{}, nil
}
