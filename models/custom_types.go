package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// StringArray adalah tipe khusus untuk menyimpan []string ke dalam kolom JSON di database.
type StringArray []string

// Value mengubah []string menjadi JSON []byte sebelum disimpan ke database.
func (a StringArray) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan membaca JSON []byte dari database dan mengubahnya kembali menjadi []string.
func (a *StringArray) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan StringArray: %v", value)
	}
	return json.Unmarshal(bytes, a)
}
