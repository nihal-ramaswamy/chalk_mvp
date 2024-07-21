package utils

import (
	"math/rand"

	"github.com/nihal-ramaswamy/chalk_mvp/internal/constants"
)

func NewCode(length int) string {
	allowedChars := constants.GetRuneUuidCharacters()
	lengthAllowedCharacters := len(allowedChars)

	uuidString := make([]byte, length)

	for i := range uuidString {
		uuidString[i] = byte(allowedChars[rand.Intn(lengthAllowedCharacters)])
	}

	return string(uuidString)
}
