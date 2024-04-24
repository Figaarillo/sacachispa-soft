package util

import "github.com/google/uuid"

func AssignIfNotEmpty(field *string, newValue string) {
	if newValue != "" {
		*field = newValue
	}
}

func AssignIfNotZero(field *int, newValue int) {
	if newValue != 0 {
		*field = newValue
	}
}

func AssignIfNotZeroFloat(field *float64, newValue float64) {
	if newValue != 0 {
		*field = newValue
	}
}

func AssignUUIDIFNotEmpty(id *uuid.UUID, newID uuid.UUID) {
	if newID.String() != "" {
		*id = newID
	}
}
