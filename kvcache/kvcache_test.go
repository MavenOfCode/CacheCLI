package kvcache

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSimpleKeyValueCache(t *testing.T) {
	t.Run("it creates new cache", func(t *testing.T) {
		testCache := NewSimpleKVCache()
		assert.NotNil(t, testCache)
	})
}

func TestCreate(t *testing.T) {
	t.Run("create works to add k-v pair to cache", func(t *testing.T) {
		testCache := NewSimpleKVCache()
		require.NotNil(t, testCache)
		key := "123"
		value := "Sooz"

		err := testCache.Create(key, value)
		assert.NoError(t, err)

		a, _ := testCache.Read(key)
		assert.Equal(t, a, value)
	})

	t.Run("create returns an error when cache is nil", func(t *testing.T) {
		testCache := SimpleKeyValueCache{nil}
		key2 := "123"
		value2 := "Sooz"

		err := testCache.Create(key2, value2)
		assert.Error(t, err)
	})

	t.Run("create returns error when given empty key", func(t *testing.T) {
		testCache := NewSimpleKVCache()
		require.NotNil(t, testCache)
		key2 := ""
		value2 := "Value"

		err2 := testCache.Create(key2, value2)
		assert.Error(t, err2)
	})

	t.Run("create returns error when given empty value", func(t *testing.T) {
		testCache := NewSimpleKVCache()
		require.NotNil(t, testCache)
		key2 := "KeyTest"
		value2 := ""

		err2 := testCache.Create(key2, value2)
		assert.Error(t, err2)
	})

	t.Run("create returns error when key already exists", func(t *testing.T) {
		testCache := NewSimpleKVCache()
		require.NotNil(t, testCache)

		key := "name"
		value := "bobby"
		err := testCache.Create(key, value)
		assert.NoError(t, err)

		key2 := "name"
		value2 := "betty"
		err2 := testCache.Create(key2, value2)
		assert.Error(t, err2)
	})
}

func TestRead(t *testing.T) {
	t.Run("it reads", func(t *testing.T) {
		testCache := NewSimpleKVCache()
		require.NotNil(t, testCache)

		key := "name"
		value := "Scott"
		_ = testCache.Create(key, value)

		f, _ := testCache.Read(key)
		assert.Equal(t, f, value)
	})

	t.Run("read successfully when given different keys", func(t *testing.T) {
		testCache := NewSimpleKVCache()
		require.NotNil(t, testCache)

		key := "name"
		value := "Benelli"

		key2 := "nickname"
		value2 := "Benny"

		_ = testCache.Create(key, value)
		_ = testCache.Create(key2, value2)

		v, _ := testCache.Read(key2)
		assert.Equal(t, v, value2)
	})

	t.Run("read returns error when given empty string", func(t *testing.T) {
		testCache := NewSimpleKVCache()
		require.NotNil(t, testCache)

		key := "name"
		value := "Scott"

		_ = testCache.Create(key, value)

		_, err := testCache.Read("")
		assert.Error(t, err)
	})

	t.Run("read returns error when given non-existent key", func(t *testing.T) {
		testCache := NewSimpleKVCache()
		require.NotNil(t, testCache)

		key := "name"
		value := "Scott"
		notExistKey := "animal"
		_ = testCache.Create(key, value)

		_, err2 := testCache.Read(notExistKey)
		assert.Error(t, err2)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("it can update", func(t *testing.T) {
		testCache := NewSimpleKVCache()
		require.NotNil(t, testCache)

		key := "name"
		value := "Benelli"
		_ = testCache.Create(key, value)

		key = "name"
		value = "Benny"
		err := testCache.Update(key, value)
		assert.Nil(t, err)

		read, _ := testCache.Read(key)
		assert.Equal(t, read, value)
	})

	t.Run("update returns error when key not in cache", func(t *testing.T) {
		testCache := NewSimpleKVCache()
		require.NotNil(t, testCache)

		key := "name"
		value := "Hero"
		err := testCache.Update(key, value)
		assert.Error(t, err)
	})

	t.Run("update returns error when key is empty string", func(t *testing.T) {
		testCache := NewSimpleKVCache()
		require.NotNil(t, testCache)

		key := ""
		value := "Benny"
		err := testCache.Update(key, value)
		assert.Error(t, err)
	})

}

func TestDelete(t *testing.T) {
	t.Run("it deletes", func(t *testing.T) {
		testCache := NewSimpleKVCache()
		require.NotNil(t, testCache)

		key := "name"
		value := "Benelli"

		_ = testCache.Create(key, value)

		err := testCache.Delete(key)
		assert.Nil(t, err)

		_, err = testCache.Read(key)
		assert.Error(t, err)

	})

	t.Run("delete returns error when key doesn't exist", func(t *testing.T) {
		testCache := NewSimpleKVCache()
		require.NotNil(t, testCache)

		key := "cat"

		err := testCache.Delete(key)
		assert.Error(t, err)
	})

	t.Run("delete returns error when given empty key string", func(t *testing.T) {
		testCache := NewSimpleKVCache()
		require.NotNil(t, testCache)

		key := ""

		err := testCache.Delete(key)
		assert.Error(t, err)
	})
}
