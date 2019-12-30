package config

import "github.com/kelseyhightower/envconfig"

// Environment - 環境変数の構造体
type Environment struct {
	Port string `envconfig:"PORT" default:"8080"`
}

// LoadEnvironment - 環境変数を読み込み
func LoadEnvironment() (Environment, error) {
	env := Environment{}
	if err := envconfig.Process("", &env); err != nil {
		return env, err
	}

	return env, nil
}
