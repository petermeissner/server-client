package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2/middleware/session"
)

// Auth parse type
type Auth struct {
	Name     string `json:"name" xml:"name" form:"name"`
	Password string `json:"password" xml:"password" form:"password"`
}

// # login_handler
// - called for side effects on session
//
// - check if credentials are ok
//
// - set appropriate values in session variables
func login_handler(auth *Auth, sess *session.Session) {
	if auth.Name == "test" && auth.Password == "user" {
		sess.Set("is_logged_in", true)
		sess.Set("name", auth.Name)
	} else {
		sess.Set("is_logged_in", false)
		sess.Set("name", nil)
	}
	if err := sess.Save(); err != nil {
		panic(err)
	}
}

// # login_handler
// - called for side effects on session
//
// - check if credentials are ok
//
// - set appropriate values in session variables
func login_handler_verbose(auth *Auth, sess *session.Session) {

	fmt.Println(auth)

	if auth.Name == "test" && auth.Password == "user" {
		fmt.Println("# Auth match")
		sess.Set("is_logged_in", true)
		sess.Set("name", auth.Name)
	} else {
		fmt.Println("# NO auth match")
		sess.Set("is_logged_in", false)
		sess.Set("name", nil)
	}

	if err := sess.Save(); err != nil {
		panic(err)
	}
}
