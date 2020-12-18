# ProjectShare API
ProjectShare is a prototype we built to explore the building of an application which allows uploading and downloading files for
projects. Also we experimented with the usage of go as backend server for react applications. 

## Prequisites
If you would like to use S3, you need credentials for it. For example you can use a `~/.aws/credentials` which contains them.  
See the S3 documentation for details.

You can also use an in-memory implementation to try this application if you don't have access to an S3 server.

## Run

### With S3
```
go run ./cmd/projectshare-server
```

### With in-memory
```
go run ./cmd/projectshare-server-memory
```
Note that it stores everything in memory and therefore you need enough RAM. It is only meant for testing and demo purposes.
After stopping the application all data gets lost.

## Project structure
### main
Main is located in `cmd/projectshare-server`.
This is done by a popular go-convention to allow to easily add other mains for different execution paths. So if another main is needed, you can just add a new package in the cmd folder. This is done here by the projectshare-server-memory main.

`projectshare-server` just instantiates and starts the server. Later it may be also responsible for loading a config and parsing command line arguments.

Also it instantiates the handlers and repositories which get passed to the handlers as they need them.

### api
The package api contains everything to serve the api endpoints. 
For this it sets up a webserver using Chi as a lightweight router.

All DTO's are in the sub-package `dto`. They should only be used by the api and by the handlers as return type.  
They should __NEVER__ be used by any repository. Instead these should only use the models defined inside of the `handler.model` package to provide a proper separation of the repository implementation and the api itself.

### handler
The handlers contain the business logic. 
They accept parameters and data from the api, do something with them and then may return a result as DTO.

### repository
The repository contains all logic which has to do with external interfaces, like database access, AWS S3 access and similar.

They have to implement the repository interfaces located in `handler/repository.go`.
These interfaces only return models defined inside of the `handler.model` package. 
**Important:** The models should be built in a way which is __NOT__ coupled with any specific repository implementation. (e.g. for example types defined in the aws or mongo libraries.)

This is because of easy replacement and mocking of different implementations.
You just have to satisfy the respective handler interface with your implementation and instantiate it in the server (or inside of a test, etc.).

Currently two implementations exist. One AWS S3 and one simple in-memory implementation for testing purposes.