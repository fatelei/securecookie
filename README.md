cookie
=========

Implement cookie for go, inspired by tornado.
It will help you generate cooke that is secured.


Usage
=========
```
package main

import (
	"fmt"
	"github.com/fatelei/securecookie"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	c := securecookie.SecureCookie{}
	c.Secret = "123"
	name := "test"
	value := "test"
	securedValue := c.CreateSecureCookie(name, value)
	cookie := http.Cookie{Name: name, Value: securedValue}
	http.SetCookie(w, &cookie)
	fmt.Fprintf(w, "%s", "hello world")
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(":8888", nil)
}
```