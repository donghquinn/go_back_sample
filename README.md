# Go Lang Backend

- Author: donghquinn

## Dependencies

- gin: Web Framework

``` shell
go get -u github.com/gin-gonic/gin
```

- Gin Usage

    - Post Body Data Parsing
        - Bind with Body Data Structures.
        - Structure can be defined with expression `json:"body1" xml:"body1" binding:"required"`
        - can be validate with ShouldBind() method

- godotenv: load .env file

``` shell
go get -u github.com/joho/godotenv
```

- prisma: Data ORM Client

``` shell
go get -u github.com/steebchen/prisma-client-go
```

## Prisma Initiate

- If you get Prisma Module, Please make schema, File name is schema.prisma. Check prisma/schema.prisma as reference on this repository.
- Run Database you want, and write its dataurl into .env file.
    - format: `[databaseType]://[userName]:[userPassword]@[hostAddr]:[hostPort]/[databaseName]?schema=[public / private]`
- Then Run command below to initiate prisma client. This command line will create the package name "db"

``` shell
go run github.com/steebchen/prisma-client-go generate
```

## File Structures

### Controllers

### data

- Prisma Client and Query Modules

### DTO

- Data Transfer Object
    - Formulized Response with throw keys; 
        - resCode: Response Code - 200 / 40* /500
        - dataRes: Actual Result of Request
        - errMsg: There would be Error message, if there were any error happend.

### libraries

- calculator: Simple functions add / sum

### module

- handler: router Handling Modules

### utilities

- shutdown: Graceful Shutdown

### types

- calc_request: Calculation Request Type Definition
