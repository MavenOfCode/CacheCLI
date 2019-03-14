package client

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	
	"CacheCLI/kvcache"
)

func TestNewCacheClient(t *testing.T) {
	t.Run("it creates new cache", func(t *testing.T) {
		testCache := NewCacheClient()
		assert.NotNil(t, testCache)
	})
}

func TestCacheClient_Create(t *testing.T) {
	
	
	t.Run("create request works", func(t *testing.T){
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusCreated)
		}))
		clientCache := &CacheClient{
			URI: testServer.URL,
			Client: http.Client{},
		}
		defer 	testServer.Close()
		err := clientCache.Create("Key", "Value")
		assert.NoError(t, err)
	})
	
	t.Run("create request works", func(t *testing.T){
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
		}))
		clientCache := &CacheClient{
			URI: testServer.URL,
			Client: http.Client{},
		}
		defer 	testServer.Close()
		err := clientCache.Create("Key", "Value")
		assert.Error(t, err)
	})
	
	//t.Run("create returns error correctly when server returns error", func(t *testing.T) {
	//	err := clientCache.Create("", "Betty")
	//	fmt.Println(err.Error())
	//	assert.Error(t,err)
	//})
}


/* Mock Server content for testing purposes */
type MockServer struct {
	port string
	cache kvcache.KeyValueCache
	router *mux.Router
}

type Data struct {
	Key string `json:"key"`
	Value string`json:"value"`
}

const headerTypeKey = "Content-Type"
const headerValue = "application/json; charset=UTF-8"


func (s *MockServer) HandleData(w http.ResponseWriter, r *http.Request) (Data, error){
	var data = Data{}
	//if body is empty error out
	if r.Body == nil {
		w.WriteHeader(http.StatusNoContent)
		_, err := w.Write([]byte("body empty"))
		if err == nil {
			return data, err
		}
		return data, nil
	}
	//if ReadAll method errors out
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		_, err2 := w.Write([]byte(err.Error()))
		if err2 != nil {
			return data, err2
		}
		return data, err
	}
	//if request body is closed without content error out
	if err := r.Body.Close(); err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		_, err2 := w.Write([]byte(err.Error()))
		if err2 != nil {
			return data, err2
		}
		return data, err
	}
	//transform request from json; if json is not correctly configured error out
	if err := json.Unmarshal(body, &data); err !=nil {
		w.WriteHeader(http.StatusUnprocessableEntity)//unprocessable entity (json failed)
		_, err2 := w.Write( []byte (err.Error()))
		if err2 !=nil {
			return data, err2
		}
		return data, err
	}
	return data, nil
}

func (s *MockServer) Put(w http.ResponseWriter, r *http.Request){
	data, err := s.HandleData(w, r)
	//if handle data method has error and data is empty, error out
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err2 := w.Write([]byte(err.Error()))
		if err2 != nil {
			return
		}
	}
	//pass encoded json to cache for storage
	err = s.cache.Create(data.Key, data.Value)
	if err !=nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err2 := w.Write([]byte(err.Error()))
		if err2 != nil {
			return
		}
		return
	}
	w.Header().Set(headerTypeKey, headerValue)
	w.WriteHeader(http.StatusCreated)
	return
}

func (s *MockServer) Get(w http.ResponseWriter, r *http.Request){
	data, err := s.HandleData(w, r)
	//if handle data method has error and data is empty, error out
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err2 := w.Write([]byte(err.Error()))
		if err2 != nil {
			return
		}
		return
	}
	//pass unmarshalled json to cache for request of data to return
	readResult, err := s.cache.Read(data.Key)
	//if Read returns error return not found status from server to client
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		_, err2 := w.Write([]byte(err.Error()))
		if err2 != nil {
			return
		}
		return
	}
	//convert string into byte slice for writer to send content back to client
	_, err1 := w.Write([]byte(readResult))
	if err1 != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		_, err2 := w.Write([]byte(err.Error()))
		if err2 != nil {
			return
		}
		return
	}
	w.Header().Set(headerTypeKey, headerValue)
	return
}

func (s *MockServer) Post(w http.ResponseWriter, r *http.Request){
	data, err := s.HandleData(w, r)
	//if handle data method has error and data is empty, error out
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err2 := w.Write([]byte(err.Error()))
		if err2 != nil {
			return
		}
		return
	}
	//pass decoded json to cache for storage update
	err = s.cache.Update(data.Key, data.Value)
	//if Update returns error pass error back to client
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		_, err2 := w.Write([]byte(err.Error()))
		if err2 != nil {
			return
		}
		return
	}
	w.Header().Set(headerTypeKey,headerValue)
	w.WriteHeader(http.StatusCreated)
	return
}

func (s *MockServer) Delete(w http.ResponseWriter, r *http.Request){
	data, err := s.HandleData(w, r)
	//if handle data method has error and data is empty, error out
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err2 := w.Write([]byte(err.Error()))
		if err2 != nil {
			return
		}
		return
	}
	//pass decoded json to cache for request of data to return
	err = s.cache.Delete(data.Key)
	//if Delete returns error return not found status from server to client
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		_, err2 := w.Write([]byte(err.Error()))
		if err2 != nil {
			return
		}
		return
	}
	w.Header().Set(headerTypeKey, headerValue)
	w.WriteHeader(http.StatusAccepted)
	return
}