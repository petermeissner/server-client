package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const server_url = "http://127.0.0.1:3000"

func get_input(r *bufio.Reader) string {
	fmt.Print("-> ")
	text, _ := r.ReadString('\n')

	// remove line endings from input string
	text = strings.Replace(text, "\n", "", -1)
	text = strings.Replace(text, "\r", "", -1)

	return text
}

func main() {
	fmt.Println("Startup")
	resp, err := http.Get(server_url)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Status)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))

	reader := bufio.NewReader(os.Stdin)
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

		text := get_input(reader)

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

			resp, err := http.Get(server_url)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("# Pinging server, status code: ", resp.StatusCode)
		}

		// start server connection and session
		if strings.Compare("start", text) == 0 {
			command_known = true

			resp, err := http.Get(server_url)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("# Pinging server, status code: ", resp.StatusCode)

			// fmt.Print("\nName?")
			// user := get_input(reader)

			// fmt.Print("Password?")
			// pw := get_input(reader)

			// execute
			// data := url.Values{"name": []string{user}, "password": []string{pw}}
			data := url.Values{"name": []string{"test"}, "password": []string{"user"}}
			post_res, err := http.PostForm(server_url+"/login", data)
			fmt.Println("\n# Done POST to server, status code:", post_res.StatusCode)
			post_body, err := io.ReadAll(resp.Body)
			for name, headers := range post_res.Header {
				for _, hdr := range headers {
					println("# " + name + ": " + hdr)
				}
			}
			for _, cookie := range post_res.Cookies() {
				fmt.Println(cookie.Name, cookie.Value)
			}
			fmt.Println("\n# Body content\n" + string(post_body))

		}
	}
}
