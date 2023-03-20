package external

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"lab02/pkg/apis/api_errors"
	"lab02/pkg/config"
	"lab02/pkg/utils"
	"net/http"
	"time"
)

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

func HandleM3O(city, forecastKey string) (*ResponseM3O, *api_errors.ErrorPage) {
	myClient := &http.Client{Timeout: 10 * time.Second}

	result := new(ResponseM3O)
	url := config.M3OURL
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
	req.Header.Set(config.ContentType, config.ContentApplicationJson)
	req.Header.Set(config.Authorization, fmt.Sprintf("Bearer %s", forecastKey))
	resp, err := myClient.Do(req)
	if err != nil {
		fmt.Printf("Response error: %+v\n", err)
	}
	fmt.Printf("Status code: %v\n", resp.StatusCode)
	if resp.StatusCode >= 500 {
		errPage := new(ErrorM3O)
		utils.UnmarshalJson(resp.Body, errPage)
		return nil, api_errors.NewErrorPage(http.StatusGone, errPage.Detail)
	} else if resp.StatusCode != 200 {
		errPage := new(ErrorM3O)
		utils.UnmarshalJson(resp.Body, errPage)
		return nil, api_errors.NewErrorPage(resp.StatusCode, errPage.Detail)
	}
	defer resp.Body.Close()

	if utils.UnmarshalJson(resp.Body, result) != nil {
		fmt.Printf("Unmarchall error: %+v\n", err)
	}

	fmt.Println(resp.Body)

	return result, nil
}

func GetLatLng(resp *ResponseM3O) (*float64, *float64, error) {
	lat := &resp.Latitude
	lng := &resp.Longitude
	if lat == nil || lng == nil {
		err := errors.New(fmt.Sprintf("Invalid parameters to create next query' lat=%v, lng=%v", lat, lng))
		return nil, nil, err
	}
	return lat, lng, nil
}
