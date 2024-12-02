package goenv

import (
	"errors"
	"fmt"
	"math/bits"
	"os"
	"reflect"
	"strconv"
)

var UnsupportedTypeError = errors.New("unsupported type")

func errorOrDefault[T comparable, TrueT any](generator func() (TrueT, error), defaultValue T) (T, error) {
	v, err := generator()
	if err != nil {
		return defaultValue, err
	}

	valueType := reflect.TypeOf(defaultValue)
	reflectedValue := reflect.ValueOf(v)

	if !reflectedValue.CanConvert(valueType) {
		return defaultValue, fmt.Errorf("cannot convert %s to %s", reflectedValue.Type().String(), valueType.String())
	}

	convertedValue := reflectedValue.Convert(valueType)
	return convertedValue.Interface().(T), nil
}

func Getenv[T comparable](key string, defaultValue T) T {
	v, _ := MustGetenv(key, defaultValue)
	return v
}

func MustGetenv[T comparable](key string, defaultValue T) (T, error) {
	envValue := os.Getenv(key)

	if envValue == "" {
		return defaultValue, nil
	}

	switch reflect.TypeOf(defaultValue).Kind() {
	case reflect.String:
		return errorOrDefault[T](func() (string, error) {
			return envValue, nil
		}, defaultValue)

	case reflect.Bool:
		return errorOrDefault[T](func() (bool, error) {
			return strconv.ParseBool(envValue)
		}, defaultValue)

	case reflect.Int:
		return errorOrDefault(func() (int, error) {
			v, err := strconv.ParseInt(envValue, 10, bits.UintSize)
			return int(v), err
		}, defaultValue)
	case reflect.Int8:
		return errorOrDefault(func() (int8, error) {
			v, err := strconv.ParseInt(envValue, 10, 8)
			return int8(v), err
		}, defaultValue)
	case reflect.Int16:
		return errorOrDefault(func() (int16, error) {
			v, err := strconv.ParseInt(envValue, 10, 16)
			return int16(v), err
		}, defaultValue)
	case reflect.Int32:
		return errorOrDefault(func() (int32, error) {
			v, err := strconv.ParseInt(envValue, 10, 32)
			return int32(v), err
		}, defaultValue)
	case reflect.Int64:
		return errorOrDefault(func() (int64, error) {
			return strconv.ParseInt(envValue, 10, 64)
		}, defaultValue)

	case reflect.Uint:
		return errorOrDefault(func() (uint, error) {
			v, err := strconv.ParseUint(envValue, 10, bits.UintSize)
			return uint(v), err
		}, defaultValue)
	case reflect.Uint8:
		return errorOrDefault(func() (uint8, error) {
			v, err := strconv.ParseUint(envValue, 10, 8)
			return uint8(v), err
		}, defaultValue)
	case reflect.Uint16:
		return errorOrDefault(func() (uint16, error) {
			v, err := strconv.ParseUint(envValue, 10, 16)
			return uint16(v), err
		}, defaultValue)
	case reflect.Uint32:
		return errorOrDefault(func() (uint32, error) {
			v, err := strconv.ParseUint(envValue, 10, 32)
			return uint32(v), err
		}, defaultValue)
	case reflect.Uint64:
		return errorOrDefault(func() (uint64, error) {
			return strconv.ParseUint(envValue, 10, 64)
		}, defaultValue)

	case reflect.Float32:
		return errorOrDefault(func() (float32, error) {
			v, err := strconv.ParseFloat(envValue, 32)
			return float32(v), err
		}, defaultValue)
	case reflect.Float64:
		return errorOrDefault(func() (float64, error) {
			return strconv.ParseFloat(envValue, 64)
		}, defaultValue)

	case reflect.Complex64:
		return errorOrDefault(func() (complex64, error) {
			v, err := strconv.ParseComplex(envValue, 64)
			return complex64(v), err
		}, defaultValue)
	case reflect.Complex128:
		return errorOrDefault(func() (complex128, error) {
			return strconv.ParseComplex(envValue, 128)
		}, defaultValue)

	case reflect.Invalid:
		fallthrough
	case reflect.Uintptr:
		fallthrough
	case reflect.Array:
		fallthrough
	case reflect.Chan:
		fallthrough
	case reflect.Func:
		fallthrough
	case reflect.Interface:
		fallthrough
	case reflect.Map:
		fallthrough
	case reflect.Pointer:
		fallthrough
	case reflect.Slice:
		fallthrough
	case reflect.Struct:
		fallthrough
	case reflect.UnsafePointer:
		fallthrough
	default:
		return defaultValue, UnsupportedTypeError
	}
}
