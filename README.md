# ipfs-hash-util

```go
package main

import (
	"fmt"
	"io/ioutil"

	"github.com/jlogelin/ipfs-hash-util"
)

func main() {
	// get the multihash sum of ipfs content
	digest, err := Sum([]byte("hello world"))
	if err != nil {
		panic(err)
	}
	// output should be "bafkqac3imvwgy3zao5xxe3de"
	fmt.Println(EncodeToString(digest))

	// try it with a large file
	contents, err := ioutil.ReadFile("/path/to/bigfile")
	if err != nil {
		panic(err)
	}

	digest2, err := Sum(contents)
	if err != nil {
		panic(err)
	}
	// output should look something like "bafybeig4bhrwgp52h3zojj5aw7p7avuwcvp5jmqpzjb2dard3rtjq6joua"
	fmt.Println(EncodeToString(digest2))
}
```
