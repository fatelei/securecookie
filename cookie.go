//Package cookie implement a simple library
//for cookie
package cookie

import (
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type list []interface{}

type SignedCookie struct {
	secret string
}

// CreateSecureCookie creates a cookie that is secured
func (p *SignedCookie) CreateSecureCookie(name string, value string) (cookie string) {
	cookie = p.createSignedValue(name, value)
	return
}

// createSignedValue creates a signed value
func (p *SignedCookie) createSignedValue(name string, value string) (signedValue string) {
	var timestamp string
	var signed_value string

	timestamp = strconv.FormatInt(time.Now().Unix(), 10)

	signed_value = p.createSignedValue(name, value, timestamp)
	ary := make(list, 3)
	ary[0] = value
	ary[1] = timestamp
	ary[2] = signed_value

	signedValue = strings.Join(ary, "|")
	return
}

// createSignature creates a signture
func (p *SignedCookie) createSignature(name string, value string, timestamp string) (signature string) {
	secret := []byte(p.secret)
	hmac := hmac.New(sha1.New, secret)
	hmac.Write([]byte(name))
	hmac.Write([]byte(value))
	hmac.Write([]byte(timestamp))
	signature = fmt.Printf("%x", hmac.Sum(nil))
	return
}

// GetSecureCookie gets the origin value
func (p *SignedCookie) GetSecureCookie(c http.Cookie, name string) (value string) {
	var signed_value string
	var signed_ary []string
	var tmp_value string
	var timestamp string
	var signature string

	signed_value = c[name]
	signed_ary = strings.Split(signed_value, "|")

	tmp_value = signed_ary[0]
	timestamp = signed_ary[1]
	signature = signed_ary[2]

}
