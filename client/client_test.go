package client

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewCacheClient(t *testing.T) {
	t.Run("it creates new cache", func(t *testing.T) {
		testCache := NewCacheClient()
		assert.NotNil(t, testCache)
	})
}

func TestCacheClient_Create(t *testing.T) {
	t.Run("create request works", func(t *testing.T) {
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusCreated)
		}))
		clientCache := &CacheClient{
			URI:    testServer.URL,
			Client: http.Client{},
		}
		defer testServer.Close()

		err := clientCache.Create("Key", "Value")

		assert.NoError(t, err)
	})

	t.Run("create request returns an error", func(t *testing.T) {
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusAccepted)
		}))
		clientCache := &CacheClient{
			URI:    testServer.URL,
			Client: http.Client{},
		}
		defer testServer.Close()

		err := clientCache.Create(" ", "Value")

		assert.Error(t, err)
	})
}

func TestCacheClient_Read(t *testing.T) {
	t.Run("read request works", func(t *testing.T) {
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))
		clientCache := &CacheClient{
			URI:    testServer.URL,
			Client: http.Client{},
		}
		defer testServer.Close()

		_, err := clientCache.Read("happy")

		assert.NoError(t, err)
	})

	t.Run("read request returns an error", func(t *testing.T) {
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusAccepted)
		}))
		clientCache := &CacheClient{
			URI:    testServer.URL,
			Client: http.Client{},
		}
		defer testServer.Close()

		_, err := clientCache.Read("")

		assert.Error(t, err)
	})
}

func TestCacheClient_Update(t *testing.T) {
	t.Run("update request works", func(t *testing.T) {
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusCreated)
		}))
		clientCache := &CacheClient{
			URI:    testServer.URL,
			Client: http.Client{},
		}
		defer testServer.Close()

		err := clientCache.Update("key", "lock")

		assert.NoError(t, err)
	})

	t.Run("update request works", func(t *testing.T) {
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusAccepted)
		}))
		clientCache := &CacheClient{
			URI:    testServer.URL,
			Client: http.Client{},
		}
		defer testServer.Close()

		err := clientCache.Update("", "lock")

		assert.Error(t, err)
	})
}

func TestCacheClient_Delete(t *testing.T) {
	t.Run("delete request works", func(t *testing.T) {
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusAccepted)
		}))
		clientCache := &CacheClient{
			URI:    testServer.URL,
			Client: http.Client{},
		}
		defer testServer.Close()

		err := clientCache.Delete("key")

		assert.NoError(t, err)
	})

	t.Run("delete request doesn't work", func(t *testing.T) {
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))
		clientCache := &CacheClient{
			URI:    testServer.URL,
			Client: http.Client{},
		}
		defer testServer.Close()

		err := clientCache.Delete("")

		assert.Error(t, err)
	})
}
