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

## File Structures

### Controllers

### libraries

- calculator: Simple functions add / sum

### module

- handler: router Handling Modules

### utilities

- shutdown: Graceful Shutdown
