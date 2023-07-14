package domain

import (
	"errors"
	"strings"
)

type ServerStatus string

const (
	Starting ServerStatus = "STARTING"
	Running  ServerStatus = "RUNNING"
	Stopping ServerStatus = "STOPPING"
	Stopped  ServerStatus = "STOPPED"
)

func (s *ServerStatus) Scan(value interface{}) error {
	if value == nil {
		return errors.New("failed to scan ServerStatus: value is nil")
	}

	// Assuming the value from the database is stored as a string
	value = strings.ToUpper(string(value.(string)))
	if stringValue, ok := value.(string); ok {
		*s = ServerStatus(stringValue)
		return nil
	}

	return errors.New("failed to scan ServerStatus: invalid value type")
}

func (s ServerStatus) Value() (interface{}, error) {
	return string(s), nil
}
