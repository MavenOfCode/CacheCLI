package server

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	
	"CacheCLI/kvcache"
)


func TestNewServerTestKeyValueCache(t *testing.T){
	t.Run("it creates a new mock server cache", func(t *testing.T) {
		testServerCache := NewServerTestKeyValueCache("key", "value");
		assert.NotNil(t, testServerCache)
	})
}

func TestServer_HandleData(t *testing.T) {
	mockCache := NewServerTestKeyValueCache("testKey", "testValue")
	server := &Server{"8080", mockCache, nil}
	
	t.Run("handle data returns data", func(t *testing.T) {
		req, err := http.NewRequest("PUT", "/", strings.NewReader(`{"key": "foo","value": "bar"}`))
		require.NoError(t, err)
		
		rr := httptest.NewRecorder()
		result, err := server.HandleData(rr, req)
		assert.Equal(t, result, Data{ "foo","bar"})
		//status code passed back is correct too
		assert.Equal(t, rr.Code, http.StatusOK )
	})
	
	t.Run("handle data returns error if body is nil", func(t *testing.T) {
		req, err := http.NewRequest("PUT", "/", nil)
		require.NoError(t, err)
		
		rr := httptest.NewRecorder()
		result, err := server.HandleData(rr, req)
		assert.Equal(t, result, Data{})
		//confirm error returned
		assert.Error(t, err)
		//status code passed back is correct too
		assert.Equal(t, rr.Code, http.StatusNoContent)
	})
	
	t.Run("handle data returns error if JSON is MALFORMED", func(t *testing.T) {
		req, err := http.NewRequest("PUT", "/", strings.NewReader(`MALFORMED JSON`))
		require.NoError(t, err)
		
		rr := httptest.NewRecorder()
		result, err := server.HandleData(rr, req)
		assert.Equal(t, result, Data{})
		//confirm error returned
		assert.Error(t, err)
		//status code passed back is correct too
		assert.Equal(t, rr.Code, http.StatusUnprocessableEntity)
		
	})
	
}

func TestServer_Put(t *testing.T) {
	mockCache := NewServerTestKeyValueCache("testKey", "testValue")
	server := &Server{"8080", mockCache, nil}
	
	t.Run("put works", func(t *testing.T) {
		req, err := http.NewRequest("PUT", "/", strings.NewReader(`{"key": "foo","value": "bar"}`))
		require.NoError(t, err)
		
		rr := httptest.NewRecorder()
		server.Put(rr, req)
		assert.Equal(t, rr.Code, http.StatusCreated)
	})
	
	t.Run("put returns error when content is empty - like malformed JSON error", func(t *testing.T) {
		req, err := http.NewRequest("PUT", "/", strings.NewReader(""))
		require.NoError(t, err)
		
		rr := httptest.NewRecorder()
		server.Put(rr, req)
		assert.Equal(t, rr.Code,http.StatusUnprocessableEntity )
	})
	
	t.Run("put returns error when json malformed", func(t *testing.T) {
		req, err := http.NewRequest("PUT", "/", strings.NewReader(`MALFORMED JSON`))
		require.NoError(t, err)
		
		rr := httptest.NewRecorder()
		server.Put(rr, req)
		assert.Equal(t, rr.Code, http.StatusUnprocessableEntity)
	})
	
	t.Run("put returns error when key is empty", func(t *testing.T) {
		req, err := http.NewRequest("PUT", "/", strings.NewReader(`{"key": "","value": "bar"}`) )
		require.NoError(t, err)
		
		rr := httptest.NewRecorder()
		server.Put(rr, req)
		assert.Equal(t, rr.Code, http.StatusBadRequest )
	})
	
	t.Run("put returns error when value is empty", func(t *testing.T) {
		req, err := http.NewRequest("PUT", "/", strings.NewReader(`{"key": "foo","value": ""}`) )
		require.NoError(t, err)
		
		rr := httptest.NewRecorder()
		server.Put(rr, req)
		assert.Equal(t, rr.Code, http.StatusBadRequest )
	})
}

func TestServer_Get(t *testing.T) {
	mockCache := NewServerTestKeyValueCache("testKey", "testValue")
	server := &Server{"8080", mockCache, nil}

	t.Run("get works and returns correct status", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/", strings.NewReader(`{"key": "testKey"}`))
		require.NoError(t, err)

		rr := httptest.NewRecorder()
		server.Get(rr, req)
		assert.Equal(t, rr.Code, http.StatusOK)
	})
	
	t.Run("get works and returns correct value", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/", strings.NewReader(`{"key": "testKey"}`))
		require.NoError(t, err)
		
		rr := httptest.NewRecorder()
		server.Get(rr, req)
		
		actual := "testValue"
		expected := rr.Body.String()
		assert.Equal(t, expected, actual)
	})
	
	
	t.Run("get returns error when key doesn't exist in cache", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/", strings.NewReader(`{"key": "true"}`))
		require.NoError(t, err)
		
		rr := httptest.NewRecorder()
		server.Get(rr, req)
		assert.Equal(t, rr.Code, http.StatusNotFound)
	})
	
	t.Run("get returns error when JSON MALFORMED", func(t *testing.T) {
		req, err := http.NewRequest("GET","/", strings.NewReader(`MALFORMED JSON`))
		require.NoError(t, err)
		
		rr := httptest.NewRecorder()
		server.Get(rr, req)
		assert.Equal(t, rr.Code, http.StatusUnprocessableEntity)
	})
}

