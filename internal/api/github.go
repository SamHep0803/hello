package api

import (
	"net/http"
)

var client = &http.Client{}

func GetGithubNotifications(token string) (int, error) {
	req, err := http.NewRequest("GET", githubUrl+"/notifications", nil)
	if err != nil {
		return 0, err
	}

	results, err := SendRequest(client, req, token)
	if err != nil {
		return 0, err
	}

	m := results.([]map[string]interface{})

	return len(m), nil
}

func GetGithubUsername(token string) (string, error) {
	req, err := http.NewRequest("GET", githubUrl+"/user", nil)
	if err != nil {
		return "", err
	}

	results, err := SendRequest(client, req, token)
	if err != nil {
		return "", err
	}

	m := results.(map[string]interface{})

	return m["login"].(string), nil
}
