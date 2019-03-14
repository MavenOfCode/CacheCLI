# Cache CLI

## Don't have Go installed?

Visit [this article](https://medium.com/@pknerd/lets-go-a-very-brif-introduction-to-go-language-e66cb5962900) for background on Go and how to install it on Mac OSX

## Instructions to run the full Client/Server CLI Application
1. Download [this](https://github.com/FavoredFortune/CacheCLI) repo into the same `go/src` directory within your home
 user directory into a new directory with `server` in the title and repeat the process with a second copy of the repo
  into a second new directory with `kvc` in the title
2. From the `server` directory where this repo/project now lives do the following:
   1. Open the `main.go` file and delete the line `cmd.Execute()`
   2. Open a terminal and type `go build -o bin/startserver`. This command builds the server executable.
   
     iii. Next, type Leave this 
   terminal 
    open until you want to end the application. 
    
      **TO END THE APPLICATION**: Quit the server by typing `Ctrl + C` in 
    this terminal. 

3.From the `kvc` directory where the second copy of this repo now lives, do the following:
   
   i. Open the `main.go` file in this project and delete the line `go server.StartServer("8080")`
    
   ii. Next in a second terminal in this copy of the project `go build -o bin/kvc`. This generates the binary 
   executable 
   that will allow you to use the KVC command line interface (CLI)
     
   iii. Follow Instructions to use the CLI below.

## Instructions to run / use the CLI
1. To execute any command in this CLI, be sure to be inside it's directory where the executable lives (inside the `bin` 
directory you just created with the `go build -o bin/kvc` command and  always type `./kvc` first
1. After writing `./kvc` you may enter your chosen command from these options followed by the required data strings as noted in the [Expected Command Behaviors ](#Expected Command Behaviors) section below : 
    *  **help**, **-h** (no other information input)
    
    *  **create** (followed by a string that will be your `key` to associate with your next string `value`) 
    
    *  **read** (followed by a string that will be your `key` to return your `value`)

    *  **update** (followed by a string that will be your `key` to associate with your next string `value` that you are updating from the original  `value`)
    
    *  **delete** followed by a string that will be your `key` to remove the `key-value` pair from your data cache)

## Expected CLI command behaviors
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

## How to run tests

 1. Navigate to the package/directory you want to test in your terminal
 2. To run the tests and see how each passes type `$ go test -v`
 3. To generate a coverage report for the package enter `$go test -cover`
 4. To see the coverage report in html follow these steps:
        
       1. in terminal `$ go test -coverprofile=coverage.out`
 
       2. in terminal `$ go tool cover -html=coverage.out -o coverage.html`
       3. in IDE right click on `coverage.html` file, choose **open in browser** option
 5. After step 4, if you want to see coverage by function in terminal `$ go tool cover -func=coverage.out`

## See [Test Coverage Reports](TestCoverageReports.md) for details on testing for each package

## See [Change logs](Change_Logs.md) for detailed examples of command behaviors throughtout development 

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

- Go testing - go help commands in terminal
- Go testing from the golang blog: https://blog.golang.org/cover
- More on funcs and test within testify/assert - docs - https://godoc.org/github.com/stretchr/testify

- CLI Client/Server API design https://thenewstack.io/make-a-restful-json-api-go/ with https://github.com/corylanou/tns-restful-json-api/tree/master/v9 Project as reference
- Server design articles: https://hackernoon.com/restful-api-design-step-by-step-guide-2f2c9f9fcdbf and https://medium.com/hashmapinc/rest-good-practices-for-api-design-881439796dc9

- Unit testing Hanlder Funcs and Servers in Go - httptest package from Golang documentation: https://golang.org/pkg/net/http/httptest/

- GoLang Blog on JSON handling in GO: https://blog.golang.org/json-and-go

- More on JSON in Go: https://golang.org/pkg/encoding/json/

- More on working with JSON in Go: http://polyglot.ninja/golang-json/

- More on byte slices and strings: https://programming.guide/go/convert-string-to-byte-slice.html
- Blog post on building http client in Go:
http://polyglot.ninja/golang-making-http-requests/

- Building an HTTP Client http://polyglot.ninja/golang-making-http-requests/

- Trouble shooting HTTP requests: https://nanxiao.me/en/fix-unsupported-protocol-scheme-issue-in-golang/

- Having two executables in one project/repo: (https://stackoverflow.com/questions/50904560/how-to-structure-go-application-to-produce-multiple-binaries/50904959)

- More on HTTP package: https://golang.org/pkg/net/http/

- More on multiple packages and binaries: https://ieftimov.com/post/golang-package-multiple-binaries/ with reference 
repo: https://github.com/fteem/fortune_teller

- More on multiple binaries approach: https://stackoverflow.com/questions/37680437/how-do-you-write-a-package-with-multiple-binaries-in-it

- More on build package from Go docs: https://golang.org/pkg/go/build/

- More on all Go commands from Go docs: https://golang.org/cmd/go/


#### [Practice CLI project](https://github.com/FavoredFortune/CobraCLI)
- See [this project](https://github.com/FavoredFortune/CobraCLI) example CLI application with instructions in the [README](https://github.com/FavoredFortune/CobraCLI/blob/master/README.md)
