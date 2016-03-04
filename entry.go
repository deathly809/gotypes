package gotypes

import (
	"fmt"
	"reflect"
)

// Type is used to determine the underlying
// type at runtime
type Type int

// Data driving around the block
var (
	conversionFunctions = map[Type]func(interface{}) (Value, error){
		Integer: func(val interface{}) (Value, error) {
			if v, ok := val.(int64); ok {
				return WrapInteger(v), nil
			}
			return nil, fmt.Errorf("Could not convert to int64")
		},
		Real: func(val interface{}) (Value, error) {
			if v, ok := val.(float64); ok {
				return WrapReal(v), nil
			}
			return nil, fmt.Errorf("Could not convert to float64")
		},
		Text: func(val interface{}) (Value, error) {
			if v, ok := val.(string); ok {
				return WrapText(v), nil
			}
			return nil, fmt.Errorf("Could not convert to string")
		},
		Boolean: func(val interface{}) (Value, error) {
			if v, ok := val.(bool); ok {
				return WrapBoolean(v), nil
			}
			return nil, fmt.Errorf("Could not convert to bool")
		},
		Unknown: func(val interface{}) (Value, error) {
			return WrapUnknown(val), nil
		},
		Array: func(val interface{}) (Value, error) {
			if v, ok := val.([]Value); ok {
				return (*array)(&v), nil
			}
			return nil, fmt.Errorf("Could not convert to array of Value")
		},
	}
)

// Types supported
const (
	Unknown   = Type(iota)
	Array     = Type(iota)
	Real      = Type(iota)
	Integer   = Type(iota)
	Boolean   = Type(iota)
	Text      = Type(iota)
	ValueType = Type(iota)
)

// Value represents any type of value, this is used
// to pass around variadic sized arrays or values not
// known until runtime
type Value interface {
	Type() Type
	Array() []Value
	Real() float64
	Integer() int64
	Bool() bool
	Text() string
	Unknown() interface{}
}

// WrapInteger takes an integer and returns a value
func WrapInteger(value int64) Value {
	return (*integer)(&value)
}

// WrapReal takes a real and returns a value
func WrapReal(value float64) Value {
	return (*real)(&value)
}

// WrapBoolean takes a boolean and returns a value
func WrapBoolean(value bool) Value {
	return (*boolean)(&value)
}

// WrapText takes a string and returns a value
func WrapText(value string) Value {
	return (*text)(&value)
}

// WrapUnknown takes an interface and returns a value
func WrapUnknown(value interface{}) Value {
	return &unknown{value: value}
}

// WrapArray takes an array and returns an array of values
func WrapArray(value interface{}, theType Type) (result Value, err error) {
	if value != nil {
		if t := reflect.TypeOf(value).Kind(); t == reflect.Slice { // is an array
			if theType == ValueType {
				tmp := value.([]Value)
				result = (*array)(&tmp)
			} else {
				s := reflect.ValueOf(value)
				tmp := make([]Value, s.Len())
				r := (*array)(&tmp)
				for i := 0; i < s.Len(); i++ {
					if v, err := conversionFunctions[theType](s.Index(i)); err != nil {
						result = nil
						break
					} else {
						(*r)[i] = v
					}
				}
				result = r
			}

		} else {
			err = fmt.Errorf("Expected array, found %s", t.String())
		}
	} else {
		err = fmt.Errorf("nil pointer exception")
	}
	return
}
