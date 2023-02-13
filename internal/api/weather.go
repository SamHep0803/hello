package api

import "net/http"

const (
	apiUrl = ""
)

func GetTemperature(token string) {
	req, err := http.NewRequest("GET", apiUrl+"/notifications", nil)
	if err != nil {
		return 0, err
	}

	results, err := SendRequest(client, req, token)
	if err != nil {
		return 0, err
	}

	return len(results), nil

}
