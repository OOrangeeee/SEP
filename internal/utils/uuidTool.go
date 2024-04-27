package utils

import "github.com/google/uuid"

type UUIDTool struct {
}

func (uT *UUIDTool) GenerateUUID() string {
	return uuid.New().String()
}
