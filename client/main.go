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

const server_url = "http://127.0.0.1:3000"

// func get_input(r *bufio.Reader) string {
// 	fmt.Print("-> ")
// 	text, _ := r.ReadString('\n')

// 	// remove line endings from input string
// 	text = strings.Replace(text, "\n", "", -1)
// 	text = strings.Replace(text, "\r", "", -1)

// 	return text
// }

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
	command_string := "Commands: \nhelp\necho on\necho off\nping\nstart"
	fmt.Println(command_string)

	echo := false
	command_known := true

	for {

		if command_known == false {
			fmt.Println("# Command unknown.")
		}
		command_known = false

		text := gbc.Get_input(">: ")
		// text := get_input(reader)

		// help command
		if strings.Compare("help", text) == 0 {
			command_known = true

			fmt.Println(command_string)
		}

		// echo command
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
		if strings.Compare("ping", text) == 0 {
			command_known = true

			resp, err := http_client.Get(server_url)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("# Pinging server, status code: ", resp.StatusCode)
		}

		// start server connection and session
		if strings.Compare("start", text) == 0 {
			command_known = true

			resp, err := http_client.Get(server_url)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("# Pinging server, status code: ", resp.StatusCode)

			// auth := gbc.Get_auth_from_term()

			// execute
			//data := url.Values{"name": []string{auth.Username}, "password": []string{auth.Password}}
			data := url.Values{"name": []string{"test"}, "password": []string{"user"}}
			post_res, err := http_client.PostForm(server_url+"/login", data)

			server_parsed_url, err := url.Parse(server_url)
			if err != nil {
				fmt.Println(err)
			}
			cookies := http_client.Jar.Cookies(server_parsed_url)

			fmt.Println("\n# Done POST to server, status code:", post_res.StatusCode)
			post_body, err := io.ReadAll(resp.Body)
			for name, headers := range post_res.Header {
				for _, hdr := range headers {
					println("# " + name + ": " + hdr)
				}
			}
			for _, cookie := range cookies {
				fmt.Println("cookies?", cookie.Name, cookie.Value)
			}
			fmt.Println("\n# Body content\n" + string(post_body))

			get_res, err := http_client.Get(server_url)
			get_body, err := io.ReadAll(get_res.Body)

			fmt.Println(string(get_body))
		}
	}
}
