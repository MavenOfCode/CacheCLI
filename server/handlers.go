package server

import (
	"code.uber.internal/sooz/key-value/.tmp/.go/goroot/src/go/doc/testdata"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	
	"CacheCLI/kvcache"
)


func Put(w http.ResponseWriter, r *http.Request){
    var simpleKeyValueCache kvcache.SimpleKeyValueCache
	
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
	if err := json.Unmarshal(body, &simpleKeyValueCache); err !=nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)//unprocessable entity (json failed)
		if err := json.NewEncoder(w).Encode(err); err !=nil {
			panic(err)
		}
	}
	skvc := kvcache.NewSimpleKVCache(simpleKeyValueCache)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(skvc); err !=nil{
		panic(err)
	}
}