package apis

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/leesper/holmes"
	"lab02/pkg/apis/api_errors"
	"lab02/pkg/apis/external"
	"lab02/pkg/config"
	"lab02/pkg/utils"
	"net/http"
)

type ResponseWeather struct {
	Location  string                  `json:"location"`
	Region    string                  `json:"region"`
	Country   string                  `json:"country"`
	Latitude  float64                 `json:"latitude"`
	Longitude float64                 `json:"longitude"`
	Timezone  string                  `json:"timezone"`
	LocalTime string                  `json:"local_time"`
	Forecast  []external.ForecastM3O  `json:"forecast"`
	AverageUV float64                 `json:"average_uv"`
	DailyUV   []external.ResultOpenUV `json:"daily_uv"`
}

func (*ResponseWeather) MergeResponses(m3o *external.ResponseM3O, uv *external.ResponseOpenUV, res *ResponseWeather) {

	avg := func(uvArr []external.ResultOpenUV) float64 {
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

func HomePage(context *gin.Context) {
	context.Data(http.StatusOK, config.TextHTML, []byte(config.HtmlPage))
}

func HandleRequest(context *gin.Context) {
	params := context.Request.URL.Query()
	fmt.Println(params)
	city := context.Query(config.QueryCity)
	forecastKey := context.Query(config.QueryForecast)
	indexKey := context.Query(config.QueryIndex)
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
	context.Data(http.StatusOK, config.TextHTML, html)
}

func HandleExternalAPIs(city, forecastKey, indexKey string) (*ResponseWeather, *api_errors.ErrorPage) {
	result := new(ResponseWeather)

	respHandleM3O, err := external.HandleM3O(city, forecastKey)
	if err != nil {
		return nil, err
	}
	logHandleM3O, _ := utils.PrettyStruct(&respHandleM3O)
	fmt.Println(logHandleM3O)

	lat, lng, _ := external.GetLatLng(respHandleM3O)

	responseHandleOpenUV, err := external.HandleOpenUV(lat, lng, indexKey)
	if err != nil {
		return nil, err
	}
	logHandleOpenUV, _ := utils.PrettyStruct(&responseHandleOpenUV)
	fmt.Println(logHandleOpenUV)

	result.MergeResponses(respHandleM3O, responseHandleOpenUV, result)
	logHandleMarshal, _ := utils.PrettyStruct(&responseHandleOpenUV)
	fmt.Println(logHandleMarshal)

	return result, nil
}

func ResponseToHTML(weather *ResponseWeather) ([]byte, error) {
	defer holmes.Start().Stop()

	jsonStr, err := json.MarshalIndent(&weather, "", "\t")
	parsed := append([]byte("<!DOCTYPE html>\n<html>\n\n<body>\n\n<p>Response:</p>\n")[:], jsonStr...)
	parsed = append(parsed, []byte("\n</body>\n</html>")[:]...)
	return parsed, err
}
