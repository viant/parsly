package lex

import (
	"reflect"
)

//Option represents a matcher option
type Option interface{}

//AssignOption assign  option to supplied assignable
func AssignOption(options []Option, assignable interface{}) bool {
	if len(options) == 0 {
		return false
	}
	assigned := false
	for i := range options {
		option := options[i]
		if option == nil {
			continue
		}
		assignableType := reflect.TypeOf(assignable).Elem()
		optionValue := reflect.ValueOf(option)
		if assignableType == optionValue.Type() || optionValue.Type().AssignableTo(assignableType) {
			assigned = true
			reflect.ValueOf(assignable).Elem().Set(optionValue)
		}
	}
	return assigned
}
