package utils

import (
	"strconv"
)

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

func IntToString[T Integer](input T) string {
	return strconv.FormatInt(int64(input), 10)
}

func IsNullOrZeroInteger[T Integer](input *T) bool {
	if input == nil {
		return true
	}

	return *input == 0
}
