package utils

import (
	"github.com/google/uuid"
)

func IsNullOrEmptyUUID(input *uuid.UUID) bool {
	if input == nil || *input == uuid.Nil {
		return true
	}

	return false
}
