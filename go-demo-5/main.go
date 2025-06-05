package main

import (
	"demo/weather/geo"
	"demo/weather/weather"
	"flag"
	"fmt"
)

func main() {
	fmt.Println("Новый проект")
	city := flag.String("city", "", "Город")
	format := flag.Int("format", 1, "Формат вывода погоды")

	flag.Parse()

	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		panic(err.Error())
	}

	weatherString := weather.GetWeather(geoData, *format)
	fmt.Println(weatherString)
}
