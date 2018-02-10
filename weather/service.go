package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Service struct {
	appID string
}

func NewWeatherService(appID string) *Service {
	return &Service{appID: appID}
}

func (s *Service) CurrentWeatherInfo() (Info, error) {
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=London,uk&appid=" + s.appID)
	if err != nil {
		return Info{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return Info{}, fmt.Errorf("unexpected HTTP status %v", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Info{}, err
	}

	var weatherInfo Info
	err = json.Unmarshal(body, &weatherInfo)
	if err != nil {
		return Info{}, err
	}

	return weatherInfo, nil
}
