package service

import (
	"github.com/AviralDixit-star/banking/domain"
	"github.com/AviralDixit-star/banking/dto"
	"github.com/AviralDixit-star/banking/errs"
)

//Port
type CustomerService interface {
	GetAllCustomers(string) ([]domain.Customer, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
}

//Service Implementation
type DefaultCustomerService struct {
	//dependency
	repo domain.CustomerRepository
}

//receiver func
func (s DefaultCustomerService) GetAllCustomers(status string) ([]domain.Customer, *errs.AppError) {
	if status == "Active" {
		status = "1"
	} else if status == "Inactive" {
		status = "0"
	} else {
		status = ""
	}
	return s.repo.FindAll(status)
}

//receiver func
func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}

	response := c.ToDto()
	return response, nil
}

//helper Func
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
