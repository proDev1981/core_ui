package core

import (
	"fmt"
	"strconv"
)

/*CONVERSIONS*/

// conversion to string
func String(str any) string {
	return fmt.Sprint(str)
}

// conversion to int
func Int(number any) int {
	switch number.(type) {
	case string:
		if val, err := strconv.Atoi(String(number)); err == nil {
			return val
		}
		return -1
	case int:
		return number.(int)
	default:
		return -1
	}
}

// conversion to float64
func Float(data any) float64 {
	return data.(float64)
}

// conversion to bool
func Bool(data any) bool {
	if val, ok := data.(bool); ok {
		return val
	}
	if val, ok := data.(string); ok {
		if val == "true" {
			return true
		}
		return false
	}
	return false
}

/* METHODS */

// assert
func Assert[T any](condition bool, a T, b T) T {
	if condition {
		return a
	}
	return b
}

// Default
func Default[T any](a *T, b T) {
	if LikeNil(*a) {
		*a = b
	}
}

// return if value like a nil
func LikeNil(data any) bool {
	if value := String(data); value == "" || value == "0" || value == "<nil>" || value == "{}" {
		return true
	}
	return false
}

// return new slice with items filtration
func Filter[D any](data []D, fn func(index int, item D) bool) (res []D) {

	for index, item := range data {
		if fn(index, item) {
			res = append(res, item)
		}
	}
	return
}

/* POSITIONS */

// last item
func Last[D any](data []D) D {
	return data[len(data)-1]
}

// firts item
func Fisrt[D any](data []D) D {
	return data[0]
}

// append
func Append[T any](base []T, value T) (res []T) {
	// no funciona
	res = append(base, value)
	return
}
