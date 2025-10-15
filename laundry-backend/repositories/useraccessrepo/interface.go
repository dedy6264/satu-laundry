package useraccessrepo

import (
	"laundry-backend/entities"
)

type UserAccessRepo interface {
	Create(access entities.UserAccess) error
	FindByID(id int) (entities.UserAccess, error)
	FindByUsername(username string) (entities.UserAccess, error)
	FindAll() ([]entities.UserAccess, error)
	FindAllWithPagination(request entities.DTRequest[entities.UserAccess]) ([]entities.UserAccess, int, error)
	Update(access entities.UserAccess) error
	UpdatePassword(id int, password string) error
	UpdateLastLogin(id int) error
	Delete(id int) error
	AuthenticateUser(username, password string) (entities.UserAccess, error)
}
