package helper

import "github.com/google/uuid"

func NewUUIDv7() (string, error) {
	uuid, err := uuid.NewV7()
	return uuid.String(), err
}
