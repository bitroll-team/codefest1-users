package model

import "github.com/google/uuid"

type UserSearchResult struct {
	Username string
	UUID     uuid.UUID
}

type ResLogin struct {
	UserId   string `json:"userid"`
	Username string `json:"username"`
	Role     string `json:"role"`
}
