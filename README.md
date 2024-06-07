# three_way_merge_api_go
 An API that attempts to automatically merge the changes in 2 files against an original file

## How to use

To use this api, simply get the dependencies:
```shell
go mod tidy
```

### Run the API

```shell
go run .
```

Then send a POST request with file form data with the following names:
original_file, a_file, b_file

### Run the web experimentation tool

If you do not want to deal with the API directly. You can run the simple 3-way merge experimentation/improvement tool.

```shell
cd frontend
npm run dev
```