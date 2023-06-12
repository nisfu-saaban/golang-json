package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapJson(t *testing.T) {

	jsonString := `{"id":"1","name":"Laptop","image_url":"http://example.com/image.png"}`
	jsonBytes := []byte(jsonString)

	var product map[string]interface{}
	_ = json.Unmarshal(jsonBytes, &product)
	fmt.Println(product)
}
