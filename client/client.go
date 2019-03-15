package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"CacheCLI/server"
)

type CacheClient struct {
	URI    string
	Client http.Client
}

func NewCacheClient() *CacheClient {
	const URI = "http://localhost:8080"
	return &CacheClient{URI: URI, Client: http.Client{}}
}

func (c *CacheClient) Create(key, value string) error {
	//create JSON with key and value by first putting into data object then marshal to JSON
	message := server.Data{Key: key, Value: value}
	byteData, err := json.Marshal(message)
	if err != nil {
		err := fmt.Errorf("server response: '%v' : bad input data", http.StatusBadRequest)
		return err
	}

	//create HTTP request
	request, err := http.NewRequest("PUT", c.URI, bytes.NewBuffer(byteData))
	if err != nil {
		err := fmt.Errorf("server response: '%v' : http request failed", http.StatusBadRequest)
		return err
	}

	//Make request to server and get a response
	client := &http.Client{}
	resp, err := client.Do(request)
	//close body
	defer func() {
		_ = resp.Body.Close()
	}()

	if err != nil {
		msg, _ := ioutil.ReadAll(resp.Body)
		err := fmt.Errorf("server response: '%v' : '%v'", http.StatusBadRequest, string(msg))
		return err
	}

	//check status to see if  sever error exists
	if resp.StatusCode != http.StatusCreated {
		msg, _ := ioutil.ReadAll(resp.Body)
		err := fmt.Errorf("server response: '%v' : '%v'", resp.StatusCode, string(msg))
		return err
	}

	return nil
}

func (c *CacheClient) Read(key string) (string, error) {
	//create JSON with key
	message := server.Data{Key: key}
	byteData, err := json.Marshal(message)
	if err != nil {
		err := fmt.Errorf("server response: '%v' : bad input data'", http.StatusBadRequest)
		return "", err
	}

	//create http request - not using http.Get because it doesn't take in body message
	request, err := http.NewRequest("GET", c.URI, bytes.NewBuffer(byteData))
	if err != nil {
		err := fmt.Errorf("server response: '%v' : http request failed", http.StatusBadRequest)
		return "", err
	}

	//send request to server
	client := &http.Client{}
	resp, err := client.Do(request)
	//close body
	defer func() {
		_ = resp.Body.Close()
	}()

	if err != nil {
		msg, _ := ioutil.ReadAll(resp.Body)
		err := fmt.Errorf("server response: '%v' : '%v'", http.StatusBadRequest, string(msg))
		return "", err
	}

	//check status - report if errors exist
	if resp.StatusCode != http.StatusOK {
		msg, _ := ioutil.ReadAll(resp.Body)
		err := fmt.Errorf("server response: '%v': '%v'", resp.StatusCode, string(msg))
		return "", err
	}

	//if no error, take response stream, convert reader to byte slice
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err == io.EOF {
		//"End of File" means end of response read, convert response to byte slice to string
		result := string(respBytes)
		return result, err
	}
	if err != nil {
		return "", err
	}

	//Read command returns a string (result of read query) and error
	return string(respBytes), nil
}

func (c *CacheClient) Update(key, value string) error {
	//creates JSON with key and value
	message := server.Data{Key: key, Value: value}
	byteData, err := json.Marshal(message)
	if err != nil {
		err := fmt.Errorf("server response: '%v' : bad input data'", http.StatusBadRequest)
		return err
	}

	//creates http request
	request, err := http.NewRequest("POST", c.URI, bytes.NewBuffer(byteData))
	if err != nil {
		err := fmt.Errorf("server response: '%v' : http request failed", http.StatusBadRequest)
		return err
	}

	//send request to server
	client := &http.Client{}
	resp, err := client.Do(request)
	defer func() {
		_ = resp.Body.Close()
	}()

	if err != nil {
		msg, _ := ioutil.ReadAll(resp.Body)
		err := fmt.Errorf("server response: '%v' : '%v'", http.StatusBadRequest, string(msg))
		return err
	}

	//check status - report if errors exist
	if resp.StatusCode != http.StatusCreated {
		msg, _ := ioutil.ReadAll(resp.Body)
		err := fmt.Errorf("server response: '%v': '%v'", resp.StatusCode, string(msg))
		return err
	}

	return nil
}

func (c *CacheClient) Delete(key string) error {
	//create JSON from key
	message := server.Data{Key: key}
	byteData, err := json.Marshal(message)
	if err != nil {
		err := fmt.Errorf("server response: '%v' : bad input data'", http.StatusBadRequest)
		return err
	}

	//create http request - not using http.Get because it doesn't take in body message
	request, err := http.NewRequest("DELETE", c.URI, bytes.NewBuffer(byteData))
	if err != nil {
		err := fmt.Errorf("server response: '%v' : http request failed", http.StatusBadRequest)
		return err
	}

	//send request to server
	client := &http.Client{}
	resp, err := client.Do(request)
	defer func() {
		_ = resp.Body.Close()
	}()

	if err != nil {
		msg, _ := ioutil.ReadAll(resp.Body)
		err := fmt.Errorf("server response: '%v' : '%v'", http.StatusBadRequest, string(msg))
		return err
	}

	//check status - report if errors exist
	if resp.StatusCode != http.StatusAccepted {
		msg, _ := ioutil.ReadAll(resp.Body)
		err := fmt.Errorf("server response: '%v': '%v'", resp.StatusCode, string(msg))
		return err
	}

	return nil
}
