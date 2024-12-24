// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Internal use only. Unauthorized use is prohibited.
// Contact: legal@focela.com

// Package reflection provides utility functions for working with Go's reflect package.
package reflection

import (
	"reflect"
)

// OriginValueAndKindOutput holds the input and original value and kind information.
type OriginValueAndKindOutput struct {
	InputValue  reflect.Value
	InputKind   reflect.Kind
	OriginValue reflect.Value
	OriginKind  reflect.Kind
}

// OriginValueAndKind retrieves and returns the original reflect value and kind.
// It resolves pointer indirections to determine the final underlying value and kind.
func OriginValueAndKind(value interface{}) (out OriginValueAndKindOutput) {
	if v, ok := value.(reflect.Value); ok {
		out.InputValue = v
	} else {
		out.InputValue = reflect.ValueOf(value)
	}
	out.InputKind = out.InputValue.Kind()
	out.OriginValue = out.InputValue
	out.OriginKind = out.InputKind

	for out.OriginKind == reflect.Ptr {
		out.OriginValue = out.OriginValue.Elem()
		out.OriginKind = out.OriginValue.Kind()
	}
	return
}

// OriginTypeAndKindOutput holds the input and original type and kind information.
type OriginTypeAndKindOutput struct {
	InputType  reflect.Type
	InputKind  reflect.Kind
	OriginType reflect.Type
	OriginKind reflect.Kind
}

// OriginTypeAndKind retrieves and returns the original reflect type and kind.
// It resolves pointer indirections to determine the final underlying type and kind.
func OriginTypeAndKind(value interface{}) (out OriginTypeAndKindOutput) {
	if value == nil {
		return
	}
	if reflectType, ok := value.(reflect.Type); ok {
		out.InputType = reflectType
	} else {
		if reflectValue, ok := value.(reflect.Value); ok {
			out.InputType = reflectValue.Type()
		} else {
			out.InputType = reflect.TypeOf(value)
		}
	}
	out.InputKind = out.InputType.Kind()
	out.OriginType = out.InputType
	out.OriginKind = out.InputKind

	for out.OriginKind == reflect.Ptr {
		out.OriginType = out.OriginType.Elem()
		out.OriginKind = out.OriginType.Kind()
	}
	return
}

// ValueToInterface converts a reflect.Value to its interface representation.
// It handles basic types, pointers, and interface kinds.
func ValueToInterface(v reflect.Value) (value interface{}, ok bool) {
	if v.IsValid() && v.CanInterface() {
		return v.Interface(), true
	}

	switch v.Kind() {
	case reflect.Bool:
		return v.Bool(), true
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int(), true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint(), true
	case reflect.Float32, reflect.Float64:
		return v.Float(), true
	case reflect.Complex64, reflect.Complex128:
		return v.Complex(), true
	case reflect.String:
		return v.String(), true
	case reflect.Ptr:
		return ValueToInterface(v.Elem())
	case reflect.Interface:
		return ValueToInterface(v.Elem())
	default:
		return nil, false
	}
}
