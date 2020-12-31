package main

import (
	"strconv"
	"strings"
)

type ints []int64

// String converts the int64 array into a comma separated string.
func (ints ints) String() string {
	intsStr := make([]string, len(ints))
	for i, v := range ints {
		intsStr[i] = strconv.FormatInt(v, 10)
	}
	return strings.Join(intsStr, ",")
}
