package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"log"
	"os"
	"bufio"
	"strings"
)

func main() {
	fmt.Printf("Username: ")
		reader := bufio.NewReader(os.Stdin)
		user, _ := reader.ReadString('\n')
		fmt.Printf("Password: ")
		reader2 := bufio.NewReader(os.Stdin)
		pass, _ := reader2.ReadString('\n')
	for {
		fmt.Printf(">")
		cmdReader := bufio.NewReader(os.Stdin)
		cmd, _ := cmdReader.ReadString('\n')
		final := strings.Replace(cmd, "\n", "", -1)
		if final == "debug" {
			client := &http.Client{}
			req, err := http.NewRequest("GET", "http://localhost/api", nil)

			req.Header.Set("password", strings.Replace(pass, "\n", "", -1))
			req.Header.Set("username", strings.Replace(user, "\n", "", -1))
			req.Header.Set("user-agent", "Official Go Wantyapps Site API")
			res, err := client.Do(req)
			if err != nil {
				fmt.Println("[\033[91m!\033[0m] Connection error.")
				os.Exit(1)
			}
			defer res.Body.Close()
			body, _ := ioutil.ReadAll(res.Body)
			if err != nil {
				log.Fatal(err)
			}

			if string(body) == "{\"success\": true}" {
				fmt.Println("[\033[92m+\033[0m] " + string(body))
			} else {
				fmt.Println("[\033[91m!\033[0m] " + string(body))
			}
		}
	}
}
