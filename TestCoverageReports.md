# Testing coverage reports - run via `go test -cover`


## KVCache: 91.7% of statements; 0.016s

```=== RUN   TestSimpleKeyValueCache
   === RUN   TestSimpleKeyValueCache/it_creates_new_cache
   --- PASS: TestSimpleKeyValueCache (0.00s)
       --- PASS: TestSimpleKeyValueCache/it_creates_new_cache (0.00s)
   === RUN   TestCreate
   === RUN   TestCreate/creates_and_reads_successfully
   === RUN   TestCreate/it_creates_successfully
   === RUN   TestCreate/_create_returns_error_when_empty_string_given_as_parameter
   === RUN   TestCreate/_create_returns_error_when_key_already_exists
   bobby
   create failed: key 'name' isn't unique: 
   bobby
   --- PASS: TestCreate (0.00s)
       --- PASS: TestCreate/creates_and_reads_successfully (0.00s)
       --- PASS: TestCreate/it_creates_successfully (0.00s)
       --- PASS: TestCreate/_create_returns_error_when_empty_string_given_as_parameter (0.00s)
       --- PASS: TestCreate/_create_returns_error_when_key_already_exists (0.00s)
   === RUN   TestRead
   === RUN   TestRead/reads_successfully
   === RUN   TestRead/read_successfully_when_given_different_keys
   === RUN   TestRead/read_returns_error_when_given_empty_string
   === RUN   TestRead/read_returns_error_when_given_invalid_key
   --- PASS: TestRead (0.00s)
       --- PASS: TestRead/reads_successfully (0.00s)
       --- PASS: TestRead/read_successfully_when_given_different_keys (0.00s)
       --- PASS: TestRead/read_returns_error_when_given_empty_string (0.00s)
       --- PASS: TestRead/read_returns_error_when_given_invalid_key (0.00s)
   === RUN   TestUpdate
   === RUN   TestUpdate/it_can_update
   === RUN   TestUpdate/update_returns_error_when_key_not_in_cache
   === RUN   TestUpdate/update_returns_error_when_key_is_empty_string
   --- PASS: TestUpdate (0.00s)
       --- PASS: TestUpdate/it_can_update (0.00s)
       --- PASS: TestUpdate/update_returns_error_when_key_not_in_cache (0.00s)
       --- PASS: TestUpdate/update_returns_error_when_key_is_empty_string (0.00s)
   === RUN   TestDelete
   === RUN   TestDelete/it_deletes
   === RUN   TestDelete/delete_returns_error_when_key_doesn't_exist
   === RUN   TestDelete/delete_returns_error_when_given_empty_key_string
   --- PASS: TestDelete (0.00s)
       --- PASS: TestDelete/it_deletes (0.00s)
       --- PASS: TestDelete/delete_returns_error_when_key_doesn't_exist (0.00s)
       --- PASS: TestDelete/delete_returns_error_when_given_empty_key_string (0.00s)
   PASS
   ok  	CacheCLI/kvcache	0.014s
```

