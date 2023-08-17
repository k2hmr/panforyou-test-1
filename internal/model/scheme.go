package model

import (
	"time"
)


type Fields struct {
	Name        string                 `json:"name"`
}

type SysInfo struct {
	CreatedAt time.Time `json:"createdAt"`
	ID          string                `json:"id"`
}

type ResData struct {
	Fields   Fields `json:"fields"`
	Sys      SysInfo   `json:"sys"`
}

type SaveData struct {
	Id string `json:"id"`
	Name string `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}



