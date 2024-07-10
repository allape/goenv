package goenv

import (
	"math/bits"
	"os"
	"strconv"
)

func errorOrDefault[T comparable, TrueT any](generator func() (TrueT, error), defaultValue T) (T, error) {
	v, err := generator()
	if err != nil {
		return defaultValue, err
	}
	return any(v).(T), nil
}

func Getenv[T comparable](key string, defaultValue T) T {
	v, _ := GetSafeEnv(key, defaultValue)
	return v
}

func GetSafeEnv[T comparable](key string, defaultValue T) (T, error) {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue, nil
	}

	switch any(defaultValue).(type) {
	case string:
		return any(value).(T), nil

	case bool:
		return errorOrDefault[T](func() (bool, error) {
			return strconv.ParseBool(value)
		}, defaultValue)

	case int:
		return errorOrDefault(func() (int, error) {
			v, err := strconv.ParseInt(value, 10, bits.UintSize)
			return int(v), err
		}, defaultValue)
	case int8:
		return errorOrDefault(func() (int8, error) {
			v, err := strconv.ParseInt(value, 10, 8)
			return int8(v), err
		}, defaultValue)
	case int16:
		return errorOrDefault(func() (int16, error) {
			v, err := strconv.ParseInt(value, 10, 16)
			return int16(v), err
		}, defaultValue)
	case int32:
		return errorOrDefault(func() (int32, error) {
			v, err := strconv.ParseInt(value, 10, 32)
			return int32(v), err
		}, defaultValue)
	case int64:
		return errorOrDefault(func() (int64, error) {
			return strconv.ParseInt(value, 10, 64)
		}, defaultValue)

	case uint:
		return errorOrDefault(func() (uint, error) {
			v, err := strconv.ParseUint(value, 10, bits.UintSize)
			return uint(v), err
		}, defaultValue)
	case uint8:
		return errorOrDefault(func() (uint8, error) {
			v, err := strconv.ParseUint(value, 10, 8)
			return uint8(v), err
		}, defaultValue)
	case uint16:
		return errorOrDefault(func() (uint16, error) {
			v, err := strconv.ParseUint(value, 10, 16)
			return uint16(v), err
		}, defaultValue)
	case uint32:
		return errorOrDefault(func() (uint32, error) {
			v, err := strconv.ParseUint(value, 10, 32)
			return uint32(v), err
		}, defaultValue)
	case uint64:
		return errorOrDefault(func() (uint64, error) {
			return strconv.ParseUint(value, 10, 64)
		}, defaultValue)

	case float32:
		return errorOrDefault(func() (float32, error) {
			v, err := strconv.ParseFloat(value, 32)
			return float32(v), err
		}, defaultValue)
	case float64:
		return errorOrDefault(func() (float64, error) {
			return strconv.ParseFloat(value, 64)
		}, defaultValue)

	case complex64:
		return errorOrDefault(func() (complex64, error) {
			v, err := strconv.ParseComplex(value, 64)
			return complex64(v), err
		}, defaultValue)
	case complex128:
		return errorOrDefault(func() (complex128, error) {
			return strconv.ParseComplex(value, 128)
		}, defaultValue)
	}

	return defaultValue, nil
}
