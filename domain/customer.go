package domain

import (
	"github.com/AviralDixit-star/banking/dto"
	"github.com/AviralDixit-star/banking/errs"
)

type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string
	Status      string
}

func (c Customer) ToDto() *dto.CustomerResponse {
	response := &dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		Status:      c.Status,
		DateOfBirth: c.DateOfBirth,
	}
	return response
}

type CustomerRepository interface {
	FindAll(string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}
