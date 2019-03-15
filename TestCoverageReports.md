# Testing coverage reports - run via `go test -cover`


## KVCache: 91.7% of statements; 0.016s (3/4/19)

## KVCache: 100% of statements; 0.14s (3/4/19)

--__**NOTE**__ Ran `go tool cover -html=c.out -o coverage.html` to find where coverage is lacking for this package. See [readme](README.md)  for full details on how to run your own.

### `go tool cover -func=coverage.out`

  ```
  CacheCLI/kvcache/kvcache.go:21:	NewSimpleKVCache	100.0%
     CacheCLI/kvcache/kvcache.go:26:	Create			100.0%
     CacheCLI/kvcache/kvcache.go:46:	Read			100.0%
     CacheCLI/kvcache/kvcache.go:54:	Update			100.0%
     CacheCLI/kvcache/kvcache.go:63:	Delete			100.0%
     total:				(statements)		100.0%
   ```
 

### `go test -v`(verbose)
```=== RUN   TestSimpleKeyValueCache
   === RUN   TestSimpleKeyValueCache/it_creates_new_cache
   --- PASS: TestSimpleKeyValueCache (0.00s)
       --- PASS: TestSimpleKeyValueCache/it_creates_new_cache (0.00s)
   === RUN   TestCreate
   === RUN   TestCreate/it_creates
   === RUN   TestCreate/create_instantiates_cache_when_cache_starts_as_nil
   === RUN   TestCreate/_create_returns_error_when_empty_string_given_as_parameter
   === RUN   TestCreate/_create_returns_error_when_key_already_exists
   bobby
   create failed: key 'name' isn't unique: 
   bobby
   --- PASS: TestCreate (0.00s)
       --- PASS: TestCreate/it_creates (0.00s)
       --- PASS: TestCreate/create_instantiates_cache_when_cache_starts_as_nil (0.00s)
       --- PASS: TestCreate/_create_returns_error_when_empty_string_given_as_parameter (0.00s)
       --- PASS: TestCreate/_create_returns_error_when_key_already_exists (0.00s)
   === RUN   TestRead
   === RUN   TestRead/it_reads
   === RUN   TestRead/read_successfully_when_given_different_keys
   === RUN   TestRead/read_returns_error_when_given_empty_string
   === RUN   TestRead/read_returns_error_when_given_invalid_key
   --- PASS: TestRead (0.00s)
       --- PASS: TestRead/it_reads (0.00s)
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
   ok  	CacheCLI/kvcache	0.016s
```

## CMD: only 56.2% of statements; 0.013s (3/4/19)

## CMD: 84.5% coverage of statements; 0.014s (3/5/19)
## 95.2% according to HTML coverage report
   
--__**NOTE**__ Ran `go tool cover -html=c.out -o coverage.html` to find where coverage is lacking for this package. 
See [readme](README.md) for full details on how to run your own.

### `go tool cover -func=coverage.out` 
```
///MOCK METHODS FOR TESTING
CacheCLI/cmd/cmd-runner.go:21:	Create			100.0%
CacheCLI/cmd/cmd-runner.go:27:	Read			83.3%
CacheCLI/cmd/cmd-runner.go:38:	Update			83.3%
CacheCLI/cmd/cmd-runner.go:49:	Delete			83.3%
CacheCLI/cmd/cmd-runner.go:61:	NewMockSimpleKVCache	100.0%
///ALL COMMANDS
CacheCLI/cmd/cmd-runner.go:66:	CreateCmd		100.0%
CacheCLI/cmd/cmd-runner.go:83:	ReadCmd			100.0%
CacheCLI/cmd/cmd-runner.go:102:	UpdateCmd		100.0%
CacheCLI/cmd/cmd-runner.go:121:	DeleteCmd		100.0%
CacheCLI/cmd/root.go:52:	Execute			0.0%
total:				(statements)		84.5%
```

