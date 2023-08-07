package utils

import "strconv"

func Prefix(array interface{}, prefix string) []string {
	ret := make([]string, 0)
	switch arr := array.(type) {
	case []int:
		for _, e := range arr {
			ret = append(ret, prefix+strconv.Itoa(e))
		}
	case []string:
		for _, e := range arr {
			ret = append(ret, prefix+e)
		}
	default:
		panic("No prefix applicable")
	}
	return ret
}
