package util

import (
	"os"
	"regexp"
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

func RedactJsonField(input string, fieldName string) string {
	m1 := regexp.MustCompile(`"` + fieldName + `":\s*".*"`)
	return m1.ReplaceAllString(input, `"`+fieldName+`": `+`"<redacted>"`)
}
