package filter

import "errors"

var (
	ErrorNotAStruct            error = errors.New("Provided element or target is not a struct")
	ErrorNoMatchingStructField       = errors.New("Key does not match a premapped target struct field")
	ErrorUnexpectedValType           = errors.New("Val type doesn't match expected target struct field type")
)
