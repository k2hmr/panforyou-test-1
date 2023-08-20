package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetSpaceID() (string, error) {
	err := loadEnv(".env")
	if err != nil {
		return "", err
	}

	spaceID := os.Getenv("SPACE_ID")
	if spaceID == "" {
		return "", fmt.Errorf("SPACE_IDが設定されていません。")
	}
	return spaceID, nil
}

func GetAccessToken() (string, error) {
	err := loadEnv(".env")
	if err != nil {
		return "", err
	}

	accessToken := os.Getenv("ACCESS_TOKEN")
	if accessToken == "" {
		return "", fmt.Errorf("ACCESS_TOKENが設定されていません。")
	}
	return accessToken, nil
}

func loadEnv(key string) error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}
	return nil
}
