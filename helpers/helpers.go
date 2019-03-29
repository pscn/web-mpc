package helpers

import (
	"strconv"
)

func ToFloat64(str string) float64 {
	val, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0.0
	}
	return val
}

func ToInt64(str string) int64 {
	val, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return val
}

func ToBool(str string) bool {
	switch str {
	case "1":
		return true
	}
	return false
}
