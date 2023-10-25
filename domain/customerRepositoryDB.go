package domain

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/AviralDixit-star/banking/errs"
	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDB struct {
	client *sql.DB
}

func (d CustomerRepositoryDB) FindAll(status string) ([]Customer, *errs.AppError) {
	var row *sql.Rows
	var err error
	if status == "" {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		row, err = d.client.Query(findAllSql)
	} else {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		row, err = d.client.Query(findAllSql, status)
	}

	if err != nil {
		log.Println("Error while executing the querring" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected Database error")
	}

	customers := make([]Customer, 0)
	for row.Next() {
		var c Customer
		err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
		if err != nil {
			return nil, errs.NewNotFoundError("No customer is there in Database")
		}
		customers = append(customers, c)
	}
	return customers, nil

}

func (d CustomerRepositoryDB) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	row := d.client.QueryRow(customerSql, id)
	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("No customer is found")
		} else {
			log.Println("Error while scaaning " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected Database error")
		}

	}
	return &c, nil

}

//helper func
func NewCustomerRepositoryDB() CustomerRepositoryDB {
	db, err := sql.Open("mysql", "root:aviral9956@tcp(localhost:3306)/banking")
	if err != nil {
		fmt.Println("here")
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return CustomerRepositoryDB{client: db}
}
