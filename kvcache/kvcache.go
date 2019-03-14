package kvcache

import (
	"fmt"
)

//interface for use by all files (public by using cap at start of name)
type KeyValueCache interface{
	Create(key, value string) error
	Read(key string) (string,error)
	Update(key, value string) error
	Delete(key string) error
}

type SimpleKeyValueCache struct{
	Data map[string]string
}

//constructor function for generating cache
func NewSimpleKVCache() *SimpleKeyValueCache{
	return &SimpleKeyValueCache{map[string]string{}}
}

/*working implementation of KVC interface*/
func (c *SimpleKeyValueCache) Create(key, value string) error{
	if c.Data == nil {
		return fmt.Errorf("create failed: cache does not exist")
	}

	if key == "" || value == "" {
		return fmt.Errorf("create failed: key '%v' and value '%v' must not be empty strings ",key, value)
	}

	if _, ok := c.Data[key]; ok {
		return fmt.Errorf("create failed: key '%v' isn't unique: ", key)
	}

	c.Data[key] = value
	fmt.Print(c)
	return nil
}

func (c *SimpleKeyValueCache) Read(key string) (string,error){
	result, ok := c.Data[key]
	if !ok {
		return "",fmt.Errorf("read failed: key '%v' not in cache", key)
	}
	return result, nil
}

func (c *SimpleKeyValueCache) Update(key, value string) error{
	_, keyExists := c.Data[key]
	if keyExists {
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
