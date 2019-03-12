package server

import (
	"encoding/json"
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

//not sure I need this, but might for testing so keeping for now
//constructor function for generating data
func NewData(key, value string) *Data{
	return &Data{key, value}
}

func (s * Server) Put(w http.ResponseWriter, r *http.Request){
	var data = Data{}
	body, err := ioutil.ReadAll(r.Body)
	//if body is empty error out
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//if request body is closed without content error out
	if err := r.Body.Close(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//transform request to json; if json is not correctly configured error out
	if err := json.Unmarshal(body, &data); err !=nil {
		w.WriteHeader(http.StatusUnprocessableEntity)//unprocessable entity (json failed)
		return
	}
	//pass encoded json to cache for storage
	err = s.cache.Create(data.Key, data.Value)
	if err !=nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set(headerTypeKey, headerValue)
	w.WriteHeader(http.StatusCreated)
	return
}

func (s *Server) Get(w http.ResponseWriter, r *http.Request){
	var data = Data{}
	body, err := ioutil.ReadAll(r.Body)
	//if request is empty error out
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := r.Body.Close(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//transform request to json; if json is not correctly configured error out
	if err := json.Unmarshal(body, &data); err !=nil {
		w.WriteHeader(http.StatusUnprocessableEntity)//unprocessable entity (json failed)
		return
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
		return
	}
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