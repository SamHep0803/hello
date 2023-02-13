package creds

import (
	"errors"

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
			return map[string]string{}, err
		}
		return map[string]string{}, err
	}

	weatherToken, err := keyring.Get(serviceName, weatherKey)
	if err != nil {
		if !errors.Is(err, keyring.ErrNotFound) {
			return map[string]string{}, err
		}
		return map[string]string{}, err
	}

	return map[string]string{"github": githubToken, "weather": weatherToken}, nil
}

func SetCreds(platformKey, platformToken string) error {
	err := keyring.Set(serviceName, platformKey, platformToken)
	if err != nil {
		return err
	}
	return nil
}
