package cmd

import (
	"CacheCLI/kvcache"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
)

type CommandRunner struct {
	cache kvcache.KeyValueCache
}

func (c *CommandRunner) CreateCmd (cmd *cobra.Command, args []string) error {
	//pre-seeding cache for read command for now since cache won't persist until CLI/Cache connection built
	cache.Create("name", "harley")
	cache.Create("animal", "horse")
	cache.Create("kitten", "Bene")

	if len(args) < 2{
		return errors.New("create failed: insufficient arguments provided")
	}

	fmt.Println(cache)
	if  cache != nil {
		createResult := cache.Create(args[0],args[1])
		if createResult == nil {
			fmt.Printf("create success:  cache '%v' ", cache)
			fmt.Println()
			return nil
		}
	}
	return errors.New("create failed: cache not initialized")
}