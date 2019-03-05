package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
)

func (c *CommandRunner) DeleteCmd(cmd *cobra.Command, args []string) error {
	//pre-seeding cache for read command for now since cache won't persist until CLI/Cache connection built
	//cache.Create("name", "harley")
	//cache.Create("animal", "horse")
	//cache.Create("kitten", "Bene")

	if len(args) < 1 {
		return errors.New("delete failed: at least one argument required ")
	}

	if cache == nil {
		return errors.New("delete failed: cache not initialized - delete failed")
	}

	deleteResult := cache.Delete(args[0])
	if deleteResult == nil {
		fmt.Printf("delete success: cache '%v' ", cache)
		fmt.Println()
		return nil
	}
	fmt.Println(deleteResult)
	return errors.New("")
}