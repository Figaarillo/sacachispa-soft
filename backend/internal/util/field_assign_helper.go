package util

func AssignIfNotEmpty(field *string, newValue string) {
	if newValue != "" {
		*field = newValue
	}
}
