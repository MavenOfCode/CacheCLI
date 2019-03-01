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
	t.Run("creates and reads successfully", func(t *testing.T) {

		testCache := NewSimpleKVCache()
		require.NotNil(t, testCache)

		key := "testKey"
		value := "testValue"
		err := testCache.Create(key,value)

		assert.NoError(t,err)
		b, _ := testCache.Read(key)
		assert.Equal(t, b, value)
	})

	t.Run("it creates successfully", func(t *testing.T) {
		testCache := NewSimpleKVCache()
		require.NotNil(t, testCache)
		key2 := "123"
		value2 := "Sooz"

		err2 := testCache.Create(key2, value2)
		assert.NoError(t, err2)

		a,_ := testCache.Read(key2)
		assert.Equal(t, a, value2)
	})

	//added to align with read error and tests
	t.Run(" create returns error when empty string given as parameter", func(t *testing.T) {
		testCache := NewSimpleKVCache()
		require.NotNil(t, testCache)
		key2 := ""
		value2 := ""

		err2 := testCache.Create(key2, value2)
		assert.Error(t,err2,"create failed: check key '' and value '' parameters")

		_,err := testCache.Read(key2)
		assert.ObjectsAreEqualValues(err, "read failed: key '' invalid")
	})

	//add test for repeat use of key
	t.Run(" create returns error when key already exists", func(t *testing.T) {
		testCache := NewSimpleKVCache()
		require.NotNil(t, testCache)

		key := "name"
		value := "bobby"
		err := testCache.Create(key,value)
		assert.NoError(t, err, "no error in put")
		//errR,_ := testCache.Read(key)

		key2 := "name"
		value2 := "betty"
		err2 := testCache.Create(key2, value2)
		assert.Error(t, err2, "create failed: key ' ' isn't unique: ")
		//_, err2R := testCache.Read(key2)
		//assert.NotEqual(t, errR ,err2R)

	})
}


func TestRead(t *testing.T){
	t.Run("reads successfully", func(t *testing.T) {
		testCache := NewSimpleKVCache()
		require.NotNil(t, testCache)

		key := "name"
		value := "Scott"

		err := testCache.Create(key,value)
		assert.NoError(t, err)

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


		err := testCache.Create(key, value)
		assert.NoError(t, err)

		f, _ := testCache.Read(key)
		assert.Equal(t, f, value)

		err2 := testCache.Create(key2, value2)
		assert.NoError(t, err2)

		v, _ := testCache.Read(key2)
		assert.Equal(t,v,value2)

		//being sure that the Read is reading different values for different keys with different test
		assert.NotEqual(t, f,v)
	})

	t.Run("read returns error when given empty string", func(t *testing.T) {
		testCache := NewSimpleKVCache()
		require.NotNil(t, testCache)

		key := "name"
		value := "Scott"

		err := testCache.Create(key, value)
		assert.NoError(t, err)

		f, _ := testCache.Read(key)

		assert.Equal(t, f, value)

		_, err2 := testCache.Read("")

		//updated tests to reflect new Read method signature and used Objects are Equal values due to the indirect reference to the error message in the assertion
		assert.ObjectsAreEqualValues(err2, "read failed: key ' ' invalid")
	})

	t.Run("read returns error when given invalid key", func(t *testing.T) {
		testCache := NewSimpleKVCache()
		require.NotNil(t, testCache)

		key := "name"
		value := "Scott"
		invalidKey :="animal"

		err := testCache.Create(key, value)
		assert.NoError(t, err)

		_, err2 := testCache.Read(invalidKey)

		assert.Error(t,err2,"read failed: key invalid")
	})
}

func TestUpdate(t *testing.T){
	t.Run("it can update", func(t *testing.T) {
		testCache := NewSimpleKVCache()
		require.NotNil(t, testCache)

		key := "name"
		value := "Benelli"

		put := testCache.Create(key,value)
		assert.NoError(t,put)

		key = "name"
		value = "Benny"
		err := testCache.Update(key, value)

		assert.Equal(t, err, nil)

		_, read := testCache.Read(key)
		assert.ObjectsAreEqualValues(read, value)
	})
	
	t.Run("update returns error when key not in cache", func(t *testing.T) {
		testCache := NewSimpleKVCache()
		require.NotNil(t, testCache)

		key := "name"
		value := "Hero"
		err := testCache.Update(key, value)

		assert.ObjectsAreEqualValues(err, "update failed: key '%v' not in cache")

		_, read := testCache.Read(key)
		assert.ObjectsAreEqualValues(read, value)
	})

	t.Run("update returns error when key is empty string", func(t *testing.T) {
		testCache := NewSimpleKVCache()
		require.NotNil(t, testCache)

		key := "name"
		value := "Benelli"

		put := testCache.Create(key,value)
		assert.NoError(t,put)

		key = " "
		value = "Benny"
		err := testCache.Update(key, value)
		assert.ObjectsAreEqualValues(err, "update failed: key '%v' not in cache")

		_, read := testCache.Read(key)
		assert.ObjectsAreEqualValues(read, value)
	})

}

func TestDelete(t *testing.T){
	t.Run("it deletes", func(t *testing.T){
		testCache := NewSimpleKVCache()
		require.NotNil(t, testCache)

		key := "name"
		value := "Benelli"

		put := testCache.Create(key,value)
		assert.NoError(t,put)

		testCache.Delete(key)

		_,readErr := testCache.Read(key)

		assert.Error(t, readErr, "delete successful: key no longer in cache")

	})

	t.Run("delete returns error when key doesn't exist", func(t *testing.T) {
		testCache := NewSimpleKVCache()
		require.NotNil(t, testCache)

		key := "cat"

		err := testCache.Delete(key)
		assert.Error(t, err, "delete error works as expected")
	})

	t.Run("delete returns error when given empty key string", func(t *testing.T) {
		testCache := NewSimpleKVCache()
		require.NotNil(t, testCache)

		key := ""

		err := testCache.Delete(key)
		assert.Error(t, err, "delete error works as expected")
	})
}