### `go test -v` (verbose testing)

  ```
  srichm :~/gocode/src/CacheCLI/cmd :[tues-cli-test !] go test -v
  === RUN   TestMockSimpleKeyValueCache
  === RUN   TestMockSimpleKeyValueCache/it_creates_a_mock_cache
  --- PASS: TestMockSimpleKeyValueCache (0.00s)
      --- PASS: TestMockSimpleKeyValueCache/it_creates_a_mock_cache (0.00s)
  === RUN   TestCommandRunner_CreateCmd
  === RUN   TestCommandRunner_CreateCmd/it_creates
  create success:  cache '&{false testValueString}' 
  === RUN   TestCommandRunner_CreateCmd/create_command_returns_error_when_cache_is_nil
  === RUN   TestCommandRunner_CreateCmd/create_command_returns_error_when_one_of_2_args_missing
  === RUN   TestCommandRunner_CreateCmd/create_command_returns_error_when_both_args_missing
  --- PASS: TestCommandRunner_CreateCmd (0.00s)
      --- PASS: TestCommandRunner_CreateCmd/it_creates (0.00s)
      --- PASS: TestCommandRunner_CreateCmd/create_command_returns_error_when_cache_is_nil (0.00s)
      --- PASS: TestCommandRunner_CreateCmd/create_command_returns_error_when_one_of_2_args_missing (0.00s)
      --- PASS: TestCommandRunner_CreateCmd/create_command_returns_error_when_both_args_missing (0.00s)
  === RUN   TestCommandRunner_ReadCmd
  === RUN   TestCommandRunner_ReadCmd/it_reads
  create success:  cache '&{true ReturnString}' 
  &{0xc0000ae5c0}
  >> value for key is:  ReturnString
  === RUN   TestCommandRunner_ReadCmd/read_command_returns_error_when_key_is_invalid
  &{0xc0000ae5c0}
  === RUN   TestCommandRunner_ReadCmd/read_command_returns_error_when_args_are_insufficient
  === RUN   TestCommandRunner_ReadCmd/read_command_returns_error_when_cache_is_nil_
  &{<nil>}
  {<nil>}
  --- PASS: TestCommandRunner_ReadCmd (0.00s)
      --- PASS: TestCommandRunner_ReadCmd/it_reads (0.00s)
      --- PASS: TestCommandRunner_ReadCmd/read_command_returns_error_when_key_is_invalid (0.00s)
      --- PASS: TestCommandRunner_ReadCmd/read_command_returns_error_when_args_are_insufficient (0.00s)
      --- PASS: TestCommandRunner_ReadCmd/read_command_returns_error_when_cache_is_nil_ (0.00s)
  === RUN   TestCommandRunner_UpdateCmd
  === RUN   TestCommandRunner_UpdateCmd/it_updates
  create success:  cache '&{true ReturnString}' 
  update success:  cache '&{true ReturnString}' 
  <nil>
  === RUN   TestCommandRunner_UpdateCmd/update_returns_error_when_invalid_key_provided
  create success:  cache '&{false ReturnString}' 
  update error
  === RUN   TestCommandRunner_UpdateCmd/update_returns_error_when_cache_is_nil
  === RUN   TestCommandRunner_UpdateCmd/update_returns_error_when_min_args_aren't_provided
  === RUN   TestCommandRunner_UpdateCmd/update_returns_error_when_key_is_empty_string
  update error
  update error
  --- PASS: TestCommandRunner_UpdateCmd (0.00s)
      --- PASS: TestCommandRunner_UpdateCmd/it_updates (0.00s)
      --- PASS: TestCommandRunner_UpdateCmd/update_returns_error_when_invalid_key_provided (0.00s)
      --- PASS: TestCommandRunner_UpdateCmd/update_returns_error_when_cache_is_nil (0.00s)
      --- PASS: TestCommandRunner_UpdateCmd/update_returns_error_when_min_args_aren't_provided (0.00s)
      --- PASS: TestCommandRunner_UpdateCmd/update_returns_error_when_key_is_empty_string (0.00s)
  === RUN   TestCommandRunner_DeleteCmd
  === RUN   TestCommandRunner_DeleteCmd/it_deletes
  delete success: cache '&{true }' 
  === RUN   TestCommandRunner_DeleteCmd/delete_command_returns_error_when_key_is_invalid
  delete error
  === RUN   TestCommandRunner_DeleteCmd/delete_command_returns_error_when_args_are_insufficient
  === RUN   TestCommandRunner_DeleteCmd/delete_command_returns_error_when_key_is_empty_string
  delete error
  === RUN   TestCommandRunner_DeleteCmd/delete_command_returns_error_when_cache_is_nil_
  --- PASS: TestCommandRunner_DeleteCmd (0.00s)
      --- PASS: TestCommandRunner_DeleteCmd/it_deletes (0.00s)
      --- PASS: TestCommandRunner_DeleteCmd/delete_command_returns_error_when_key_is_invalid (0.00s)
      --- PASS: TestCommandRunner_DeleteCmd/delete_command_returns_error_when_args_are_insufficient (0.00s)
      --- PASS: TestCommandRunner_DeleteCmd/delete_command_returns_error_when_key_is_empty_string (0.00s)
      --- PASS: TestCommandRunner_DeleteCmd/delete_command_returns_error_when_cache_is_nil_ (0.00s)
  PASS
  ok  	CacheCLI/cmd	0.013s
```

## Server: only 45.7% of statements; 0.017s (3/11/19)

## Server: 61.9% of statements; 0.015s (3/13/19)*
 *some if/err statements untestable in this framework as they are 
  package commands

--**NOTE** Ran go tool cover `-html=c.out -o coverage.html` to find where coverage is lacking for this package. See 
[readme](README.md) for full details on how to run your own.



### `go tool cover -func=coverage.out`


    ```
    //NOT TESTING PACKAGE PROVIDED METHODS
    CacheCLI/server/server.go:33:	StartServer	0.0%
    
    //HANDLER HELPER FUNCTION
    CacheCLI/server/server.go:76:	HandleData	48.6%
    
    //SERVER HANDLER FUNCTIONS
    CacheCLI/server/server.go:126:	Put		77.8%
    CacheCLI/server/server.go:153:	Get		63.0%
    CacheCLI/server/server.go:193:	Post		78.9%
    CacheCLI/server/server.go:222:	Delete		78.9%
    total:				(statements)	61.9%
    ```

