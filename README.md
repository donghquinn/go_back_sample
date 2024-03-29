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

- validator: Request Validator. You don't need to install package manually.
    - How to make custom validator: [Korean Reference](https://velog.io/@moonyoung/Golang-%ED%95%84%EB%93%9C-%EA%B2%80%EC%82%AC-validate-%EB%9D%BC%EC%9D%B4%EB%B8%8C%EB%9F%AC%EB%A6%AC-%EC%95%8C%EC%95%84%EB%B3%B4%EA%B8%B0)
    - Unlike node.js, you can validate requests by tags.
    - And of course, you can make custom tags as well: [Korean Reference](https://velog.io/@beardfriend/gogin-%EC%9A%94%EC%B2%AD-Validation%EC%97%90-%EB%8C%80%ED%95%B4-%EC%95%8C%EC%95%84%EB%B3%B4%EC%9E%90)

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

### validators
