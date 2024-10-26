package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

const api_key = "65b6b00cd98945309ae61314242410"
const api_url = "https://api.weatherapi.com/v1/forecast.json?key=" + api_key //65b6b00cd98945309ae61314242410&q=London&days=1&aqi=no&alerts=no

type WeatherResponse struct {
	Location Location
	Current  Current
}

type Location struct {
	Name    string  `json:"name"`
	Region  string  `json:"region"`
	Country string  `json:"country"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
}

type Current struct {
	Temp_C float32
	Temp_F float32
}

func GetWeather(city string, wg *sync.WaitGroup) {
//func GetWeather(city string) {
   defer wg.Done()	
   days := 1
   aqi := "no"
   alerts := "no"
   // convert days to string 
   daysStr := strconv.Itoa(days)

   url := api_url + "&q=" + city + "&days=" + daysStr + "&aqi=" + aqi + "&alerts=" + alerts	
   
   // Call the http response 
   resp, err := http.Get(url)
   if err != nil {
	 log.Printf("Error fetching weather data for %s: %v\n", city, err)
	 return 
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
    log.Printf("Error: received non-200 response for %s: %v\n", city, resp.Status)
	return 
   }

   var weather WeatherResponse 
   if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
      log.Printf("Error decoding response for %s: %v\n", city, err)
	  return
   }
   // Print the weather report in a table format 
   
   fmt.Printf("| %-30s | %s\n", "City", weather.Location.Name)  
   fmt.Printf("| %-30s | %s\n", "Country", weather.Location.Country) 
   fmt.Printf("| %-30s | %s\n", "Region", weather.Location.Region) 
   fmt.Printf("| %-30s | %.4f\n", "Latitude", weather.Location.Lat) 
   fmt.Printf("| %-30s | %.4f\n", "Longitude", weather.Location.Lon) 
   fmt.Printf("| %-30s | %.2f °C\n", "Temperature", weather.Current.Temp_C) 
   fmt.Printf("| %-30s | %.2f °F\n", "Temperature", weather.Current.Temp_F) 
   fmt.Println("=======================================")
}

func main() {
   fmt.Println("=======================================") 
   fmt.Println("| WEATHER REPORT |") 
   fmt.Println("=======================================")
   cities := []string{"London", "Mumbai", "Delhi", "Bangalore", "Shimla"}
   now := time.Now()	
   var wg sync.WaitGroup
	for _, city := range cities {
		wg.Add(1)
		go GetWeather(city, &wg)
		//GetWeather(city)
	}
	wg.Wait()
	fmt.Println("Total time elapsed:",time.Since(now))
	
	
}