### `go test -v` (verbose testing)

```=== RUN   TestNewServerTestKeyValueCache
   === RUN   TestNewServerTestKeyValueCache/it_creates_a_new_mock_server_cache
   --- PASS: TestNewServerTestKeyValueCache (0.00s)
       --- PASS: TestNewServerTestKeyValueCache/it_creates_a_new_mock_server_cache (0.00s)
   === RUN   TestServer_HandleData
   === RUN   TestServer_HandleData/handle_data_returns_data
   === RUN   TestServer_HandleData/handle_data_returns_error_if_body_is_nil
   === RUN   TestServer_HandleData/handle_data_returns_error_if_JSON_is_MALFORMED
   --- PASS: TestServer_HandleData (0.00s)
       --- PASS: TestServer_HandleData/handle_data_returns_data (0.00s)
       --- PASS: TestServer_HandleData/handle_data_returns_error_if_body_is_nil (0.00s)
       --- PASS: TestServer_HandleData/handle_data_returns_error_if_JSON_is_MALFORMED (0.00s)
   === RUN   TestServer_Put
   === RUN   TestServer_Put/put_works
   === RUN   TestServer_Put/put_returns_error_when_content_is_empty_-_like_malformed_JSON_error
   === RUN   TestServer_Put/put_returns_error_when_json_malformed
   === RUN   TestServer_Put/put_returns_error_when_key_is_empty
   === RUN   TestServer_Put/put_returns_error_when_value_is_empty
   --- PASS: TestServer_Put (0.00s)
       --- PASS: TestServer_Put/put_works (0.00s)
       --- PASS: TestServer_Put/put_returns_error_when_content_is_empty_-_like_malformed_JSON_error (0.00s)
       --- PASS: TestServer_Put/put_returns_error_when_json_malformed (0.00s)
       --- PASS: TestServer_Put/put_returns_error_when_key_is_empty (0.00s)
       --- PASS: TestServer_Put/put_returns_error_when_value_is_empty (0.00s)
   === RUN   TestServer_Get
   === RUN   TestServer_Get/get_works_and_returns_correct_status
   === RUN   TestServer_Get/get_works_and_returns_correct_value
   === RUN   TestServer_Get/get_returns_error_when_key_doesn't_exist_in_cache
   === RUN   TestServer_Get/get_returns_error_when_JSON_MALFORMED
   --- PASS: TestServer_Get (0.00s)
       --- PASS: TestServer_Get/get_works_and_returns_correct_status (0.00s)
       --- PASS: TestServer_Get/get_works_and_returns_correct_value (0.00s)
       --- PASS: TestServer_Get/get_returns_error_when_key_doesn't_exist_in_cache (0.00s)
       --- PASS: TestServer_Get/get_returns_error_when_JSON_MALFORMED (0.00s)
   === RUN   TestServer_Post
   === RUN   TestServer_Post/post_works
   === RUN   TestServer_Post/post_returns_error_when_key_doesn't_exist_in_cache
   === RUN   TestServer_Post/post_returns_error_when_JSON_MALFORMED
   === RUN   TestServer_Post/post_returns_error_when_content_is_empty_-_like_malformed_JSON_error
   === RUN   TestServer_Post/post_returns_error_when_key_is_empty
   === RUN   TestServer_Post/post_returns_error_when_value_is_empty
   --- PASS: TestServer_Post (0.00s)
       --- PASS: TestServer_Post/post_works (0.00s)
       --- PASS: TestServer_Post/post_returns_error_when_key_doesn't_exist_in_cache (0.00s)
       --- PASS: TestServer_Post/post_returns_error_when_JSON_MALFORMED (0.00s)
       --- PASS: TestServer_Post/post_returns_error_when_content_is_empty_-_like_malformed_JSON_error (0.00s)
       --- PASS: TestServer_Post/post_returns_error_when_key_is_empty (0.00s)
       --- PASS: TestServer_Post/post_returns_error_when_value_is_empty (0.00s)
   === RUN   TestServer_Delete
   === RUN   TestServer_Delete/delete_works
   === RUN   TestServer_Delete/delete_returns_error_when_JSON_MALFORMED
   === RUN   TestServer_Delete/delete_returns_error_when_key_is_empty
   === RUN   TestServer_Delete/delete_returns_error_when_key_is_not_in_cache
   --- PASS: TestServer_Delete (0.00s)
       --- PASS: TestServer_Delete/delete_works (0.00s)
       --- PASS: TestServer_Delete/delete_returns_error_when_JSON_MALFORMED (0.00s)
       --- PASS: TestServer_Delete/delete_returns_error_when_key_is_empty (0.00s)
       --- PASS: TestServer_Delete/delete_returns_error_when_key_is_not_in_cache (0.00s)
   PASS
   ok  	CacheCLI/server	0.014s
   ```