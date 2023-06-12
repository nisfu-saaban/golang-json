package golangjson

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestDecode(t *testing.T) {
	reader, _ := os.Open("product.json")
	decoder := json.NewDecoder(reader)

	product := &Product{}
	_ = decoder.Decode(product)
	fmt.Println(product)
}
