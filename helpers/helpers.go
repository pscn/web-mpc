package helpers

import (
	"strconv"
)

// ToFloat the string
func ToFloat(str string) float32 {
	val, err := strconv.ParseFloat(str, 32)
	if err != nil { // FIXME: should we handle this?
		return 0.0
	}
	return float32(val)
}

// ToInt the string
func ToInt(str string) int {
	val, err := strconv.ParseInt(str, 10, 64)
	if err != nil { // FIXME: should we handle this?
		return 0
	}
	return int(val)
}

// ToBool the string
func ToBool(str string) bool {
	switch str {
	case "1":
		return true
	}
	return false
}

// eof
