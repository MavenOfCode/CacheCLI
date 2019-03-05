package cmd

import (
	"CacheCLI/kvcache"
	"fmt"
)

type CommandRunner struct {
	cache kvcache.KeyValueCache
}

/* MockCache struct and implementation of KVC interface for testing of KVC CLI commands */
type MockKeyValueCache struct{
	WillFail bool
	ReturnString string
}

func (m *MockKeyValueCache) Create(key,value string) error{
	return nil
}

func (m *MockKeyValueCache) Read(key string) (string,error){
	if m.WillFail {
		return m.ReturnString, nil
	}
	return "", fmt.Errorf("read error")
}

func (m *MockKeyValueCache) Update(key, value string) error  {
	return nil
}

func (m *MockKeyValueCache) Delete(key string) error{
	return nil
}

//constructor function for generating test MockCache
func NewMockSimpleKVCache() *MockKeyValueCache{
	return &MockKeyValueCache{}
}
