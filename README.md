# Remote [![Go Reference](https://pkg.go.dev/badge/github.com/201R/remote.svg)](https://pkg.go.dev/github.com/201R/remote) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

 
Remote is a Golang http request package to make request like POST GET PUT ... on a domain or an API 

## Installation

To install remote package, you need to install Go and set your Go workspace first.

1. The first need [Go](https://golang.org/) installed (**version 1.12+ is required**), then you can use the below Go command to install Gin.

```sh
$ go get -u github.com/201RichK/remote
```

2. Import it in your code:

```go
import "github.com/201R/remote"
```
3. Quick start
 
```sh
# assume the following codes in example.go file
$ cat example.go
```

```go
package main

import (
    "github.com/201R/remote"
    "fmt"
)

func main() {
    //Create new config 
    client := remote.NewRemote(remote.Config{})

    res, err := client.GET(remote.Options{
        URL: "https://google.com",
    })

    if err != nil {
        fmt.Println(err)
    }

    if err := json.NewDecoder(os.Stdout).Decode(res.Body); err != nil {
        fmt.Println(err)
    }
    
}
```