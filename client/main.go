package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	fmt.Println("Startup")
	resp, err := http.Get("http://127.0.0.1:3000")
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

	echo := false
	start := true

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')

		// remove line endings from input string
		text = strings.Replace(text, "\n", "", -1)
		text = strings.Replace(text, "\r", "", -1)

		// startup and help
		if start == true || strings.Compare("help", text) == 0 {
			fmt.Println("Commands: \nhelp\necho on\necho off\nping")
			start = false
		}

		// echo command
		if echo {
			fmt.Println(text)
			fmt.Println(hex.EncodeToString([]byte(text)))
		}
		if strings.Compare("echo on", text) == 0 {
			echo = true
		}
		if strings.Compare("echo off", text) == 0 {
			echo = false
		}

		// ping server
		if strings.Compare("ping", text) == 0 {
			resp, err := http.Get("http://127.0.0.1:3000")
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("# Pinging server, status code: ", resp.StatusCode)
		}

	}
}
