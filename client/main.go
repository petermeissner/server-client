package main

import (
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"

	gbc "github.com/petermeissner/golang-basic-cred/library"
)

// CONSTANTS
const server_url = "http://127.0.0.1:3000"

// MAIN
func main() {

	jar, err := cookiejar.New(nil)
	if err != nil {
		// error handling
	}
	http_client := http.Client{
		Timeout: time.Duration(3) * time.Second,
		Jar:     jar,
	}

	fmt.Println("Startup")
	resp, err := http_client.Get(server_url)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Status)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))

	fmt.Println("Simple Shell")
	fmt.Println("---------------------")

	// startup and help
	command_string := "Commands: \nhelp\necho on\necho off\nping\ntest login\nstart"
	fmt.Println(command_string)

	echo := false
	command_known := true

	for {

		// handle unknown commands
		if command_known == false {
			fmt.Println("# Command unknown.")
		}
		command_known = false

		// Get user input
		//
		text := gbc.Get_input(">: ")

		// help command
		//
		if strings.Compare("help", text) == 0 {
			command_known = true

			fmt.Println(command_string)
		}

		// echo command on/off
		//
		if echo {
			fmt.Println(text)
			fmt.Println(hex.EncodeToString([]byte(text)))
		}
		if strings.Compare("echo on", text) == 0 {
			command_known = true

			echo = true
		}
		if strings.Compare("echo off", text) == 0 {
			command_known = true

			echo = false
		}

		// ping server
		//
		if strings.Compare("ping", text) == 0 {
			command_known = true

			resp, err := http_client.Get(server_url)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("# Pinging server, status code: ", resp.StatusCode)
		}

		// test if login workd
		//
		if strings.Compare("test login", text) == 0 {
			// acknowledge that command will be processed
			command_known = true
			cli_test_login(http_client, resp)
		}

		// start server connection and session
		//
		if strings.Compare("start", text) == 0 {
			// acknowledge that command will be processed
			command_known = true
			cli_start(http_client, resp)
		}
	}
}

// the main client-server-loop
//
// ...
func cli_start(http_client http.Client, resp *http.Response) {

}

// cli_test_login
//
// ...
func cli_test_login(http_client http.Client, resp *http.Response) {

	// get authentication input and send it to server
	//
	auth := gbc.Get_auth_from_term()
	data := url.Values{"username": []string{auth.Username}, "password": []string{auth.Password}}
	post_res, err := http_client.PostForm(server_url+"/login", data)

	// report results
	//

	// headers
	fmt.Println("\n# Done POST to server, status code:", post_res.StatusCode)
	for username, headers := range post_res.Header {
		for _, hdr := range headers {
			println("# " + username + ": " + hdr)
		}
	}

	// cookies
	server_parsed_url, err := url.Parse(server_url)
	if err != nil {
		fmt.Println(err)
	}
	cookies := http_client.Jar.Cookies(server_parsed_url)
	for _, cookie := range cookies {
		fmt.Println("cookies?", cookie.Name, cookie.Value)
	}

	// body
	post_body, err := io.ReadAll(resp.Body)
	fmt.Println("\n# Body content\n" + string(post_body))

	get_res, err := http_client.Get(server_url)
	get_body, err := io.ReadAll(get_res.Body)

	fmt.Println(string(get_body))
}
