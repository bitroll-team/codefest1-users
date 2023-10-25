package model

import "github.com/google/uuid"

type UserSearchResult struct {
	Username string
	UUID     uuid.UUID
}
