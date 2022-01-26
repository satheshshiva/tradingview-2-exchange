package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRedact(t *testing.T) {
	ip := `{
		"strategy": "BB strategy",
			"passphrase": "Rn&6@TZQ1*0c0X#tA",
			"side":"sell"
	}`
	ip = RedactJsonField(ip, "passphrase")
	assert.NotContains(t, ip, "Rn&6@TZQ1*0c0X#tA", "redaction did not work")
	assert.Contains(t, ip, `"strategy": "BB strategy"`, "removed the prev field value")
	assert.Contains(t, ip, `"side":"sell"`, "removed the next field value")

	ip = `{
		"strategy": "BB strategy",
			"passphrase": "Rn&6@TZQ1*0c0X#tA"
	}`
	ip = RedactJsonField(ip, "passphrase")
	assert.NotContains(t, ip, "Rn&6@TZQ1*0c0X#tA", "redaction did not work")
	assert.Contains(t, ip, `"strategy": "BB strategy"`, "removed the prev field value")

	ip = `{
		"strategy": "BB strategy",
			"passphrase":"Rn&6@TZQ1*0c0X#tA"       ,
			"side":"sell"
	}`
	ip = RedactJsonField(ip, "passphrase")
	assert.NotContains(t, ip, "Rn&6@TZQ1*0c0X#tA", "redaction did not work")
	assert.Contains(t, ip, `"strategy": "BB strategy"`, "removed the prev field value")
	assert.Contains(t, ip, `"side":"sell"`, "removed the next field value")

	ip = `{
		"passphrase":"Rn&6@TZQ1*0c0X#tA"
	}`
	ip = RedactJsonField(ip, "passphrase")
	assert.NotContains(t, ip, "Rn&6@TZQ1*0c0X#tA", "redaction did not work")
}
