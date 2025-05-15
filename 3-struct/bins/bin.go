package bins

import "time"

type BinList struct {
	Bins []Bin `json:"bins"`
}

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
}

func NewBin() *Bin {
	return &Bin{
		Id:        "1",
		Private:   false,
		CreatedAt: time.Now(),
		Name:      "первый",
	}
}