## CMD:   currently only 56.2% of statements; 0.013s

   ```
   === RUN   TestMockKeyValueCache
   === RUN   TestMockKeyValueCache/it_creates_a_mock_cache
   --- PASS: TestMockKeyValueCache (0.00s)
       --- PASS: TestMockKeyValueCache/it_creates_a_mock_cache (0.00s)
   === RUN   TestCommandRunner_CreateCmd
   === RUN   TestCommandRunner_CreateCmd/it_creates
   create success:  cache '&{map[name:harley animal:horse kitten:Bene testString:testValueString]}' 
   === RUN   TestCommandRunner_CreateCmd/create_command_returns_error_when_cache_is_nil
   create success:  cache '&{map[name:harley animal:horse kitten:Bene testString:testValueString keyTest:testValueString]}' 
   === RUN   TestCommandRunner_CreateCmd/create_command_returns_error_when_one_of_2_args_missing
   === RUN   TestCommandRunner_CreateCmd/create_command_returns_error_when_both_args_missing
   --- PASS: TestCommandRunner_CreateCmd (0.00s)
       --- PASS: TestCommandRunner_CreateCmd/it_creates (0.00s)
       --- PASS: TestCommandRunner_CreateCmd/create_command_returns_error_when_cache_is_nil (0.00s)
       --- PASS: TestCommandRunner_CreateCmd/create_command_returns_error_when_one_of_2_args_missing (0.00s)
       --- PASS: TestCommandRunner_CreateCmd/create_command_returns_error_when_both_args_missing (0.00s)
   === RUN   TestCommandRunner_DeleteCmd
   === RUN   TestCommandRunner_DeleteCmd/it_deletes
   === RUN   TestCommandRunner_DeleteCmd/delete_command_returns_error_when_key_is_invalid
   read failed: key 'false' invalid or cache empty
   === RUN   TestCommandRunner_DeleteCmd/delete_command_returns_error_when_key_is_empty_string
   read failed: key '' invalid or cache empty
   === RUN   TestCommandRunner_DeleteCmd/delete_command_returns_error_when_cache_is_nil_
   read failed: key 'false' invalid or cache empty
   --- PASS: TestCommandRunner_DeleteCmd (0.00s)
       --- PASS: TestCommandRunner_DeleteCmd/it_deletes (0.00s)
       --- PASS: TestCommandRunner_DeleteCmd/delete_command_returns_error_when_key_is_invalid (0.00s)
       --- PASS: TestCommandRunner_DeleteCmd/delete_command_returns_error_when_key_is_empty_string (0.00s)
       --- PASS: TestCommandRunner_DeleteCmd/delete_command_returns_error_when_cache_is_nil_ (0.00s)
   === RUN   TestCommandRunner_ReadCmd
   === RUN   TestCommandRunner_ReadCmd/it_reads
   === RUN   TestCommandRunner_ReadCmd/read_command_returns_error_when_key_is_invalid
   read failed: key 'false' invalid or cache empty
   === RUN   TestCommandRunner_ReadCmd/read_command_returns_error_when_args_are_insufficient
   read failed: at least one argument required
   === RUN   TestCommandRunner_ReadCmd/read_command_returns_error_when_cache_is_nil_
   read failed: key 'false' invalid or cache empty
   --- PASS: TestCommandRunner_ReadCmd (0.00s)
       --- PASS: TestCommandRunner_ReadCmd/it_reads (0.00s)
       --- PASS: TestCommandRunner_ReadCmd/read_command_returns_error_when_key_is_invalid (0.00s)
       --- PASS: TestCommandRunner_ReadCmd/read_command_returns_error_when_args_are_insufficient (0.00s)
       --- PASS: TestCommandRunner_ReadCmd/read_command_returns_error_when_cache_is_nil_ (0.00s)
   === RUN   TestCommandRunner_UpdateCmd
   === RUN   TestCommandRunner_UpdateCmd/it_updates
   create success:  cache '&{map[testString:testValueString keyTest:testValueString ReturnString:hi name:harley animal:horse kitten:Bene]}' 
   update success:  cache '&{map[testString:testValueString keyTest:testValueString ReturnString:bye name:harley animal:horse kitten:Bene]}' 
   === RUN   TestCommandRunner_UpdateCmd/update_returns_error_when_invalid_key_provided
   update failed: key 'key' not in cache
   === RUN   TestCommandRunner_UpdateCmd/update_returns_error_when_cache_is_nil
   update failed: key 'key' not in cache
   === RUN   TestCommandRunner_UpdateCmd/update_returns_error_when_min_args_aren't_provided
   === RUN   TestCommandRunner_UpdateCmd/update_returns_error_when_key_is_empty_string
   update failed: key '' not in cache
   --- PASS: TestCommandRunner_UpdateCmd (0.00s)
       --- PASS: TestCommandRunner_UpdateCmd/it_updates (0.00s)
       --- PASS: TestCommandRunner_UpdateCmd/update_returns_error_when_invalid_key_provided (0.00s)
       --- PASS: TestCommandRunner_UpdateCmd/update_returns_error_when_cache_is_nil (0.00s)
       --- PASS: TestCommandRunner_UpdateCmd/update_returns_error_when_min_args_aren't_provided (0.00s)
       --- PASS: TestCommandRunner_UpdateCmd/update_returns_error_when_key_is_empty_string (0.00s)
   PASS
   ok  	CacheCLI/cmd	0.015s
   ```
   
--__**NOTE**__ Ran `go tool cover -html=c.out -o coverage.html` to find where coverage is lacking for this package. See [file](cmd/coverage.html)
Part is the cmd-runner mock constructor function and methods, but there are more tests to be done to improve coverage.

### Coverage by Command
 - 42.9% cmd-runner.go
 - 91.7% create-cmd.go
 - 75% read-cmd.go
 - 0% root.go (for Execute command)
 - 92% update-cmd.go
 - 0% delete-cmd.go
