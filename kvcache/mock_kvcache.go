package kvcache

import (
	"fmt"
	"strconv"
)

/* implementation of the KeyValueCache interface for testing */
type MockKeyValueCache struct {
	Success      bool
	ReturnString string
}

//constructor function useful for testing
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
