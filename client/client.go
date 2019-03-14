package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	
	"CacheCLI/server"
)

const URI = "http://localhost:8080"

type CacheClient struct{
	URI string
	Client http.Client
}

func NewCacheClient()*CacheClient{
	return &CacheClient{URI:URI,Client:http.Client{}}
}
func (c *CacheClient) Create(key,value string) error{
	//create JSON with key and value by first putting into data object then marshal to JSON
	message := server.Data{Key: key, Value: value}
	byteData, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}
	
	//create HTTP request (no built in http.Put" in Go http package so using .NewRequest)
	request, err := http.NewRequest("PUT", URI, bytes.NewBuffer(byteData))
	if err != nil {
		log.Fatalln(err)
	}
	//Send request to server (make request)
	client := &http.Client{}
	resp, err := client.Do(request)
	
	//check status to see if error exists
	if err != nil {
		log.Fatalln(err)
	}
	
	//if no error, take response stream, set response interface and hydrate it with response body
	var result string
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Fatalln(err)
	}
	
	//Create command returns string to terminal(log on server side for verification)
	log.Println(result)
	fmt.Println(result)
	
	//close body (because this returns an error couldn't find simple way to defer and handle error so just making it
	// last call in method)
	err = resp.Body.Close()
	if err != nil {
		log.Fatalln(err)
	}
	return nil
}

func (c *CacheClient) Read(key string) (string, error){
	//create JSON with key
	message := server.Data{Key: key}
	byteData, err := json.Marshal(message)
	if err !=nil {
		log.Fatalln(err)
	}

	//create http request - not using http.Get because it doesn't take in body message
	request, err := http.NewRequest("GET", URI, bytes.NewBuffer(byteData))
	if err != nil {
		log.Fatalln(err)
	}
	
	//send request to server
	client := &http.Client{}
	resp, err := client.Do(request)
	
	//check status - report if errors exist
	if err != nil {
		log.Fatalln(err)
	}
	
	//turn  response into byte slice and then into string with decoder/decode
	var result string
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Fatalln(err)
	}
	
	//log result on server side for verification
	log.Println(result)
	
	//close body response
	err = resp.Body.Close()
	if err != nil {
		log.Fatalln(err)
	}
	return result, nil
}

func (c *CacheClient) Update(key, value string) error{
	//creates JSON with key and value
	message := server.Data{Key: key, Value: value}
	byteData, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}
	
	//creates http request
	request, err := http.NewRequest("POST", URI, bytes.NewBuffer(byteData))
	if err != nil {
		log.Fatalln(err)
	}
	
	//send request to server
	client := &http.Client{}
	resp, err := client.Do(request)
	
	//check status - report if errors exist
	if err != nil {
		log.Fatalln(err)
	}
	
	//if no error exists, decode json and return result
	var result string
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Fatalln(err)
	}
	
	//Update command returns string to terminal(log on server side for verification)
	log.Println(result)
	fmt.Println(result)
	
	//close body
	err = resp.Body.Close()
	if err != nil {
		log.Fatalln(err)
	}
	return nil
}

func (c *CacheClient) Delete(key string) error{
	//create JSON from key
	message := server.Data{Key: key}
	byteData, err := json.Marshal(message)
	if err !=nil {
		log.Fatalln(err)
	}
	
	//create http request - not using http.Get because it doesn't take in body message
	request, err := http.NewRequest("DELETE", URI, bytes.NewBuffer(byteData))
	if err != nil {
		log.Fatalln(err)
	}
	
	//send request to server
	client := &http.Client{}
	resp, err := client.Do(request)
	
	//check status - report if errors exist
	if err != nil {
		log.Fatalln(err)
	}
	
	//turn  response into byte slice and then into string with decoder/decode
	var result string
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Fatalln(err)
	}
	
	//Delete command returns string to terminal; also log result on server side for verification
	log.Println(result)
	fmt.Println(result)
	
	//close body response
	err = resp.Body.Close()
	if err != nil {
		log.Fatalln(err)
	}
	return  nil
}