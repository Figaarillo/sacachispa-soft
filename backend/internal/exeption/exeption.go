package exeption

import "errors"

var ErrorMissingParam = errors.New("error: url param is empty or not found")
