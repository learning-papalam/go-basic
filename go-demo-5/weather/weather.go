package weather

import (
	"demo/weather/geo"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func GetWeather(geo *geo.Geo, format int) string {
	var urlStr string = fmt.Sprint("https://wttr.in/", geo.City)
	baseUrl, err := url.Parse(urlStr)
	if err != nil {
		panic(err.Error())
	}

	params := url.Values{}
	params.Add("format", fmt.Sprint(format))
	baseUrl.RawQuery = params.Encode()

	r, err := http.Get(baseUrl.String())
	if err != nil {
		panic(err.Error())
	}

	defer r.Body.Close()

	if r.StatusCode != 200 {
		fmt.Println("Status code", r.StatusCode)
		return ""
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	return string(body)
}
