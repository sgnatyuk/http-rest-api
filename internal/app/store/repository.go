package store

import (
	"github.com/sgnatyuk/http-rest-api/internal/app/model"
)

type UserRepository interface {
	Create(user *model.User) error
	Find(int) (*model.User, error)
	FindByEmail(string) (*model.User, error)
}
