package utils

import "github.com/google/uuid"

func GenerateUUID() (uuid.UUID, error) {
	random, err := uuid.NewRandom()
	if err != nil {
		return [16]byte{}, err
	}
	return random, nil
}
