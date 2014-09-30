securecookie
==============

Implement cookie for go, inspired by tornado.
It will help you generate cooke that is secured.

Install
==============
+ Intstall through go

```
go get github.com/fatelei/securecookie
```
+ Install by manual

```
> git clone git@github.com:fatelei/securecookie.git
> go install securecookie 
```


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