# LOWKEYLOBOS Demo Project

## About this project: 
The project is a take home assessment to demonstrate the API design and consumption in GoLang.

## Running the project:
1. clone the project
2. execute `go build` in root of project
3. execute `go run main.go` in root of project
4. the API is served in `http://localhost:10000/`

## Available endpoints:
The API uses Basic Authentication with username `lowkeylobos2022` and password `bijaya-sharma`. 

1. `GET /`: home 
2. `GET /metadata`: get list of metadata
3. `GET /metadata/{resourceId}`: get designated metadata
4. `POST /metadata`: create new metadata
5. `DELETE /metadata/{resourceId}`: delete a metadata 

# NOTE: 
For demo purpose I have used a dummy data to show the capabality of API building. Also I have added a single call to demonstrate the internal service call. 



