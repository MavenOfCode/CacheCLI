package server

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	
	"CacheCLI/kvcache"
)
 
func TestServer_Put(t *testing.T) {
	
	mockCache := kvcache.NewMockSimpleKVCache(true, "test")
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

//func TestServer_Get(t *testing.T) {
//	mockCache := kvcache.NewMockSimpleKVCache(true, "test")
//	server := &Server{"8080", mockCache, nil}
//
//	t.Run("get works", func(t *testing.T) {
//		req, err := http.NewRequest("GET", "/GET", strings.NewReader(`{"key": "foo","value": "bar"}`))
//		require.NoError(t, err)
//
//		rr := httptest.NewRecorder()
//		server.Get()
//
//	})
//
//}