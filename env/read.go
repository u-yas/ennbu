package env

import (
	"github.com/joho/godotenv"
)

func ReadEnv(envFilePath string) (map[string]string, error) {
	env, err := godotenv.Read(envFilePath)
	if err != nil {
		return nil, err
	}

	return env, nil
}
