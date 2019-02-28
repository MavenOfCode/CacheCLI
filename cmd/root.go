package cmd

import (
	"KVCache/kvcache"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

//use constructor from kvchache package for global access to struct and it's private fields per Troy
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
	Use:"read",
	Short: "read given key and return value",
	Args: cobra.MinimumNArgs(1),
	Long: "read value string out to command line from key-value cache given key string input from command line",
	RunE: func(cmd *cobra.Command, args []string) error  {
		fmt.Println(args, len(args))
		if cache == nil {
			return errors.New("cache empty - read failed: ")
		}
		readResult, err := cache.Read(args[0])
		if err !=nil {
			return err
		}
		fmt.Println(">>", readResult)
		return nil
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
	RootCmd.AddCommand(putCmd, readCmd)
	RootCmd.Execute()
}

