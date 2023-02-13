package api

import "net/http"

func GetTemperature(token string) (int, error) {
	req, err := http.NewRequest("GET", weatherUrl+"/notifications", nil)
	if err != nil {
		return 0, err
	}

	results, err := SendRequest(client, req, token)
	if err != nil {
		return 0, err
	}

	return len(results), nil

}
