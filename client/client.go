package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
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
		return err
	}
	
	//create HTTP request (no built in http.Put" in Go http package so using .NewRequest)
	request, err := http.NewRequest("PUT", URI, bytes.NewBuffer(byteData))
	if err != nil {
		log.Fatalln(err)
		return err
	}
	
	//Send request to server (make request)
	client := &http.Client{}
	resp, err := client.Do(request)
	
	//close body
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}()
	
	if err != nil {
		log.Fatalln(err)
		return err
	}
	
	//check status to see if  sever error exists
	if resp.StatusCode != http.StatusCreated {
		err := fmt.Errorf("create failed: '%v'", resp.StatusCode)
		log.Fatalln(err)
		return err
	}
	
	//if no error, take response stream, convert reader to byte slice
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err == io.EOF{
		//"End of File" means end of response read, convert response to byte slice to string
		log.Println(string(respBytes))
		
		return err
	}
	if err != nil {
		log.Println(err)
		log.Fatalln(err)
		
		return err
	}
	
	//Create command returns string to terminal(log on server side for verification)
	//convert byte slice to string
	log.Println(string(respBytes))
	fmt.Println(string(respBytes))
	
	return nil
}

func (c *CacheClient) Read(key string) (string, error){
	//create JSON with key
	message := server.Data{Key: key}
	byteData, err := json.Marshal(message)
	if err !=nil {
		log.Fatalln(err)
		return "", err
	}

	//create http request - not using http.Get because it doesn't take in body message
	request, err := http.NewRequest("GET", URI, bytes.NewBuffer(byteData))
	if err != nil {
		log.Fatalln(err)
		return "", err
	}
	
	//send request to server
	client := &http.Client{}
	resp, err := client.Do(request)
	
	//close body
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}()

	if err != nil {
		log.Fatalln(err)
		return "", err
	}
	
	//check status - report if errors exist
	if resp.StatusCode != http.StatusOK {
		err := fmt.Errorf("read failed: '%v'", resp.StatusCode)
		log.Fatalln(err)
		return "", err
	}
	
	//if no error, take response stream, convert reader to byte slice
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err == io.EOF{
		//"End of File" means end of response read, convert response to byte slice to string
		result := string(respBytes)
		log.Println(result)
		
		return result, err
	}
	if err != nil {
		log.Println(err)
		log.Fatalln(err)
		
		return "", err
	}
	
	//log result on server side for verification
	log.Println(string(respBytes))
	
	//close body response
	err = resp.Body.Close()
	if err != nil {
		log.Fatalln(err)
	}
	
	//Read command returns a string (result of read query) and error
	return string(respBytes), nil
}

func (c *CacheClient) Update(key, value string) error{
	//creates JSON with key and value
	message := server.Data{Key: key, Value: value}
	byteData, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	
	//creates http request
	request, err := http.NewRequest("POST", URI, bytes.NewBuffer(byteData))
	if err != nil {
		log.Fatalln(err)
		return err
	}
	
	//send request to server
	client := &http.Client{}
	resp, err := client.Do(request)
	
	//close body
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}()
	
	if err != nil {
		log.Fatalln(err)
	}
	
	//check status - report if errors exist
	if resp.StatusCode != http.StatusCreated{
		err := fmt.Errorf("update failed: '%v'", resp.StatusCode)
		log.Fatalln(err)
		return err
	}
	
	//if no error exists, convert read to byte slice then string and return success string result
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err == io.EOF{
		//"End of File" means end of response read, convert response to byte slice to string
		log.Println(string(respBytes))
		
		return err
	}
	if err != nil {
		log.Println(err)
		log.Fatalln(err)
		
		return err
	}
	
	//Update command returns string to terminal(log on server side for verification)
	log.Println(string(respBytes))
	fmt.Println(string(respBytes))
	return nil
}

func (c *CacheClient) Delete(key string) error{
	//create JSON from key
	message := server.Data{Key: key}
	byteData, err := json.Marshal(message)
	if err !=nil {
		log.Fatalln(err)
		return err
	}
	
	//create http request - not using http.Get because it doesn't take in body message
	request, err := http.NewRequest("DELETE", URI, bytes.NewBuffer(byteData))
	if err != nil {
		log.Fatalln(err)
		return err
	}
	
	//send request to server
	client := &http.Client{}
	resp, err := client.Do(request)
	
	//close body
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}()
	
	if err != nil {
		log.Fatalln(err)
	}
	
	//check status - report if errors exist
	if resp.StatusCode != http.StatusAccepted {
		err := fmt.Errorf("delete failed :'%v'", resp.StatusCode)
		log.Fatalln(err)
		return err
	}
	
	//turn  response into byte slice and then into string
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err == io.EOF{
		//"End of File" means end of response read, convert response to byte slice to string
		log.Println(string(respBytes))
		
		return err
	}
	if err != nil {
		log.Println(err)
		log.Fatalln(err)
		
		return err
	}
	
	//Delete command returns string to terminal; also log result on server side for verification
	log.Println(string(respBytes))
	fmt.Println(string(respBytes))
	
	return  nil
}