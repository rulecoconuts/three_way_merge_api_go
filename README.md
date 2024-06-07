# three_way_merge_api_go
 An API that attempts to automatically merge the changes in 2 files against an original file. This project is designed to improve the efficiency and correctness of the 3-way merge algorithm. It is a component of the Langresus project.

![Demo](https://github.com/rulecoconuts/three_way_merge_api_go/blob/160f27c87b9f8149e3676e1d994db6b0c6198603/merge_web.png "Demo")


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

### Run the experimentation tool web app

If you do not want to deal with the API directly, you can run the simple 3-way merge experimentation/improvement tool web app.

```shell
cd frontend
npm run dev
```

### Skip the hassle and run all components at once

The following command will install the necessary packages for the API and experimentation tool before running them, assuming that npm and Go have been installed.

```shell
bash run_all.sh
```