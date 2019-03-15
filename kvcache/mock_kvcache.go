package kvcache

import (
	"fmt"
	"strconv"
)

/* MockCache struct that is a member of the  KVC interface for testing of KVC CLI commands */
type MockKeyValueCache struct {
	Success      bool
	ReturnString string
}

//go insists that the put method is required to implement the KVC interface for the NewMockSimpleKVCache constructor.
// ..I've invalidated the files/cache several times and it still requires this - so putting it in for now but NOT
// USED - can't find it's use ANYWHERE
func (m *MockKeyValueCache) Put(key, value string) error {
	m.Success, _ = strconv.ParseBool(key)
	m.ReturnString = value
	return nil
}

//constructor function for generating test MockCache
func NewMockSimpleKVCache(success bool, retString string) KeyValueCache {
	return &MockKeyValueCache{success, retString}
}

func (m *MockKeyValueCache) Create(key, value string) error {
	m.Success, _ = strconv.ParseBool(key)
	m.ReturnString = value
	return nil
}

func (m *MockKeyValueCache) Read(key string) (string, error) {
	if m == nil {
		return "", fmt.Errorf("update error: cache empty")
	}
	m.Success, _ = strconv.ParseBool(key)
	if m.Success {
		return m.ReturnString, nil
	}
	return "", fmt.Errorf("read error")
}

func (m *MockKeyValueCache) Update(key, value string) error {
	if m == nil {
		return fmt.Errorf("update error: cache empty")
	}
	m.Success, _ = strconv.ParseBool(key)
	if m.Success {
		return nil
	}
	return fmt.Errorf("update error")
}

func (m *MockKeyValueCache) Delete(key string) error {
	if m == nil {
		return fmt.Errorf("update error: cache empty")
	}
	m.Success, _ = strconv.ParseBool(key)
	if m.Success {
		return nil
	}
	return fmt.Errorf("delete error")
}
