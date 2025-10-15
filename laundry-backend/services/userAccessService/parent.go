package useraccessservice

import (
	"laundry-backend/services"
)

type UserAccessService struct {
	service services.UsecaseService
}

func ApiUserAccessService(service services.UsecaseService) UserAccessService {
	return UserAccessService{
		service: service,
	}
}
