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

## File Structures

### Controllers

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
