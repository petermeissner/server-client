package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2/middleware/session"
)

// Auth parse type
type Auth struct {
	Username string `json:"username" xml:"username" form:"username"`
	Password string `json:"password" xml:"password" form:"password"`
}

// # login_handler
// - called for side effects on session
//
// - check if credentials are ok
//
// - set appropriate values in session variables
func login_handler(auth *Auth, sess *session.Session) {
	if auth.Username == "test" && auth.Password == "user" {
		sess.Set("is_logged_in", true)
		sess.Set("Username", auth.Username)
	} else {
		sess.Set("is_logged_in", false)
		sess.Set("Username", nil)
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

	if auth.Username == "test" && auth.Password == "user" {
		fmt.Println("# Auth match")
		sess.Set("is_logged_in", true)
		sess.Set("Username", auth.Username)
	} else {
		fmt.Println("# NO auth match")
		sess.Set("is_logged_in", false)
		sess.Set("Username", nil)
	}

	if err := sess.Save(); err != nil {
		panic(err)
	}
}
