package register

import (
	"fmt"
	"net/http"
)

type Registration struct {
	Username string
	Email    string
}

var registrations = []Registration{}

func (o *Registration) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create an order")
}
func (o *Registration) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("List all orders")
}
func (o *Registration) GetByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get an order by ID")
}
func (o *Registration) UpdateByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update an order by ID")
}
func (o *Registration) DeleteByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete an order by ID")
}
