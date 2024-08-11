package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

func FetchWeather(city string, ch chan<- string, wg *sync.WaitGroup) interface{} {

	var data struct {
		Main struct {
			Temp float64 `json:"temp"`
		} `json:"main"`
	}
	defer wg.Done()

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching weather for %s: %s\n", city, err)
		return data
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Printf("Error decoding weather data for %s: %s\n", city, err)
		return data
	}
	ch <- fmt.Sprintf("this is the data for: %s", city)
	return data

}
