# ipfs-hash-util

```
package main

import (
	"fmt"

	u "github.com/jlogelin/ipfs-hash-util"
)

func main() {
	u.Sum(byte[]("hello world"))

	digest, err := Sum([]byte("hello world"))
	if err != nil {
		t.Fatal(err)
	}
	// "bafkqac3imvwgy3zao5xxe3de"
	fmt.Println(u.EncodeToString(digest))
}
```