package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"CacheCLI/kvcache"
)

type CommandRunner struct {
	cache kvcache.KeyValueCache
}

/*Commands for CLI using CommandRunner*/
func (c *CommandRunner) CreateCmd(cmd *cobra.Command, args []string) (string, error) {
	if len(args) != 2 {
		return "", fmt.Errorf("create failed: insufficient arguments provided")
	}
	
	if c.cache != nil {
		err := c.cache.Create(args[0], args[1])
		if err == nil {
			res := fmt.Sprintf("create success: cache '%v'\n ", c.cache)
			return res, nil
		}
		return "", fmt.Errorf("create failed: '%v' ", err)
	}
	return "", fmt.Errorf("create failed: cache is nil")
}

func (c *CommandRunner) ReadCmd(cmd *cobra.Command, args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("read failed: only one argument required")
	}
	
	if c.cache == nil {
		return "", fmt.Errorf("read failed: cache nil")
	}
	
	readResult, err := c.cache.Read(args[0])
	if err != nil {
		return "", err
	}
	res := fmt.Sprintf(">> value for key '%v' is '%v'", args[0], readResult)
	return res, nil
}

func (c *CommandRunner) UpdateCmd(cmd *cobra.Command, args []string) (string, error) {
	if len(args) != 2 {
		return "", fmt.Errorf("update failed: two arguments required")
	}
	
	if c.cache == nil {
		return "", fmt.Errorf("update failed: cache is nil ")
	}
	
	err := c.cache.Update(args[0], args[1])
	if err == nil {
		res := fmt.Sprintf("update success: cache '%v' \n", c.cache)
		return res, nil
	}
	return "", fmt.Errorf("update failed: '%v'", err)
}

func (c *CommandRunner) DeleteCmd(cmd *cobra.Command, args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("delete failed: one argument required ")
	}
	
	if c.cache == nil {
		return "", fmt.Errorf("delete failed: cache is nil")
	}
	
	err := c.cache.Delete(args[0])
	if err == nil {
		res := fmt.Sprintf("delete success: cache '%v' \n", c.cache)
		
		return res, nil
	}
	return "", fmt.Errorf("delete failed: '%v'", err)
}
