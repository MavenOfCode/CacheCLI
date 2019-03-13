package server

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	
	"CacheCLI/kvcache"
)

//JSON literal and object for server to take in and return as needed
type Data struct {
	Key string `json:"key"`
	Value string`json:"value"`
}

//func (d *Data) Read()

//object that implements the handler methods
//has a port member to listen to a port
//has access to the KVC for a long running process
type Server struct {
	port string
	cache kvcache.KeyValueCache
	router *mux.Router
}

 const headerTypeKey = "Content-Type"
 const headerValue = "application/json; charset=UTF-8"

func StartServer(port string) {
	server := &Server{port:port, cache:kvcache.NewSimpleKVCache()}
	
	//moved all router construction components inside this constructor so it can access the server instance and its
	// methods
	routes := Routes{
		Route{
			"PUT",
			"/",
			server.Put,
		},
		Route{
			"POST",
			"/",
			server.Post,
		},
		Route{
			"GET",
			"/",
			server.Get,
		},
		Route{
			"DELETE",
			"/",
			server.Delete,
		},
	}
	
	router := mux.NewRouter().StrictSlash(true)
	for _,route := range routes{
		var handler http.Handler
		
		handler = route.HandlerFuc
		//put logger here later when I get to that
		
		router.
			Methods(route.Method).
			Path(route.URI).
			Handler(handler)
	}
	log.Fatal(http.ListenAndServe(port, router))
}

func (s *Server) HandleData(w http.ResponseWriter, r *http.Request) (Data, error){
	var data = Data{}
	//if body is empty error out
	if r.Body == nil {
		w.WriteHeader(http.StatusNoContent)
		result := []byte("body empty")
		response, err2 := w.Write(result)
		if err2 != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			return data, err2
		}
		w.WriteHeader(response)
		return data, fmt.Errorf("error with data request: '%v'", result)
	}
	//if ReadAll method errors out
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		result := []byte(err.Error())
		response, err2 := w.Write(result)
		if err2 != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			return data, err2
		}
		w.WriteHeader(response)
		return data, fmt.Errorf(err.Error())
	}
	//if request body is closed without content error out
	if err := r.Body.Close(); err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		result := []byte("body won't close")
		response, err2 := w.Write(result)
		if err2 != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			return data, err2
		}
		w.WriteHeader(response)
		return data, fmt.Errorf("error with data request: '%v'", result)
	}
	//transform request from json; if json is not correctly configured error out
	if err := json.Unmarshal(body, &data); err !=nil {
		w.WriteHeader(http.StatusUnprocessableEntity)//unprocessable entity (json failed)
		result := []byte (err.Error())
		response, err2 := w.Write(result)
		if err2 !=nil {
			w.WriteHeader(http.StatusExpectationFailed)
			return data, err2
		}
		w.WriteHeader(response)
		return data, err
	}
	return data, nil
}

func (s *Server) Put(w http.ResponseWriter, r *http.Request){
	data, err := s.HandleData(w, r)
	fmt.Println(w)
	//if handle data method has error and data is empty, error out
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err2 := w.Write([]byte(err.Error()))
		if err2 != nil {
			w.WriteHeader(http.StatusExpectationFailed)
			return
		}
	}
	//pass encoded json to cache for storage
	err = s.cache.Create(data.Key, data.Value)
	if err !=nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(w)
		_, err2 := w.Write([]byte(err.Error()))
		if err2 != nil {
			w.WriteHeader(http.StatusExpectationFailed)
			return
		}
		return
	}
	w.Header().Set(headerTypeKey, headerValue)
	w.WriteHeader(http.StatusCreated)
	return
}

func (s *Server) Get(w http.ResponseWriter, r *http.Request){
	data, err := s.HandleData(w, r)
	//if handle data method has error and data is empty, error out
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err2 := w.Write([]byte(err.Error()))
		if err2 != nil {
			w.WriteHeader(http.StatusExpectationFailed)
			return
		}
	}
	//pass unmarshalled json to cache for request of data to return
	readResult, err := s.cache.Read(data.Key)
	//if Read returns error return not found status from server to client
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	//convert string into byte slice for writer to send content back to client
	result := []byte(readResult)
	response, err := w.Write(result)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	w.Header().Set(headerTypeKey, headerValue)
	w.WriteHeader(response)
	return
}

func (s *Server) Post(w http.ResponseWriter, r *http.Request){
	var data = Data{}
	body, err := ioutil.ReadAll(r.Body)
	//if request is empty error out
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//if request body is empty error out
	if err := r.Body.Close(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//transform request from json; if json is not correctly configured error out
	if err := json.Unmarshal(body, &data); err !=nil {
		w.WriteHeader(http.StatusUnprocessableEntity)//unprocessable entity (json failed)
		return
	}
	//pass decoded json to cache for storage update
	err = s.cache.Update(data.Key, data.Value)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set(headerTypeKey,headerValue)
	w.WriteHeader(http.StatusCreated)
	return
}

func (s *Server) Delete(w http.ResponseWriter, r *http.Request){
	var data = Data{}
	body, err := ioutil.ReadAll(r.Body)
	//if request is empty error out
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//if request body is empty error out
	if err := r.Body.Close(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//transform request from json; if json is not correctly configured error out
	if err := json.Unmarshal(body, &data); err !=nil {
		w.WriteHeader(http.StatusUnprocessableEntity)//unprocessable entity (json failed)
		return
	}
	//pass decoded json to cache for request of data to return
	err = s.cache.Delete(data.Key)
	//if Delete returns error return not found status from server to client
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set(headerTypeKey, headerValue)
	w.WriteHeader(http.StatusAccepted)
	return
	
	
}