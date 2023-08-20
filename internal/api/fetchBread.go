package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/k2hmr/panforyou-test-1/internal/config"
	"github.com/k2hmr/panforyou-test-1/internal/model"
)

func FetchBread(entryId string) (*model.SaveData, error) {
	spaceId, err := config.GetSpaceID()
	if err != nil {
		return nil, err
	}
	accessToken, err := config.GetAccessToken()
	if err != nil {
		return nil, err
	}
	apiUrl := config.BuildSafeURL(spaceId, entryId, accessToken)
	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		return nil, err
	}
	
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("APIリクエストに失敗しました。: %s", resp.Status)
	}

	var resData model.ResData
	err = json.NewDecoder(resp.Body).Decode(&resData)
	if err != nil {
		return nil, err
	}
	saveData := model.SaveData{Id: resData.Sys.ID, Name: resData.Fields.Name, CreatedAt: resData.Sys.CreatedAt}

	return &saveData, nil
}
