package common

import (
	"os"
	"strconv"
)

func GetEnv(key, def string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}

	return def
}

func GetIntEnv(key string, def int) int {
	if val, ok := os.LookupEnv(key); ok {
		if res, err := strconv.Atoi(val); err == nil {
			return res
		}
	}

	return def
}

func GetBoolEnv(key string, def bool) bool {
	if val, ok := os.LookupEnv(key); ok {
		res, _ := strconv.ParseBool(val)
		return res
	}

	return def
}
