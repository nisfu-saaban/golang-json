package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonObject(t *testing.T) {
	customer, _ := NewCustomer(1, 30, "Bambang", "Yuwono", true, []string{"Voly", "Makan"})
	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestDecodeJSON(t *testing.T) {
	jsonstring := `{"Id":1,"FirstName":"Bambang","LastName":"Yuwono","Age":30,"Married":true}`
	jsonBytes := []byte(jsonstring)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
}

func TestDecodeJsonArray(t *testing.T) {
	jsonstring := `{"Id":1,"FirstName":"Bambang","LastName":"Yuwono","Age":30,"Married":true,"Hobbies":["Bola","Selam"]}`
	jsonBytes := []byte(jsonstring)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
}

func TestJsonComplex(t *testing.T) {
	student := &Student{
		StudentID: "1",
		FirstName: "Wahyu",
		LastName:  "Waluyo",
		Subject: []Subject{
			{
				Id:          "1",
				SubjectName: "MTK",
				Duration:    "3 Jam",
			},
			{
				Id:          "2",
				SubjectName: "Penjas",
				Duration:    "Secapeknya",
			},
		},
	}

	bytes, _ := json.Marshal(student)
	fmt.Println(string(bytes))
}

func TestDecodeJsonComplex(t *testing.T) {
	jsonstring := `{"StudentID":"1","FirstName":"Wahyu","LastName":"Waluyo","Subject":[{"Id":"1","SubjectName":"MTK","Duration":"3 Jam"},{"Id":"2","SubjectName":"Penjas","Duration":"Secapeknya"}]}`
	jsonBytes := []byte(jsonstring)

	student := &Student{}
	err := json.Unmarshal(jsonBytes, student)
	if err != nil {
		panic(err)
	}
	fmt.Println(student)
}
