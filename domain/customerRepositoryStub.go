package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customersMock := []Customer{
		{Id: "123", Name: "Deepti", City: "Nagpur", Zipcode: "202311", DateOfBirth: "2000-02-01", Status: "1"},
		{Id: "645", Name: "Raj", City: "Kanpur", Zipcode: "202311", DateOfBirth: "2000-02-01", Status: "1"},
		{Id: "765", Name: "Hira", City: "Pune", Zipcode: "202311", DateOfBirth: "2000-02-01", Status: "1"},
		{Id: "865", Name: "Dice", City: "Kashmir", Zipcode: "202311", DateOfBirth: "2000-02-01", Status: "1"},
	}
	return CustomerRepositoryStub{customers: customersMock}
}
