package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"log"
)

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost/api", nil)
	req.Header.Set("password", "PASSWORD")
	req.Header.Set("username", "TEST")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("[\033[91m!\033[0m] Connection error.")
		return
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
