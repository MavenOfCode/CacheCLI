package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
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

 var headerTypeKey = "Content-Type"
 var headerValue = "application/json; charset=UTF-8"

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
		w.Header().Set("Status: ", "http.StatusBadRequest")
		return
	}
	//if request body is closed without content error out
	if err := r.Body.Close(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//transform request to json; if json is not correctly configured error out
	if err := json.Unmarshal(body, &data); err !=nil {
		w.Header().Set(headerTypeKey, headerValue )
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
		w.Header().Set("Status: ", "http.StatusBadRequest")
		return
	}
	if err := r.Body.Close(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Status: ", "http.StatusBadRequest")
		return
	}
	//transform request to json; if json is not correctly configured error out
	if err := json.Unmarshal(body, &data); err !=nil {
		w.Header().Set(headerTypeKey, headerValue)
		w.WriteHeader(http.StatusUnprocessableEntity)//unprocessable entity (json failed)
		return
	}
	//pass encoded json to cache for request of data to return
	readReq, err := s.cache.Read(data.Key)
	//if Read returns error return not found status from server to client
	if err != nil {
		w.Header().Set(headerTypeKey, headerValue)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	//if Read returns string (and error not nil) then encode response for return to client
	if err := json.NewEncoder(w).Encode(readReq); err != nil {
		w.Header().Set(headerTypeKey, headerValue)
		w.WriteHeader(http.StatusOK)
		return
	}
}

func (s *Server) Post(w http.ResponseWriter, r *http.Request){
	var data = Data{}
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
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
	//transform request to json; if json is not correctly configured error out
	if err := json.Unmarshal(body, &data); err !=nil {
		w.Header().Set(headerTypeKey, headerValue)
		w.WriteHeader(http.StatusUnprocessableEntity)//unprocessable entity (json failed)
		return
	}
	//pass encoded json to cache for storage update
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
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	//if request is empty error out
	if err != nil {
		panic(err)
	}
	//if request body is empty error out
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	//transform request to json; if json is not correctly configured error out
	if err := json.Unmarshal(body, &data); err !=nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)//unprocessable entity (json failed)
		if err := json.NewEncoder(w).Encode(err); err !=nil {
			panic(err)
		}
	}
	//pass encoded json to cache for request of data to return
	deleteKVC := s.cache.Delete(data.Key)
	//if Delete returns error return not found status from server to client
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	//if Read returns string (and error not nil) then encode response for return to client
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusAccepted)
	if err := json.NewEncoder(w).Encode(deleteKVC); err != nil {
		panic(err)
	}
}