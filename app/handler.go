package app

import (
	"encoding/json"
	"net/http"

	"github.com/AviralDixit-star/banking/service"
	"github.com/gorilla/mux"
)

type Customer struct {
	Name    string `json:"fullName"`
	City    string `json:"city"`
	ZipCode string `json:"zipCode"`
}

//Handler Function
// func Greet(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "hello World")
// }

type CustomerHandlers struct {
	service service.CustomerService
}

//Handler Function
func (ch *CustomerHandlers) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	// customers := []Customer{
	// 	{Name: "Deepti", City: "Nagpur", ZipCode: "202311"},
	// 	{Name: "Ashish", City: "Kanpur", ZipCode: "202122"},
	// }
	status := r.URL.Query().Get("status")
	customers, err := ch.service.GetAllCustomers(status)
	if err != nil {
		WriteResponse(w, err.Code, err.AsMessage())
	} else {
		WriteResponse(w, http.StatusOK, customers)
	}
}

//Handler Function
func (ch *CustomerHandlers) GetCustomerById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customersResponse, err := ch.service.GetCustomer(vars["customer_id"])
	if err != nil {
		WriteResponse(w, err.Code, err.AsMessage())
	} else {
		WriteResponse(w, http.StatusOK, customersResponse)
	}

}

func WriteResponse(w http.ResponseWriter, code int, response interface{}) {
	w.Header().Add("content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response)
}

// func getCustomer(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	fmt.Fprint(w, vars["customer_id"])
// }

// func CreateCustomer(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Post request recieved")
// }
