//Package cookie implement a simple library
//for cookie
package securecookie

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type SecureCookie struct {
	Secret string
}

// CreateSecureCookie creates a cookie that is secured
func (p *SecureCookie) CreateSecureCookie(name string, value string) (cookie string) {
	cookie = p.createSignedValue(name, value)
	return
}

// createSignedValue creates a signed value
func (p *SecureCookie) createSignedValue(name string, value string) (signedValue string) {
	var timestamp string
	var signed_value string

	timestamp = strconv.FormatInt(time.Now().Unix(), 10)

	signed_value = p.createSignature(name, value, timestamp)
	encodedValue := base64.StdEncoding.EncodeToString([]byte(value))
	ary := make([]string, 3)
	ary[0] = encodedValue
	ary[1] = timestamp
	ary[2] = signed_value

	signedValue = strings.Join(ary, "|")
	return
}

// createSignature creates a signture
func (p *SecureCookie) createSignature(name string, value string, timestamp string) (signature string) {
	secret := []byte(p.Secret)
	mac := hmac.New(sha1.New, secret)
	mac.Write([]byte(name))
	mac.Write([]byte(value))
	mac.Write([]byte(timestamp))
	signature = hex.EncodeToString(mac.Sum(nil))
	return
}

// GetSecureCookie gets the origin value
func (p *SecureCookie) GetSecureCookie(cookie http.Cookie, name string) (value string, err error) {
	var signedValue string
	var signedAry []string

	signedValue = cookie.Value
	signedAry = strings.Split(signedValue, "|")

	signature := p.createSignature(name, signedAry[0], signedAry[1])

	if isNotEqual := hmac.Equal([]byte(signedAry[2]), []byte(signature)); isNotEqual {
		err = errors.New("Invalid cookie signature")
		return
	}

	timestamp, err := strconv.Atoi(signedAry[1])
	if err != nil {
		return
	}

	now := time.Now().Unix()

	if int64(timestamp) < now-31*86400 {
		err = errors.New("Expired cookie")
		return
	}

	if int64(timestamp) > now+31*86400 {
		err = errors.New("Cookie timestamp in future; possible tampering.")
		return
	}

	byteValue, err := base64.StdEncoding.DecodeString(signedAry[0])
	if err == nil {
		value = string(byteValue)
		return
	}
	return
}
