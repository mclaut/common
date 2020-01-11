package common

import (
	"strconv"
	"strings"
)

func ParseFloat(s string) (float64, error) {
	s = trim(s)
	if len(s) == 0 {
		return 0, nil
	}
	return strconv.ParseFloat(s, 64)
}

func ParseInt(s string) (int64, error) {
	s = trim(s)
	if len(s) == 0 {
		return 0, nil
	}
	return strconv.ParseInt(s, 64, 8)
}

func CheckTime(s string) bool {
	s = trim(s)
	if len(s) == 0 {
		return false
	}
	return true
}

func trim(s string) string {
	return strings.Trim(s, " \r\n")
}
