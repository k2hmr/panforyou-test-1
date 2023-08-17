package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/k2hmr/panforyou-test-1/internal/model"
)


func FetchFood(entryID string) (*model.SaveData, error) {
	err := godotenv.Load(".env")
	var spaceID = os.Getenv("SPACE_ID")
	var accessToken = os.Getenv("ACCESS_TOKEN")

	apiUrl := fmt.Sprintf("https://cdn.contentful.com/spaces/%s/entries/%s?access_token=%s", spaceID, entryID, accessToken)
	client := &http.Client{}

	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		return nil, err
	}
	
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %s", resp.Status)
	}

	var resData model.ResData
	err = json.NewDecoder(resp.Body).Decode(&resData)
	if err != nil {
		return nil, err
	}
	saveData := model.SaveData{ Id: resData.Sys.ID, Name: resData.Fields.Name, CreatedAt: resData.Sys.CreatedAt }

	return &saveData, nil
}