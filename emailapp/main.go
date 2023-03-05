package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

type apiConfigData struct {
	OpenWeatherMapApiKey string `json:"OpenWeatherMapApiKey`
}

type weatherData struct {
	Name string `json:"name`
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json:"main"`
}

func loadApiConfig(filename string) (apiConfigData, error) {
	bytes, err := ioutil.ReadFile(filename)

	if err != nil {
		return apiConfigData{}, err
	}

	var c apiConfigData

	err = json.Unmarshal(bytes, &c)
	if err != nil {
		return apiConfigData{}, err
	}
	return c, nil
}
func query(city string) (weatherData, error) {
	apiConfig := os.Getenv("API_KEY")
	// if err != nil {
	// 	return weatherData{}, err
	// }

	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=" + apiConfig + "&q=" + city)
	if err != nil {
		return weatherData{}, err
	}

	defer resp.Body.Close()

	var d weatherData
	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return weatherData{}, err
	}
	return d, nil
}

var tmpl *template.Template

func init() {
	tmpl, _ = template.ParseFiles("/templates/*.html")
}
func GetData() {
	http.HandleFunc("/weather/",
		func(w http.ResponseWriter, r *http.Request) {
			city := r.FormValue("city")
			data, err := query(city)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			tmpl.ExecuteTemplate(w, "test.html", data)
		})

	http.ListenAndServe(":8080", nil)
}
