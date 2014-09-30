package securecookie

import (
	"net/http"
	"testing"
)

func TestGetSecureCookie(t *testing.T) {
	secureCookie := SecureCookie{Secret: "123"}

	cookie := http.Cookie{Name: "test", Value: "dGVzdDEyMw==|1412056511|7613cc57d905c0cdb78a5e4c9db8d3cab4e925ec"}
	_, err := secureCookie.GetSecureCookie(cookie, "test")

	if err != nil {
		t.Error("GetSecureCookie failed. Result %v", err)
	}
}
