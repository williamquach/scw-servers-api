package domain

import (
	"errors"
	"strings"
)

type ServerType string

const (
	Small  ServerType = "SMALL"
	Medium ServerType = "MEDIUM"
	Large  ServerType = "LARGE"
)

func (s *ServerType) Scan(value interface{}) error {
	if value == nil {
		return errors.New("failed to scan ServerType: value is nil")
	}

	// Assuming the value from the database is stored as a string
	value = strings.ToUpper(string(value.(string)))
	if stringValue, ok := value.(string); ok {
		*s = ServerType(stringValue)
		return nil
	}

	return errors.New("failed to scan ServerType: invalid value type")
}

func (s *ServerType) Value() (interface{}, error) {
	return string(*s), nil
}
