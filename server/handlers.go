package server

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	
	"CacheCLI/kvcache"
)

//JSON literal and object for server to take in and return as needed
type Data struct {
	key string `json:"key"`
	value string`json:"value"`
}

//object that implements the handler methods
//has a port member to listen to a port
//has access to the KVC for a long running process
//has access to the JSON object for parsing in/out data as requested by the client
type Server struct {
	port string
	cache kvcache.KeyValueCache
	Data Data
}

//not sure I need this, but might for testing so keeping for now
//constructor function for generating data
func NewData(key, value string) *Data{
	return &Data{key, value}
}

var data = Data{}

func (s *Server) Put(w http.ResponseWriter, r *http.Request){
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
	//pass encoded json to cache for storage
	simpleKVC := s.cache.Create(data.key, data.value)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(simpleKVC); err !=nil{
		panic(err)
	}
}