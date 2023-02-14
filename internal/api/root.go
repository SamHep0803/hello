package api

import (
	"encoding/json"
	"io"
	"net/http"
)

func SendRequest(client *http.Client, req *http.Request, token string) (interface{}, error) {
	req.Header.Add("Accept", "application/json")
	if token != "" {
		req.Header.Add("Authorization", "Token "+token)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var m interface{}
	json.Unmarshal(body, &m)

	return m, nil

}
