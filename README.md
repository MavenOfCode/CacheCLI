# Cache CLI

## Purpose
- Work with Go
- Practice writing tests in Go
- Learn more about structs and interfaces
- Learn to work with Cobra and other Go libraries
- Learn about gRPC and how to connect CLI with simple string key-value cache servers like [KeyValueCache](https://github.com/FavoredFortune/KeyValueCache)
- Grow project over time to add other technologies and achieve other learning goals

## Technologies
- Go
- Markdown (MD)
- Goland (IDE)
- Stretchr/testify API (github.com/stretchr/testify)
- Cobra API (https://github.com/spf13/cobra)
- Flag API (from Go https://golang.org/pkg/flag/)
_more soon_

## User Stories
- As a user I want to put in code arguments like `create animal horse` that work with a Go CLI 'client' to take the commands in a terminal and add the values `animal` and `horse` as a `key:value` pair in this `SimpleKeyValueCache` construct
- As a user I want to type in the command `read animal` to the CLI and have the Go application return `>>horse`
- As a developer I want unit tests for each method (`Create`, `Read`, `Update`, `Delete`) that prove it works with both good and bad input
- As a developer I want to use the Cobra library to build the CLI (per Troy Dai)
* As a developer I want to use the Go Flag library in the CLI (per Troy & Scott)
- As a developer I want unit tests for each command 
- As a developer I want integrated tests for the CLI
- More developer stories details coming soon from Scott

## Instructions to build the CLI
1. Download both [this](https://github.com/FavoredFortune/CacheCLI) repo into the same `go/src` directory within your home user directory

2. In your terminal, go the the `go/src` directory where the project repo now lives
3. Type `go build -o bin/kvc` into the command line of the terminal. This generates the binary executable that will allow you to use the command line interface (CLI)


## Instructions to run / use the CLI
1. To execute any command in this CLI, be sure to be inside it's directory where the file lives (inside the `bin` directory you just created with the `go build -o bin/kvc` command and  always type `./kvc` first
1. After writing `./kvc` you may enter your chosen command from these options followed by the required data strings as noted in the [Expected Command Behaviors ](#Expected Command Behaviors) section below : 
    *  **help**, **-h** (no other information input)
    
    *  **create** (followed by a string that will be your `key` to associate with your next string `value`) 
    
    *  **read** (followed by a string that will be your `key` to return your `value`)

    *  **update** (followed by a string that will be your `key` to associate with your next string `value` that you are updating from the original  `value`)
    
    *  **delete** followed by a string that will be your `key` to remove the `key-value` pair from your data cache)

## Expected command behaviors
- **`create`** puts a `key` string and a `value` string into a the designed simple value cache struct as a key:value pair
 
  **Looks like this when entered in the terminal:** `./kvc create <key> <value>`
  
  **Expected output of successful command:** `create success: cache [full cache data key-value pair sets]`
  
  **Possible errors that may occur at time of input:*
    * empty string provided as key or value
    * key or value not provided
    * key already exists in cache (each key must be unique)
  
- **`read`** takes an input string `key` finds that in the cache (checking to be sure cache exists and key exists) and return it's paired `value` string

   **Looks like this when entered in the terminal:** `./kvc read <key>`
   
   **Expected output of successful command:**
   `>> value for key is: <value>`
   
   **Possible errors that may occur at time of input**:
     
    * empty string provided as key 
    * key provided doesn't exist in cache
    * key provided isn't the key you're looking for
   
- **`update`** allows user to input any existing `key` string and change it's `value` string(after checking that cache and key exists)

   **Looks like this when entered in the terminal:** `./kvc update <key> <value>`
   
   **Expected output of successful command:**
   `update success: cache [full cache data key-value pair sets]`
   
  **Possible errors that may occur at time of input**:
    
   * key doesn't already exist in cache to be updated with new value
   * key is an empty string


- **`delete`** allows user to delete `key-value` string pair by inputting just the `key` from the cache (after checking that the cache and key exist)

   **Looks like this when entered in the terminal:** `./kvc delete <key>`
   
   **Expected output of successful command:**
   `delete success: cache [full cache data key-value pair sets, with target key-value pair removed]`
   
  **Possible errors that may occur at time of input**:
    
   * empty string provided as key 
   * key provided doesn't exist in cache
   * key is an empty string
   


## See [Change logs](Change_Logs.md) for detailed examples of command behaviors throughtout development 

## Future state of CLI
- Include flags for `key` that are `-k` or `-key` and `-v` and `-value` for value
- Have verbose manual option available with typing in command `man`
- Future `Read` command feature could have note if extra args passed in and return message like "I notice extra words in your `read` command, did you mean to `put` or `update` a key value pair instead of `read`?"


## Resources
- Mentors and co-workers: Scott Hornberger and Troy Dai
- gitignore generated on gitignore.io
- Article on testing in Go: https://medium.com/@matryer/5-simple-tips-and-tricks-for-writing-unit-tests-in-golang-619653f90742
- BETTER article on testing in Go with your own written unit tests: https://www.calhoun.io/how-to-test-with-go/
- On building a simple CLI in Go: https://blog.rapid7.com/2016/08/04/build-a-simple-cli-tool-with-golang/
- More on flags in Go: https://gobyexample.com/command-line-flags
- If I want to build a CLI using the Go CLI library: https://tutorialedge.net/golang/building-a-cli-in-go/

- For help on Cobra work around for auto generating command files - after using command with -t for package name, must go get file from gocode folder and put in app folder
https://github.com/spf13/cobra/pull/817

- Tutorial on building a CLI with Cobra: https://ordina-jworks.github.io/development/2018/10/20/make-your-own-cli-with-golang-and-cobra.html

- Make sub commands shared to convert to functions: https://stackoverflow.com/questions/43747075/cobra-commander-how-to-call-a-command-from-another-command

- Cobra Documentation: https://github.com/spf13/cobra

- Go Documentation on flags: https://golang.org/pkg/flag/

- More on flags: https://flaviocopes.com/go-command-line-flags/

- Adjust build creation to help make not pushing up go binary files easier: https://stackoverflow.com/questions/9952061/git-ignore-compiled-google-go


#### [Practice CLI project](https://github.com/FavoredFortune/CobraCLI)
- See [this project](https://github.com/FavoredFortune/CobraCLI) example CLI application with instructions in the [README](https://github.com/FavoredFortune/CobraCLI/blob/master/README.md)
