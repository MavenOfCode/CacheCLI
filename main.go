// Copyright © 2019 Sooz Richman (FavoredFortune - GitHub)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"KVCache/kvcache"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
)

//use constructor from kvchache package for global access to struct and it's private fields per Troy
var cache = kvcache.NewSimpleKVCache()

func main() {

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

	//attach subcommands to rootcommand
	RootCmd.AddCommand(putCmd, readCmd)
	RootCmd.Execute()


}




