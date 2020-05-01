package sdp

import (
	"fmt"
	"strings"
)

func uint32ArrayToString(a []uint, delim string) string {
	return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), delim), "[]")
}

func intArrayToString(a []int, delim string) string {
	return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), delim), "[]")
}

func contains(arr []string, item string) bool {
	for _, str := range arr {
		if str == item {
			return true
		}
	}
	return false
}
