package external

import (
	"fmt"
	"lab02/pkg/apis/api_errors"
	"lab02/pkg/config"
	"lab02/pkg/utils"
	"net/http"
	"time"
)

type ErrorUV struct {
	Err string `json:"error"`
}

func NewErrorUV(err string) *ErrorUV {
	return &ErrorUV{Err: err}
}

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

func HandleOpenUV(lat, lng *float64, indexKey string) (*ResponseOpenUV, *api_errors.ErrorPage) {
	myClient := &http.Client{Timeout: 10 * time.Second}

	result := new(ResponseOpenUV)
	url := config.UVURL

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Request error: %+v\n", err)
	}

	q := req.URL.Query()
	q.Add(config.Lat, fmt.Sprintf("%v", *lat))
	q.Add(config.Lng, fmt.Sprintf("%v", *lng))
	req.URL.RawQuery = q.Encode()

	req.Header.Set(config.XAccessToken, indexKey)
	resp, err := myClient.Do(req)
	if err != nil {
		fmt.Printf("Response error: %+v\n", err)
	}
	fmt.Printf("Status code: %v\n", resp.StatusCode)
	if resp.StatusCode >= 500 {
		errPage := new(ErrorUV)
		utils.UnmarshalJson(resp.Body, errPage)
		return nil, api_errors.NewErrorPage(http.StatusGone, errPage.Err)
	} else if resp.StatusCode != 200 {
		errPage := new(ErrorUV)
		utils.UnmarshalJson(resp.Body, errPage)
		return nil, api_errors.NewErrorPage(resp.StatusCode, errPage.Err)
	}
	defer resp.Body.Close()

	if utils.UnmarshalJson(resp.Body, result) != nil {
		fmt.Printf("Unmarchall error: %+v\n", err)
	}

	fmt.Println(resp.Body)

	return result, nil
}
