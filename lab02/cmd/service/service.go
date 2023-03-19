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
	"github.com/leesper/holmes"
)

var htmlPage = "<!DOCTYPE html>\n<html>\n<body>\n\n<p>Show information about city</p>\n\n" +
	"<form action=\"http://localhost:8080/weather\" id=\"frm1\" method=\"get\">\n" +
	"    City: <input name=\"city\" type=\"text\"><br><br>\n    API Forecast: <input name=\"forecast\" type=\"text\"><br><br>\n" +
	"    API UX Index: <input name=\"index\" type=\"text\"><br><br>\n" +
	"    <input onclick=\"myFunction()\" type=\"button\" value=\"Submit\">\n</form>\n\n<script>\n" +
	"    function myFunction() {\n        document.getElementById(\"frm1\").submit();\n" +
	"    }\n</script>\n\n</body>\n</html>"

type ErrorPage struct {
	StatusCode int    `json:"status_code"`
	Err        string `json:"message"`
}

func NewErrorPage(status int, err string) *ErrorPage {
	return &ErrorPage{
		StatusCode: status,
		Err:        err,
	}
}

type ErrorM3O struct {
	Id     string `json:"id"`
	Code   int    `json:"code"`
	Detail string `json:"detail"`
	Status string `json:"status"`
}

func NewErrorM3O(id string, code int, detail, status string) *ErrorM3O {
	return &ErrorM3O{
		Id:     id,
		Code:   code,
		Detail: detail,
		Status: status,
	}
}

type ErrorUV struct {
	Err string `json:"error"`
}

func NewErrorUV(err string) *ErrorUV {
	return &ErrorUV{Err: err}
}

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

	router.GET("/weather", HandleRequest)
	router.GET("/", HomePage)

	router.Run("localhost:8080")
}

func HomePage(context *gin.Context) {
	context.Data(http.StatusOK, "text/html; charset=utf-8", []byte(htmlPage))
}

func HandleRequest(context *gin.Context) {
	params := context.Request.URL.Query()
	fmt.Println(params)
	city := context.Query("city")
	forecastKey := context.Query("forecast")
	indexKey := context.Query("index")
	fmt.Println(city + " " + forecastKey + " " + indexKey)
	resp, errorPage := HandleExternalAPIs(city, forecastKey, indexKey)

	if errorPage != nil {
		context.IndentedJSON(errorPage.StatusCode, errorPage)
		return
	}
	html, err := ResponseToHTML(resp)
	if err != nil {
		fmt.Printf("parsing html error")
	}
	context.Data(http.StatusOK, "text/html; charset=utf-8", html)
}

func HandleExternalAPIs(city, forecastKey, indexKey string) (*ResponseWeather, *ErrorPage) {
	result := new(ResponseWeather)

	respHandleM3O, err := HandleM3O(city, forecastKey)
	if err != nil {
		return nil, err
	}
	logHandleM3O, _ := PrettyStruct(&respHandleM3O)
	fmt.Println(logHandleM3O)

	lat, lng, _ := GetLatLng(respHandleM3O)

	responseHandleOpenUV, err := HandleOpenUV(lat, lng, indexKey)
	if err != nil {
		return nil, err
	}
	logHandleOpenUV, _ := PrettyStruct(&responseHandleOpenUV)
	fmt.Println(logHandleOpenUV)

	result.MergeResponses(respHandleM3O, responseHandleOpenUV, result)
	logHandleMarshal, _ := PrettyStruct(&responseHandleOpenUV)
	fmt.Println(logHandleMarshal)

	return result, nil
}

func HandleM3O(city, forecastKey string) (*ResponseM3O, *ErrorPage) {
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
	if resp.StatusCode >= 500 {
		errPage := new(ErrorM3O)
		UnmarshalJson(resp.Body, errPage)
		return nil, NewErrorPage(http.StatusGone, errPage.Detail)
	} else if resp.StatusCode != 200 {
		errPage := new(ErrorM3O)
		UnmarshalJson(resp.Body, errPage)
		return nil, NewErrorPage(resp.StatusCode, errPage.Detail)
	}
	defer resp.Body.Close()

	if UnmarshalJson(resp.Body, result) != nil {
		fmt.Printf("Unmarchall error: %+v\n", err)
	}

	fmt.Println(resp.Body)

	return result, nil
}

func HandleOpenUV(lat, lng *float64, indexKey string) (*ResponseOpenUV, *ErrorPage) {
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
	if resp.StatusCode >= 500 {
		errPage := new(ErrorM3O)
		UnmarshalJson(resp.Body, errPage)
		return nil, NewErrorPage(http.StatusGone, errPage.Detail)
	} else if resp.StatusCode != 200 {
		errPage := new(ErrorUV)
		UnmarshalJson(resp.Body, errPage)
		return nil, NewErrorPage(resp.StatusCode, errPage.Err)
	}
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

func ResponseToHTML(weather *ResponseWeather) ([]byte, error) {
	defer holmes.Start().Stop()

	jsonStr, err := json.MarshalIndent(&weather, "", "\t")
	parsed := append([]byte("<!DOCTYPE html>\n<html>\n\n<body>\n\n<p>Response:</p>\n")[:], jsonStr...)
	parsed = append(parsed, []byte("\n</body>\n</html>")[:]...)
	return parsed, err
}
