package administrationservice

import (
	"laundry-backend/services"
)

type AuthService struct {
	service services.UsecaseService
}

func ApiAuthService(service services.UsecaseService) AuthService {
	return AuthService{
		service: service,
	}
}