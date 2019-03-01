package cmd

import (
	"KVCache/kvcache"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

//use constructor from kvcache package for global access to struct and it's private fields per Troy
var cache = kvcache.NewSimpleKVCache()

//make root command not executable without subcommand by not providing a 'Run' for the 'rootCmd'
var RootCmd = &cobra.Command{Use:"cli"}
var putCmd = &cobra.Command{
	Use:   "put",
	Args: cobra.MinimumNArgs(2),
	Short: "put key-value pair",
	Long:  "put key value strings into the key-value cache",
	RunE: func(cmd *cobra.Command, args []string) error {
		if cache == nil {
			return errors.New("cache not initialized - put failed: ")
		}
		putResult := cache.Put(args[0],args[1])
		if putResult == nil {
			fmt.Printf("put success:  cache '%v' ", cache)
			fmt.Println()
			return nil
		}
		return errors.New("put fail")
	},
}

//trying use of minimum args in command to avoid writing RunE function with error to test for args length
var readCmd = &cobra.Command{
	Use:  "read",
	Short: "read given key and return value",
	Args: cobra.MinimumNArgs(1),
	Long: "read value string out to command line from key-value cache given key string input from command line",
	RunE: func(cmd *cobra.Command, args []string) error  {
		fmt.Println(args, len(args))
		if cache == nil {
			return errors.New("cache empty - read failed: ")
		}
		//pre-seeding cache for read command for now since cache won't persist until CLI/Cache connection built
		cache.Put("name", "harley")
		cache.Put("animal", "horse")
		cache.Put("kitten", "Bene")
		readResult, err := cache.Read(args[0])
		fmt.Println(readResult)
		fmt.Println(err)
		if err !=nil {
			return err
		}
		fmt.Println(">>", readResult)
		return nil
	},
}

var updateCmd = &cobra.Command{
	Use:  "update",
	Args: cobra.MinimumNArgs(2),
	Short: "update key-value pair",
	Long:  "update key value strings into the key-value cache",
	RunE: func(cmd *cobra.Command, args []string) error {
		if cache == nil {
			return errors.New("cache not initialized - put failed: ")
		}
		//pre-seeding cache for read command for now since cache won't persist until CLI/Cache connection built
		cache.Put("name", "harley")
		cache.Put("animal", "horse")
		cache.Put("kitten", "Bene")
		updateResult := cache.Update(args[0],args[1])
		if updateResult == nil {
			fmt.Printf("update success:  cache '%v' ", cache)
			fmt.Println()
			return nil
		}
		return errors.New("update fail")
	},
}

var deleteCmd = &cobra.Command{
	Use:  "delete",
	Args: cobra.MinimumNArgs(1),
	Short: "delete key-value pair",
	Long:  "delete key value strings into the key-value cache",
	RunE: func(cmd *cobra.Command, args []string) error {
		if cache == nil {
			return errors.New("cache not initialized - put failed: ")
		}
		//pre-seeding cache for read command for now since cache won't persist until CLI/Cache connection built
		cache.Put("name", "harley")
		cache.Put("animal", "horse")
		cache.Put("kitten", "Bene")
		deleteResult := cache.Delete(args[0])
		if deleteResult == nil {
			fmt.Printf("delete success: cache '%v' ", cache)
			fmt.Println()
			return nil
		}
		return errors.New("delete fail")
	},
}


// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//attach subcommands to rootcommand
	RootCmd.AddCommand(putCmd)
	RootCmd.AddCommand(readCmd)
	RootCmd.AddCommand(updateCmd)
	RootCmd.AddCommand(deleteCmd)
	RootCmd.Execute()
}

