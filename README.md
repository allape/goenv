# os.Getenv wrapper

## Installation

```shell
go get -u github.com/allape/goenv
```

## Example

```go
package main

import (
	"fmt"
	"github.com/allape/goenv"
	"reflect"
)

type StringAlias string

func (s StringAlias) Print() {
	fmt.Println(s)
}

func main() {
	value, err := goenv.MustGetenv("ENV_VAR", 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(value)                        // 1
	fmt.Println(reflect.TypeOf(value).Kind()) // int // see https://golang.org/pkg/reflect/#Kind for more information

	str, err := goenv.MustGetenv("ENV_VAR", StringAlias("string alias"))
	if err != nil {
		panic(err)
	}
	str.Print() // string alias
}

```
