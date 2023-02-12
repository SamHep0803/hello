package api

import (
	"net/http"
)

const (
	apiUrl = "https://api.github.com"
)

var client = &http.Client{}

func GetGithubNotifications(token string) (int, error) {
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
