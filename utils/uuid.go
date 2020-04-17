package utils

import (
	uuid "github.com/satori/go.uuid"
)

//IsValidUUID - to validate UUID
func IsValidUUID(input string) bool {
	return !(uuid.FromStringOrNil(input) == uuid.Nil)
}
