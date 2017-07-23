// Package filter provides abstractions to filter an arbitrary slice of structs
// based on certain criteria
package filter

import (
	"errors"
	"reflect"
)

type Constraints interface {
	// Add a new constraint on a struct field (equality)
	Add(key string, val interface{}) error
	// Validate ensures that elem meets the specified constraints
	Validate(elem interface{}) (bool, error)
}

type constraints struct {
	constraintMap map[string]interface{}
	// targetStructType represent the type of the target struct elements
	// to be filtered
	targetStructType reflect.Type
	// Maps target struct fields to their type
	targetTypeMap map[string]reflect.Type
}

func (c *constraints) Add(key string, val interface{}) error {
	// Ensure key actually maps to a target field struct name
	expType, ok := c.targetTypeMap[key]
	if !ok {
		return ErrorNoMatchingStructField
	}

	// Validate key's val type
	if reflect.TypeOf(val) != expType {
		return ErrorUnexpectedValType
	}
	// Add to constraint map
	c.constraintMap[key] = val
	return nil
}

func (c *constraints) Validate(elem interface{}) (bool, error) {
	// Ensure that the elem is a struct TODO: Change to ptr to struct later
	if !isStruct(elem) {
		return false, ErrorNotAStruct
	}

	for key, val := range c.constraintMap {
		// Fetch the struct field from elem
		eq := reflect.DeepEqual(reflect.ValueOf(elem).FieldByName(key).Interface(),
			reflect.ValueOf(val).Interface())
		if !eq {
			return false, nil
		}
	}
	return true, nil
}

func isStruct(t interface{}) bool {
	return reflect.ValueOf(t).Kind() == reflect.Struct
}

func isSlice(t interface{}) bool {
	return reflect.ValueOf(t).Kind() == reflect.Slice
}

// NewConstraints generates a new Constraints representation based on the
// targetStruct that the constraints will be applied to during filtering.
func NewConstraints(targetStruct interface{}) (Constraints, error) {
	// Ensure targetStruct is a struct
	if !isStruct(targetStruct) {
		return nil, ErrorNotAStruct
	}

	// Construct mapping of target struct field names to their types
	tTypeMap := make(map[string]reflect.Type)
	tType := reflect.TypeOf(targetStruct)
	for i := 0; i < tType.NumField(); i++ {
		tTypeMap[tType.Field(i).Name] = tType.Field(i).Type // Map target struct field to type
	}

	cMap := make(map[string]interface{})
	c := constraints{targetStructType: tType, targetTypeMap: tTypeMap, constraintMap: cMap}
	return &c, nil
}

func Filter(in interface{}, cons Constraints) (out []interface{}, err error) {
	if !isSlice(in) {
		return []interface{}{nil}, errors.New("")
	}
	// Map to []interface{}
	//TODO(suyashkumar): Can we be more efficient?
	inVal := reflect.ValueOf(in)
	inLen := inVal.Len()
	inSlc := make([]interface{}, inLen)

	for i := 0; i < inLen; i++ {
		inSlc[i] = inVal.Index(i).Interface()
	}

	out = make([]interface{}, len(inSlc))
	outIndex := 0
	for _, e := range inSlc {
		ok, err := cons.Validate(e)

		if err != nil {
			// Error in validation of this element
			return []interface{}{nil}, err
		}

		if ok {
			// Add to filtered array
			out[outIndex] = e
			outIndex++
		}
	}

	return out[:outIndex], nil
}
