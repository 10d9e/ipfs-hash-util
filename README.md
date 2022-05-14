# ipfs-hash-util

```go
package main

import (
	"fmt"

	u "github.com/jlogelin/ipfs-hash-util"
)

func main() {
	digest, err := u.Sum([]byte("hello world"))
	if err != nil {
		t.Fatal(err)
	}
	// "bafkqac3imvwgy3zao5xxe3de"
	fmt.Println(u.EncodeToString(digest))
}
```
