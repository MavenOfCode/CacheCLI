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
 
func TestServer_Put(t *testing.T) {
	
	mockCache := NewServerTestKeyValueCache("testKey", "testValue")
	server := &Server{"8080", mockCache, nil}
	
	t.Run("put works", func(t *testing.T) {
		req, err := http.NewRequest("PUT", "/PUT", strings.NewReader(`{"key": "foo","value": "bar"}`))
		require.NoError(t, err)
		
		rr := httptest.NewRecorder()
		server.Put(rr, req)
		assert.Equal(t, rr.Code, http.StatusCreated)
	})
	
	t.Run("put returns error when content is empty - like malformed JSON error", func(t *testing.T) {
		
		req, err := http.NewRequest("PUT", "PUT",strings.NewReader(""))
		require.NoError(t, err)
		
		rr := httptest.NewRecorder()
		server.Put(rr, req)
		assert.Equal(t, rr.Code,http.StatusUnprocessableEntity )
	})
	
	t.Run("put returns error when json malformed", func(t *testing.T) {
		req, err := http.NewRequest("PUT", "/PUT",strings.NewReader(`MALFORMED JSON`))
		require.NoError(t, err)
		
		rr := httptest.NewRecorder()
		server.Put(rr, req)
		assert.Equal(t, rr.Code, http.StatusUnprocessableEntity)
	})
}

func TestServer_Get(t *testing.T) {
	mockCache := kvcache.NewMockSimpleKVCache(true, "test")
	server := &Server{"8080", mockCache, nil}

	t.Run("get works", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/GET", strings.NewReader(`{"key": "success"}`))
		require.NoError(t, err)

		rr := httptest.NewRecorder()
		server.Get(rr, req)
		assert.Equal(t, rr.Code, http.StatusAccepted)
	})
	
	t.Run("get returns error when key doesn't exist in cache", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/GET", strings.NewReader(`{"success": "true"}`))
		require.NoError(t, err)
		
		rr := httptest.NewRecorder()
		server.Get(rr, req)
		assert.Equal(t, rr.Code, http.StatusNotFound)
		
	})

}

type ServerTestKeyValueCache struct {
	Key string `json:"key"`
	Value string `json:"value"`
}


func NewServerTestKeyValueCache(key, value string) kvcache.KeyValueCache {
	return &ServerTestKeyValueCache{key,value}
}

func (st *ServerTestKeyValueCache) Create(key, value string) error {
	st.Key = key
	st.Value = value
	return nil
}

func (st *ServerTestKeyValueCache) Read(key string) (string, error) {
	if st == nil {
		return "", fmt.Errorf("update error: cache empty")
	}
	if st.Key != ""{
		return st.Value, nil
	}
	return "", fmt.Errorf("read error")
}

func (st *ServerTestKeyValueCache) Update(key, value string) error {
	if st == nil {
		return fmt.Errorf("update error: cache empty")
	}
	st.Key = key
	st.Value  = value
	return nil
}

func (st *ServerTestKeyValueCache) Delete(key string) error {
	if st == nil {
		return fmt.Errorf("update error: cache empty")
	}
	if st.Key != "" {
		return nil
	}
	return fmt.Errorf("delete error")
}
