# three_way_merge_api_go
 An API that attempts to automatically merge the changes in 2 files against an original file

## How to use

To use this api, simply get the dependencies:
```shell
go mod tidy
```

Run the API

```shell
go run .
```

Then send a POST request with file form data with the following names:
original_file, a_file, b_file