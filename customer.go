package golangjson

import (
	"errors"
)

type Customer struct {
	Id        int
	FirstName string
	LastName  string
	Age       int
	Married   bool
	Hobbies   []string
}

func NewCustomer(id, age int, firstName, lastName string, married bool, hobbies []string) (*Customer, error) {
	customer := &Customer{
		Id:        id,
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
		Married:   married,
		Hobbies:   hobbies,
	}
	if customer != nil {
		return customer, nil
	}
	return nil, errors.New("Customer is Invalid")
}
