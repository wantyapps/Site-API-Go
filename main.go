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
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("password", "PASSWORD")
	req.Header.Set("username", "TEST")
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}
