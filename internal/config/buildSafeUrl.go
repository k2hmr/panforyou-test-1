package config

import (
	"fmt"
	"net/url"
)

func BuildSafeURL(spaceID,entryID, accessToken string) string {
	baseURL := "https://cdn.contentful.com"
	escapedSpaceId := url.QueryEscape(spaceID)
	escapedEntryId := url.QueryEscape(entryID)
	escapedAccessToken := url.QueryEscape(accessToken)

	safeURL := fmt.Sprintf("%s/spaces/%s/entries/%s?access_token=%s", baseURL, escapedSpaceId, escapedEntryId, escapedAccessToken)
	return safeURL
}