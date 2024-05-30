package exeption

import "errors"

var (
	ErrorMissingParam    = errors.New("error: url param is empty or not found")
	ErrorInvalidBody     = errors.New("error: invalid body provided")
	ErrInvalidPagination = errors.New("error: invalid pagination provided")
)
