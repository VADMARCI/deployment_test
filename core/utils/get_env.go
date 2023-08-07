package utils

import (
	"os"
	"strconv"
)

func GetEnvString(name string) string {

	return os.Getenv(name)
}
func GetEnvBool(name string) bool {
	s := GetEnvString(name)
	i, err := strconv.ParseBool(s)
	if nil != err {
		return false
	}
	return i
}
func GetEnvInt(name string) int64 {
	s := GetEnvString(name)
	i, err := strconv.ParseInt(s, 10, 0)
	if nil != err {
		return 0
	}
	return i
}
func GetEnvFloat(name string) float64 {
	s := GetEnvString(name)
	i, err := strconv.ParseFloat(s, 64)
	if nil != err {
		return 0
	}
	return i
}
