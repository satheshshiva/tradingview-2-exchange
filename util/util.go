package util

import "os"

var isProd = false

func init() {
	if _, ok := os.LookupEnv("PRODUCTION"); ok {
		isProd = true
	}
}
func IsProd() bool {
	return isProd
}
