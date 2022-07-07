package util

import "os"

func Getenv(key, fallback string) string {
	var (
		val     string
		isExist bool
	)
	val, isExist = os.LookupEnv(key)
	if !isExist {
		val = fallback
	}
	return val
}
