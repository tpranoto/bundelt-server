package common

import "strconv"

func ConvToInt64(input string) (int64, error) {
	return strconv.ParseInt(input, 10, 64)
}
