package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ResponseWeather struct {
	Location  string         `json:"location"`
	Region    string         `json:"region"`
	Country   string         `json:"country"`
	Latitude  float64        `json:"latitude"`
	Longitude float64        `json:"longitude"`
	Timezone  string         `json:"timezone"`
	LocalTime string         `json:"local_time"`
	Forecast  []ForecastM3O  `json:"forecast"`
	AverageUV float64        `json:"average_uv"`
	DailyUV   []ResultOpenUV `json:"daily_uv"`
}

////////////////////// ====================================

type PostForecastM3O struct {
	Days     int    `json:"days"`
	Location string `json:"location"`
}

type ForecastM3O struct {
	Date         string  `json:"date"`
	MaxTempC     float64 `json:"max_temp_c"`
	MaxTempF     float64 `json:"max_temp_f"`
	MinTempC     float64 `json:"min_temp_c"`
	MinTempF     float64 `json:"min_temp_f"`
	AvgTempC     float64 `json:"avg_temp_c"`
	AvgTempF     float64 `json:"avg_temp_f"`
	WillItRain   bool    `json:"will_it_rain"`
	ChanceOfRain int     `json:"chance_of_rain"`
	Condition    string  `json:"condition"`
	IconUrl      string  `json:"icon_url"`
	Sunrise      string  `json:"sunrise"`
	Sunset       string  `json:"sunset"`
	MaxWindMph   float64 `json:"max_wind_mph"`
	MaxWindKph   float64 `json:"max_wind_kph"`
}

type ResponseM3O struct {
	Location  string        `json:"location"`
	Region    string        `json:"region"`
	Country   string        `json:"country"`
	Latitude  float64       `json:"latitude"`
	Longitude float64       `json:"longitude"`
	Timezone  string        `json:"timezone"`
	LocalTime string        `json:"local_time"`
	Forecast  []ForecastM3O `json:"forecast"`
}

////////////////////// ====================================

type ResultOpenUV struct {
	Uv          float64   `json:"uv"`
	UvTime      time.Time `json:"uv_time"`
	SunPosition struct {
		Azimuth  float64 `json:"azimuth"`
		Altitude float64 `json:"altitude"`
	} `json:"sun_position"`
}

type ResponseOpenUV struct {
	Result []ResultOpenUV `json:"result"`
}

////////////////////// ====================================

func main() {
	router := gin.Default()
	router.GET("/weather/:city/:forecast/:index", HandleRequest)

	router.Run("localhost:8080")
}

func HandleRequest(c *gin.Context) {
	city := c.Param("city")
	forecastKey := c.Param("forecast")
	indexKey := c.Param("index")
	fmt.Println(city + " " + forecastKey + " " + indexKey)
	resp, err := HandleExternalAPIs(city, forecastKey, indexKey)

	if err != nil {

	}
	c.IndentedJSON(http.StatusOK, resp)
}

func HandleExternalAPIs(city, forecastKey, indexKey string) (*ResponseWeather, error) {
	result := new(ResponseWeather)

	respHandleM3O, _ := HandleM3O(city, forecastKey)
	logHandleM3O, _ := PrettyStruct(&respHandleM3O)
	fmt.Println(logHandleM3O)

	lat, lng, _ := GetLatLng(respHandleM3O)

	responseHandleOpenUV, _ := HandleOpenUV(lat, lng, indexKey)
	logHandleOpenUV, _ := PrettyStruct(&responseHandleOpenUV)
	fmt.Println(logHandleOpenUV)

	result.MergeResponses(respHandleM3O, responseHandleOpenUV, result)
	logHandleMarshal, _ := PrettyStruct(&responseHandleOpenUV)
	fmt.Println(logHandleMarshal)

	return result, nil
}

func HandleM3O(city, forecastKey string) (*ResponseM3O, error) {
	myClient := &http.Client{Timeout: 10 * time.Second}

	result := new(ResponseM3O)
	url := "https://api.m3o.com/v1/weather/Forecast"
	values := PostForecastM3O{
		Location: city,
		Days:     1,
	}
	data, err := json.Marshal(values)
	if err != nil {
		fmt.Printf("Marchall error: %+v\n", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(data))
	if err != nil {
		fmt.Printf("Request error: %+v\n", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", forecastKey))
	resp, err := myClient.Do(req)
	if err != nil {
		fmt.Printf("Response error: %+v\n", err)
	}
	fmt.Printf("Status code: %v\n", resp.StatusCode)
	defer resp.Body.Close()

	if UnmarshalJson(resp.Body, result) != nil {
		fmt.Printf("Unmarchall error: %+v\n", err)
	}

	fmt.Println(resp.Body)

	return result, nil
}

func HandleOpenUV(lat, lng *float64, indexKey string) (*ResponseOpenUV, error) {
	myClient := &http.Client{Timeout: 10 * time.Second}

	result := new(ResponseOpenUV)
	url := "https://api.openuv.io/api/v1/forecast"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Request error: %+v\n", err)
	}

	q := req.URL.Query()
	q.Add("lat", fmt.Sprintf("%v", *lat))
	q.Add("lng", fmt.Sprintf("%v", *lng))
	req.URL.RawQuery = q.Encode()

	req.Header.Set("x-access-token", "openuv-7c2uarlffc2bth-io")
	resp, err := myClient.Do(req)
	if err != nil {
		fmt.Printf("Response error: %+v\n", err)
	}
	fmt.Printf("Status code: %v\n", resp.StatusCode)
	defer resp.Body.Close()

	if UnmarshalJson(resp.Body, result) != nil {
		fmt.Printf("Unmarchall error: %+v\n", err)
	}

	fmt.Println(resp.Body)

	return result, nil
}

////////////////////// ====================================

func GetLatLng(resp *ResponseM3O) (*float64, *float64, error) {
	lat := &resp.Latitude
	lng := &resp.Longitude
	if lat == nil || lng == nil {
		err := errors.New(fmt.Sprintf("Invalid parameters to create next query' lat=%v, lng=%v", lat, lng))
		return nil, nil, err
	}
	return lat, lng, nil
}

func (*ResponseWeather) MergeResponses(m3o *ResponseM3O, uv *ResponseOpenUV, res *ResponseWeather) {

	avg := func(uvArr []ResultOpenUV) float64 {
		sum := 0.0
		for _, element := range uvArr {
			sum += element.Uv
		}
		return sum / float64(len(uvArr))
	}(uv.Result)

	res.Location = m3o.Location
	res.Region = m3o.Region
	res.Country = m3o.Country
	res.Latitude = m3o.Latitude
	res.Longitude = m3o.Longitude
	res.Timezone = m3o.Timezone
	res.LocalTime = m3o.LocalTime
	res.Forecast = m3o.Forecast
	res.DailyUV = uv.Result
	res.AverageUV = avg

}

func UnmarshalJson(r io.ReadCloser, target interface{}) error {
	return json.NewDecoder(r).Decode(target)
}

func PrettyStruct(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}

	return string(val), nil
}
