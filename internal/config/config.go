package config

import (
	"errors"
	"os"
)

const (
	ServerURLEnv    = "SERVER_URL"
	ApiUrlEnv       = "API_URL"
	ApiAuthKeyEnv   = "WS_AUTH_TOKEN"
	PasswordFileEnv = "PASSWORD_FILE"
)

var (
	ErrInvalidServerUrl    = errors.New("invalid server url")
	ErrInvalidAPIUrl       = errors.New("invalid api url")
	ErrInvalidAPIAuthToken = errors.New("invalid api auth token")
	ErrMissingPasswordFile = errors.New("missing password file path")
)

type Config struct {
	serverURL    string
	apiUrl       string
	apiAuthKey   string
	passwordFile string
}

func (c Config) PasswordFile() string {
	return c.passwordFile
}

func (c Config) ServerURL() string {
	return c.serverURL
}

func (c Config) ApiUrl() string {
	return c.apiUrl
}

func (c Config) ApiAuthKey() string {
	return c.apiAuthKey
}

func New() (*Config, error) {
	url := os.Getenv(ServerURLEnv)
	if url == "" {
		return nil, ErrInvalidServerUrl
	}
	apiUrl := os.Getenv(ApiUrlEnv)
	if apiUrl == "" {
		return nil, ErrInvalidAPIUrl
	}
	token := os.Getenv(ApiAuthKeyEnv)
	if token == "" {
		return nil, ErrInvalidAPIAuthToken
	}
	pwdFile := os.Getenv(PasswordFileEnv)
	if pwdFile == "" {
		return nil, ErrMissingPasswordFile
	}
	return &Config{
		serverURL:    url,
		apiUrl:       apiUrl,
		apiAuthKey:   token,
		passwordFile: pwdFile,
	}, nil
}
