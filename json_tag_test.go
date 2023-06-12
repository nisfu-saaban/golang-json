package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func TestJsonTag(t *testing.T) {
	product := &Product{
		Id:       "1",
		Name:     "Laptop",
		ImageURL: "http://example.com/image.png",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}
