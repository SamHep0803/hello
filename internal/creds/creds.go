package creds

import (
	"errors"
	"fmt"

	"github.com/zalando/go-keyring"
)

const (
	serviceName = "hello"
	githubKey   = "github"
)

func GetCreds() (string, error) {
	githubToken, err := keyring.Get(serviceName, githubKey)
	if err != nil {
		if !errors.Is(err, keyring.ErrNotFound) {
			return "", fmt.Errorf("failed to get github token: %w", err)
		}
		return "", err
	}

	return githubToken, nil
}

func SetCreds(platformKey, platformToken string) error {
	err := keyring.Set(serviceName, platformKey, platformToken)
	if err != nil {
		return err
	}
	return nil
}
