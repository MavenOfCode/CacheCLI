package cmd

import (
	_ "github.com/spf13/cobra"
)
//
//func (c *cmdRunner) CreateCmd (cmd *cobra.Command, args []string) error {
//	if cache == nil {
//		return errors.New("cache not initialized - create failed: ")
//	}
//	//pre-seeding cache for read command for now since cache won't persist until CLI/Cache connection built
//	cache.Create("name", "harley")
//	cache.Create("animal", "horse")
//	cache.Create("kitten", "Bene")
//
//	createResult := cache.Create(args[0],args[1])
//	if createResult == nil {
//		fmt.Printf("create success:  cache '%v' ", cache)
//		fmt.Println()
//		return nil
//	}
//	return errors.New("create fail")
//}