package utils

import (
	"github.com/google/uuid"
)

func GenerateUUID() string {
	newUUID := uuid.New()
	return newUUID.String()
}
