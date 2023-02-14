package api

import (
	"fmt"
	"net/http"

	"github.com/spf13/viper"
)

func GetTemperature(lon, lat float64, token string) (int64, error) {
	req, err := http.NewRequest("GET", weatherDataUrl+"/weather", nil)
	if err != nil {
		return 0, err
	}

	q := req.URL.Query()
	q.Add("lat", fmt.Sprintf("%f", lat))
	q.Add("lon", fmt.Sprintf("%f", lon))
	q.Add("units", viper.GetString("weather.units"))
	q.Add("appid", token)
	req.URL.RawQuery = q.Encode()

	results, err := SendRequest(client, req, "")
	if err != nil {
		return 0, err
	}

	m := results.(map[string]interface{})
	d := m["main"].(map[string]interface{})

	return int64(d["temp"].(float64)), nil
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
