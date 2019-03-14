package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func Quicktest() {
	client := http.Client{}
	req, err := http.NewRequest("PUT", "http://localhost:8080", strings.NewReader(`{"key": "foo","value": "bar"}`))
	if err != nil {
		log.Fatalln(err)
	}
	
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	
	req, err = http.NewRequest("GET", "http://localhost:8080", strings.NewReader(`{"key": "foo"}`))
	if err != nil {
		log.Fatalln(err)
	}
	
	resp, err = client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	resbytes, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(resbytes))
	
}
