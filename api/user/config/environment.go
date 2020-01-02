package config

import (
	"golang.org/x/xerrors"

	"github.com/kelseyhightower/envconfig"
)

// Environment - 環境変数の構造体
type Environment struct {
	Port                             string `envconfig:"PORT" default:"8080"`
	GoogleApplicationCredentialsPath string `envconfig:"GOOGLE_APPLICATION_CREDENTIALS_PATH" required:"true"`
}

// LoadEnvironment - 環境変数を読み込み
func LoadEnvironment() (Environment, error) {
	env := Environment{}
	if err := envconfig.Process("", &env); err != nil {
		return env, xerrors.Errorf("Failed to LoadEnvironment: %w", err)
	}

	return env, nil
}