func TestServer_Post(t *testing.T){
	mockCache := NewServerTestKeyValueCache("testKey", "testValue")
	server := &Server{"8080", mockCache, nil}
	
	t.Run("post works", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/", strings.NewReader(`{"key": "testKey","value": "fooBar"}`))
		require.NoError(t, err)
		
		rr := httptest.NewRecorder()
		server.Post(rr, req)
		assert.Equal(t, rr.Code, http.StatusCreated)
	})
	
	t.Run("post returns error when key doesn't exist in cache", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/", strings.NewReader(`{"key": "zuperTrooper","value": "fooBar"}`))
		require.NoError(t, err)
		
		rr := httptest.NewRecorder()
		server.Post(rr, req)
		assert.Equal(t, rr.Code, http.StatusBadRequest)
	})
	
	t.Run("post returns error when JSON MALFORMED", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/", strings.NewReader(`MALFORMED JSON`))
		require.NoError(t, err)
		
		rr := httptest.NewRecorder()
		server.Post(rr, req)
		assert.Equal(t, rr.Code, http.StatusUnprocessableEntity)
		
	})
	
	t.Run("post returns error when content is empty - like malformed JSON error", func(t *testing.T) {
		
		req, err := http.NewRequest("POST", "/", strings.NewReader(""))
		require.NoError(t, err)
		
		rr := httptest.NewRecorder()
		server.Post(rr, req)
		assert.Equal(t, rr.Code,http.StatusUnprocessableEntity )
	})
	
	t.Run("post returns error when key is empty", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/", strings.NewReader(`{"key": "","value": "bar"}`) )
		require.NoError(t, err)
		
		rr := httptest.NewRecorder()
		server.Post(rr, req)
		assert.Equal(t, rr.Code, http.StatusBadRequest )
	})
	
	t.Run("post returns error when value is empty", func(t *testing.T) {
		req, err := http.NewRequest("PUT", "/", strings.NewReader(`{"key": "foo","value": ""}`) )
		require.NoError(t, err)
		
		rr := httptest.NewRecorder()
		server.Post(rr, req)
		assert.Equal(t, rr.Code, http.StatusBadRequest )
	})
}

func TestServer_Delete(t *testing.T) {
	mockCache := NewServerTestKeyValueCache("testKey", "testValue")
	server := &Server{"8080", mockCache, nil}
	
	t.Run("delete works", func(t *testing.T) {
		req, err := http.NewRequest("DELETE", "/", strings.NewReader(`{"key": "testKey"}`))
		require.NoError(t, err)
		
		rr := httptest.NewRecorder()
		server.Delete(rr, req)
		assert.Equal(t, rr.Code, http.StatusAccepted)
	})
	
	t.Run("delete returns error when JSON MALFORMED", func(t *testing.T) {
		req, err := http.NewRequest("DELETE", "/", strings.NewReader(`MALFORMED JSON`))
		require.NoError(t, err)
		
		rr := httptest.NewRecorder()
		server.Delete(rr, req)
		assert.Equal(t, rr.Code, http.StatusUnprocessableEntity)
	})
	
	t.Run("delete returns error when key is empty", func(t *testing.T) {
		req, err := http.NewRequest("DELETE", "/", strings.NewReader(`{"key": ""}`))
		assert.NoError(t, err)
		
		rr := httptest.NewRecorder()
		server.Delete(rr, req)
		assert.Equal(t, rr.Code, http.StatusNotFound)
	})
	
	t.Run("delete returns error when key is not in cache", func(t *testing.T) {
		req, err := http.NewRequest("DELETE", "/", strings.NewReader(`{"key": "foo"}`))
		assert.NoError(t, err)
		
		rr := httptest.NewRecorder()
		server.Delete(rr, req)
		assert.Equal(t, rr.Code, http.StatusNotFound)
	})
}

/*Mock Cache structure specifically to test Server handler function implementation*/

type ServerTestKeyValueCache struct {
	Key string `json:"key"`
	Value string `json:"value"`
}

func NewServerTestKeyValueCache(key, value string) kvcache.KeyValueCache {
	return &ServerTestKeyValueCache{key,value}
}

func (st *ServerTestKeyValueCache) Create(key, value string) error {
	if key == "" || value == ""{
		return fmt.Errorf("update error: key or value is empty")
	}
	st.Key = key
	st.Value = value
	return nil
}

func (st *ServerTestKeyValueCache) Read(key string) (string, error) {
	if st == nil {
		return "", fmt.Errorf("read error: cache empty")
	}
	if key == ""{
		return "", fmt.Errorf("read error: key is empty")
	}
	if key != st.Key {
		return "", fmt.Errorf("read error: key not in cache")
	}
	return st.Value, nil
	
}

func (st *ServerTestKeyValueCache) Update(key, value string) error {
	if st == nil {
		return fmt.Errorf("update error: cache empty")
	}
	if key == "" || value == ""{
		return fmt.Errorf("update error: key or value is empty")
	}
	if key != st.Key {
		return fmt.Errorf("update error: key is not in cache")
	}
	st.Key = key
	st.Value  = value
	return nil
}

func (st *ServerTestKeyValueCache) Delete(key string) error {
	if st == nil {
		return fmt.Errorf("delete error: cache empty")
	}
	if key == "" {
		return fmt.Errorf("delete error: key is empty")
	}
	if key != st.Key {
		return fmt.Errorf("delete error: key is not in cache")
	}
	st.Key = key
	st.Key = ""
	st.Value = ""
	return nil
}
