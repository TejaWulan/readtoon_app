package utility

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func MappingUnmarshal(data any) map[string]any {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil
	}

	var result map[string]any
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return nil
	}
	return result
}

func GenerateUUID() string {
	return uuid.New().String()
}

func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Now() time.Time {
	return time.Now()
}
