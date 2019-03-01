#Change Logs
Provide logs of terminal output/tests with date & time stamps to note progress/changes/concerns

- 2/28/19 4:29 PM

    * _**NOTE:**_ All commands have pre-populated cache data as needed, due to current lack of connection to cache server. **THIS WILL CHANGE**
    
    - [x] All commands show in help menu
    - [x] Put command message shows success
    - [ ] Read command still fails
    - [x] Update command works
    - [x] Delete command works
    
    * _**COMMENT**_ currently using terminal as intermediate test harness. Unit/Integration tests next after commands all work in terminal.
    
    ```
    richm :~/gocode/src/CacheCLI :[thurs-new-cli !] go build -o cli
       srichm :~/gocode/src/CacheCLI :[thurs-new-cli !] ./cli put name ducati
       put success:  cache '&{map[name:ducati]}' 
       srichm :~/gocode/src/CacheCLI :[thurs-new-cli !] ./cli update kitten Benny
       update success:  cache '&{map[name:harley animal:horse kitten:Benny]}' 
       srichm :~/gocode/src/CacheCLI :[thurs-new-cli !] ./cli read kitten
       [kitten] 1
       
       read failed: key 'kitten' invalid or cache empty
       Error: read failed: key 'kitten' invalid or cache empty
       Usage:
         cli read [flags]
       
       Flags:
         -h, --help   help for read
       
       srichm :~/gocode/src/CacheCLI :[thurs-new-cli !] ./cli delete animal
       delete success: cache '&{map[name:harley kitten:Bene]}' 
       srichm :~/gocode/src/CacheCLI :[thurs-new-cli !] ./cli help
       Usage:
         cli [command]
       
       Available Commands:
         delete      delete key-value pair
         help        Help about any command
         put         put key-value pair
         read        read given key and return value
         update      update key-value pair
       
       Flags:
         -h, --help   help for cli
       
       Use "cli [command] --help" for more information about a command.
   ```
- 2/28/19 4:47 PM

    * _**NOTE:**_ All commands have pre-populated cache data as needed, due to current lack of connection to cache server. **THIS WILL CHANGE**
    
    - [x] All commands show in help menu
    - [x] Put command message shows success
    - [x] Read command still fails
    - [x] Update command works
    - [x] Delete command works
    
    * _**COMMENT**_ currently using terminal as intermediate test harness. Unit/Integration tests next after commands all work in terminal.
    
    * PROOF OF WORKING COMMANDS BELOW
    
    ```srichm :~/gocode/src/CacheCLI :[thurs-new-cli !] go build -o cli
       srichm :~/gocode/src/CacheCLI :[thurs-new-cli !] ./cli help
       Usage:
         cli [command]
       
       Available Commands:
         delete      delete key-value pair
         help        Help about any command
         put         put key-value pair
         read        read given key and return value
         update      update key-value pair
       
       Flags:
         -h, --help   help for cli
       
       Use "cli [command] --help" for more information about a command.
       srichm :~/gocode/src/CacheCLI :[thurs-new-cli !] ./cli put kitten harley
       put success:  cache '&{map[kitten:harley]}' 
       srichm :~/gocode/src/CacheCLI :[thurs-new-cli !] ./cli read kitten
       [kitten] 1
       <nil>
       >> Bene
       srichm :~/gocode/src/CacheCLI :[thurs-new-cli !] ./cli update kitten ducati
       update success:  cache '&{map[animal:horse kitten:ducati name:harley]}' 
       srichm :~/gocode/src/CacheCLI :[thurs-new-cli !] ./cli delete horse
       delete failed: key 'horse' not in cache
       Error: delete fail
       Usage:
         cli delete [flags]
       
       Flags:
         -h, --help   help for delete
       
       srichm :~/gocode/src/CacheCLI :[thurs-new-cli !] ./cli delete animal
       delete success: cache '&{map[name:harley kitten:Bene]}'```