package utils

import (
	"github.com/go-basic/uuid"
	"strings"
)

//Generate 32-bit uuid string
func CreateUUID() string {
	id := uuid.New()
	split := strings.Split(id, "-")
	uuidStr := strings.Join(split, "")
	return uuidStr
}
