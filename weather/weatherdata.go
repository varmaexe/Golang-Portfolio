package weather

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
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
	apiConfig, err := loadApiConfig(".apiConfig")
	if err != nil {
		return weatherData{}, err
	}

	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=" + apiConfig.OpenWeatherMapApiKey + "&q=" + city)
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

func GetWeatherDataFromApi(c *gin.Context) {
	// city := strings.SplitN(r.URL.Path, "/", 3)[2]
	city := c.Request.FormValue("city")
	data, err := query(city)
	if err != nil {
		c.AbortWithError(c.Writer.Status(), err)
		return
	}
	// c.Header().Set("Content-Type", "application/json; charset=utf-8")
	c.Bind(data)
	// c.JSON(200, gin.H{
	// 	"message": data,
	// })
	c.HTML(200, "weatherhome.html", data)
}
