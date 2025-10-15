package inquiryservice

import (
	"laundry-backend/services"
)

type InquiryService struct {
	service services.UsecaseService
}

func ApiInquiryService(service services.UsecaseService) InquiryService {
	return InquiryService{
		service: service,
	}
}
