package kvcache

import (
	"fmt"
)

//interface for use by all files (public by using cap at start of name)
type KeyValueCache interface{
	Create(key, value string) error
	Read(key string) (string,error)
	Update(key,value string) error
	Delete(key string) error
}


type SimpleKeyValueCache struct{
	Data map[string]string
}

//constructor function for generating cache
func NewSimpleKVCache() *SimpleKeyValueCache{
	return &SimpleKeyValueCache{map[string]string{}}
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

/*working implementation of KVC interface*/
func (c *SimpleKeyValueCache) Create(key,value string) error{
	
	if c == nil {
		c.Data[key] = value
		return nil
		
	}

	//added if statement to match read behavior and logic for empty string
	if key =="" || value =="" {
		return fmt.Errorf("create failed: check key '%v' and value '%v' parameters  ",key, value)
	}

	//added to check if key exists and reject put if key does already exist
	if _, ok := c.Data[key]; ok {
		return fmt.Errorf("create failed: key '%v' isn't unique: ", key)
	}
	c.Data[key] = value

	return nil
}

//updated interface and method to return both string and error when realized SKVC wouldn't return an error when an empty string was entered as a key - not cool
func (c *SimpleKeyValueCache) Read(key string) (string,error){
	err := c.Data[key]
	if err == ""{
		return "",fmt.Errorf("read failed: key '%v' invalid or cache empty", key)
	}
	return err, nil
}

func (c *SimpleKeyValueCache) Update(key, value string) error{
	_, keyExists := c.Data[key]
	if keyExists == false {
		c.Data[key] = value
		return nil
	}
	return fmt.Errorf("update failed: key '%v' not in cache", key)
}

func (c *SimpleKeyValueCache) Delete(key string) error{
	_, keyExist := c.Data[key]
	if keyExist{
		delete(c.Data, key)
		return nil
	}
	return fmt.Errorf("delete failed: key '%v' not in cache",key)
}
