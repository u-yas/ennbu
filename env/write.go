package env

import (
	"github.com/joho/godotenv"
)

func WriteEnv(content map[string]string, envFilePath string) error {
	err := godotenv.Write(content, envFilePath)
	if err != nil {
		return err
	}

	return nil
}
