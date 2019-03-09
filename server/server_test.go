package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	
	"CacheCLI/kvcache"
)
 
func TestMockServer_Put(t *testing.T) {
	
	mockCache := kvcache.NewMockSimpleKVCache(true, "test")
	server := &Server{"8080", mockCache, nil}
	req, err := http.NewRequest("PUT", "/PUT", strings.NewReader(`{"key": "foo","value": "bar"}`))
	if err !=nil{
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	
	server.Put(rr, req)
	
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("PUT handler returned wrong status code: got '%v' want '%v'", status, http.StatusOK)
	}
	
	

}