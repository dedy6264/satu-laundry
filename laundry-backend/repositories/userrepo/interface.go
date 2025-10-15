package userrepo

import (
	"laundry-backend/entities"
)

type UserRepo interface {
	FindByEmail(email string) (entities.User, error)
	Create(user entities.User) error
}
