package creds

import (
	"errors"
	"fmt"

	"github.com/zalando/go-keyring"
)

const (
	serviceName = "hello"
	githubKey   = "github"
	weatherKey  = "weather"
)

func GetCreds() (map[string]string, error) {
	githubToken, err := keyring.Get(serviceName, githubKey)
	if err != nil {
		if !errors.Is(err, keyring.ErrNotFound) {
			return map[string]string{}, fmt.Errorf("failed to get github token: %w", err)
		}
		return map[string]string{}, err
	}

	weatherToken, err := keyring.Get(serviceName, weatherKey)
	if err != nil {
		if !errors.Is(err, keyring.ErrNotFound) {
			return map[string]string{}, fmt.Errorf("failed to get weather token: %w", err)
		}
		return map[string]string{}, err
	}

	return map[string]string{"github": githubToken, "weather": weatherToken}, nil
}

func GetGithubToken() (string, error) {
	githubToken, err := keyring.Get(serviceName, githubKey)
	if err != nil {
		if !errors.Is(err, keyring.ErrNotFound) {
			return "", fmt.Errorf("failed to get github token: %w", err)
		}
		return "", err
	}

	return githubToken, nil
}

func GetWeatherToken() (string, error) {
	weatherToken, err := keyring.Get(serviceName, weatherKey)
	if err != nil {
		if !errors.Is(err, keyring.ErrNotFound) {
			return "", fmt.Errorf("failed to get weather token: %w", err)
		}
		return "", err
	}

	return weatherToken, nil
}

func SetCreds(platformKey, platformToken string) error {
	err := keyring.Set(serviceName, platformKey, platformToken)
	if err != nil {
		return err
	}
	return nil
}
