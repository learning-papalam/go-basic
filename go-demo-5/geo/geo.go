package geo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Geo struct {
	City string `json:"city"`
}

type cityPopulationResponse struct {
	Error bool `json:"error"`
}

func GetMyLocation(city string) (*Geo, error) {
	if city != "" {
		isCity := checkCity(city)
		if !isCity {
			panic("Такого города нет")
		}
		return &Geo{
			City: city,
		}, nil
	}

	r, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		return nil, err
	}

	defer r.Body.Close()

	if r.StatusCode != 200 {
		fmt.Println(r.StatusCode)
		return nil, errors.New("NOT200")
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	var geo Geo
	json.Unmarshal(body, &geo)
	return &geo, nil
}

func checkCity(city string) bool {
	postBody, _ := json.Marshal(map[string]string{
		"city": city,
	})

	r, err := http.Post("https://countriesnow.space/api/v0.1/countries/population/cities", "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		panic("City not found")
	}
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return false
	}

	var populationResponse cityPopulationResponse
	json.Unmarshal(body, &populationResponse)
	return !populationResponse.Error
}
