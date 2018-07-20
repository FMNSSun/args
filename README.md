# args

A package to turn variadic arguments into maps.

## Example


### Unchecked

```go
func f(args... interface{}) {
	m, err := Args(args...)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(m)
}

func main() {
	f("age", 9, "name", "Hans")
}
```

### Checked

```go
package main

import (
	. "github.com/FMNSSun/args"

	"fmt"
	"reflect"
)

var ts TypesSpec = TypesSpec {
	"age" : &TypeSpec {
			Type: reflect.TypeOf(0),
			Optional: true,
	},
	"name" : &TypeSpec {
		Type: reflect.TypeOf(""),
		Optional: false,
	},
}

func f(args... interface{}) {
	m, err := ArgsChecked(ts, args...)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(m)
}

func main() {
	f("age", 9, "name", "Hans")
}
```