package app

import (
	"log"
	"net/http"

	"github.com/AviralDixit-star/banking/domain"
	"github.com/AviralDixit-star/banking/service"
	"github.com/gorilla/mux"
)

func Start() {

	//mux := http.NewServeMux()
	router := mux.NewRouter()

	//wiring
	//ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDB())}

	//Define routes
	router.HandleFunc("/customers", ch.GetAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customer/{customer_id:[0-9]+}", ch.GetCustomerById).Methods(http.MethodGet)
	// router.HandleFunc("/customers", CreateCustomer).Methods(http.MethodPost)
	// router.HandleFunc("/customers/{customer_id}", getCustomer).Methods(http.MethodGet)

	err := http.ListenAndServe("localhost:8000", router)
	if err != nil {
		log.Fatal(err)
	}
}
