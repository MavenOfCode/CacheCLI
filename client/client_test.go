package client

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
	clientCache := NewCacheClient()
	//client := &http.Client{}
	
	t.Run("create request works", func(t *testing.T){
		err := clientCache.Create("foo", "bar")
		require.Error(t, err)
		
		rr := httptest.NewRecorder()
		//status code passed back is correct too
		assert.Equal(t, rr.Code, http.StatusOK )
	})
}
