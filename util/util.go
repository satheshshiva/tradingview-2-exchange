package util

import (
	"os"
	"strings"
)

var isProd = false

func init() {
	if isProdEnv, ok := os.LookupEnv("PRODUCTION"); ok {
		if strings.ToLower(isProdEnv) == "true" {
			isProd = true
		}
	}
}
func IsProd() bool {
	return isProd
}
