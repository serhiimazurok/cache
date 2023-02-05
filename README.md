# Ninja Cache

Simple in-memory cache for go.

How to use:
```go
package main

import (
	"fmt"
	"github.com/serhiimazurok/cache"
)

func main() {
	c := cache.New()
	
	c.Set("test", "test")
	fmt.Println(c.Get("test"))
	
	c.Delete("test")
	fmt.Println(c.Get("test"))
}
```

