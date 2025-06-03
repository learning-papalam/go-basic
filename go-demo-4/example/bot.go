package example

import (
	"encoding/json"
	"os"
)

type People struct {
	Name string `json:"name"`
	Age  int8   `json:"age"`
	Sex  string `json:"-"`
}

func serializer(p *People) []byte {
	data, err := json.Marshal(p)
	if err != nil {
		panic(err.Error())
	}
	return data
}

func save(text []byte) {
	file, err := os.Create("text.json")
	if err != nil {
		panic(err.Error())
	}

	defer file.Close()

	_, err = file.Write(text)
	if err != nil {
		panic(err.Error())
	}
}

func deserializer(p *People, text []byte) {
	err := json.Unmarshal(text, p)
	if err != nil {
		panic(err.Error())
	}
}

func OpenFile() []byte {
	data, err := os.ReadFile("users.json")
	if err != nil {
		panic(err.Error())
	}
	return data
}

type Product struct {
	ProductId int      `json:"productId"`
	Price     *float64 `json:"price,omitempty"`
}

func newProduct() *Product {
	return &Product{
		ProductId: 1,
	}
}
