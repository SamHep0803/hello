package api

import (
	"fmt"
	"net/http"

	"github.com/spf13/viper"
)

var (
	lon, lat float64
)

func GetTemperature(cityName, token string) (int64, error) {
	results, err := SendWeatherRequest(cityName, "/weather", token)
	if err != nil {
		return 0, err
	}

	m := results["main"].(map[string]interface{})

	return int64(m["temp"].(float64)), nil
}

func GetForecast(cityName, token string) (string, error) {
	results, err := SendWeatherRequest(cityName, "/weather", token)
	if err != nil {
		return "", err
	}

	w := results["weather"].([]interface{})[0].(map[string]interface{})
	main := w["main"]
	description := w["description"]

	return fmt.Sprintf("%s: %s", main, description), nil
}

func GetCoordinatesByCity(cityName string, token string) (map[string]float64, error) {
	req, err := http.NewRequest("GET", weatherGeoUrl+"/direct", nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("q", viper.GetString("weather.city"))
	q.Add("appid", token)
	req.URL.RawQuery = q.Encode()

	results, err := SendRequest(client, req, "")
	if err != nil {
		return nil, err
	}

	m := results.([]interface{})[0].(map[string]interface{})

	return map[string]float64{
		"lat": m["lat"].(float64),
		"lon": m["lon"].(float64),
	}, nil
}

func SendWeatherRequest(cityName, endpoint, token string) (map[string]interface{}, error) {
	req, err := http.NewRequest("GET", weatherDataUrl+endpoint, nil)
	if err != nil {
		return nil, err
	}

	coords, err := GetCoordinatesByCity(cityName, token)
	if err != nil {
		return nil, err
	}

	lat = coords["lat"]
	lon = coords["lon"]

	q := req.URL.Query()
	q.Add("lat", fmt.Sprintf("%f", lat))
	q.Add("lon", fmt.Sprintf("%f", lon))
	q.Add("units", viper.GetString("weather.units"))
	q.Add("appid", token)
	req.URL.RawQuery = q.Encode()

	results, err := SendRequest(client, req, "")
	if err != nil {
		return nil, err
	}

	m := results.(map[string]interface{})
	return m, nil

}
