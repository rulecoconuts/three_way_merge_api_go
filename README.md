# three_way_merge_api_go
 An API that attempts to automatically merge the changes in 2 files against an original file. This project is designed to improve the efficiency and correctness of the 3-way merge algorithm. It is a component of the Langresus project.

 ![Screen Recording flow](https://drive.google.com/uc?export=view&id=18QRm8eSqMwGqqMZ5hz6zoeB9Sy1rOANK "Screen Recording flow")

![Demo](https://github.com/rulecoconuts/three_way_merge_api_go/blob/160f27c87b9f8149e3676e1d994db6b0c6198603/merge_web.png "Demo")

<br/>

![Demo Conflict Highlights](https://github.com/rulecoconuts/three_way_merge_api_go/blob/435da6190d2a022941cdcd57e74ab9ba3b2ac3eb/merge_web_conflict_highlight.png "Demo Conflict Highlightts")



## How to use

To use this api, simply get the dependencies:
```shell
go mod tidy
```

### Run the API

```shell
go run .
```

Send a POST request to http://localhost:8000/merge with file form data with the following names:
original_file, a_file, b_file

### Run the experimentation tool web app

If you do not want to deal with the API directly, you can run the simple 3-way merge experimentation/improvement tool web app.

Run the following command: 
```shell
cd frontend
npm run dev
```

Then visit the web app at http://localhost:3000

### Skip the hassle and run all components at once

The following command will install the necessary packages for the API and experimentation tool before running them, assuming that npm and Go have been installed.

```shell
bash run_all.sh
```

You can also just run the entire project as a Docker container by running the following command:

```shell
bash run_docker.sh
```