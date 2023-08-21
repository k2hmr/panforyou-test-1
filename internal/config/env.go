package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/k2hmr/panforyou-test-1/internal/validator"
)

func GetSpaceId() (string, error) {
	err := loadEnv(".env")
	if err != nil {
		return "", err
	}

	spaceId := os.Getenv("SPACE_ID")
	if spaceId == "" {
		return "", fmt.Errorf("SPACE_IDが設定されていません。")
	}
	if !validator.IsSpaceIdValid(spaceId) {
		return "", fmt.Errorf("SPACE_IDは不正な形式です。")
	}
	return spaceId, nil
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
	if !validator.IsAccessTokenValid(accessToken) {
		return "", fmt.Errorf("ACCESS_TOKENは不正な形式です。")
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
