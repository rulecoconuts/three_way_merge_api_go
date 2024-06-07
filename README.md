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

![Demo](https://github.com/rulecoconuts/three_way_merge_api_go/blob/160f27c87b9f8149e3676e1d994db6b0c6198603/merge_web.png "Demo")