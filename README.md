# filter
Filter provides the ability to quickly and easily filter an arbitrary slice of structs based on certain dynamic constraints **at runtime**. This makes it simple to ask questions like: "Give me all `Thing`s where A="HI" and B=2" where the actual constraints themselves are dynamic and determined at runtime. 

Though there are limited compile time safety checks (due to the need to use `interface{}`), there are extensive runtime checks in an attempt to ensure programs don't crash or do bad things (you'll get `error`s returned to you if you attempt to use this package incorrectly).

If you do not need dynamic constraints at runtime, or do not need a general filtering function, you may wish to implement someting more specific and perhaps better for your use.

Short Example:

```go
package main

import (
	"fmt"
	"github.com/suyashkumar/filter"
)

type Thing struct {
	A string
	B int64
}

func main() {
	// Thing Array
	in := []Thing{
		Thing{A: "HI", B: 1},
		Thing{A: "HI2", B: 2},
	}

	cons, _ := filter.NewConstraints(Thing{}) // Creates a new Constraints for the Thing struct
	cons.Add("A", "HI")                       // Add a constraint to filter on Thing.A = "HI"

	out, _ := filter.Filter(in, cons) // Filter

	// Map out []interface{} into your desired type
	outArr := make([]Thing, len(out))
	for i := 0; i < len(out); i++ {
		if out[i] == nil {
			continue
		}
		outArr[i] = out[i].(Thing)
	}

	fmt.Println(outArr) // [{HI 1} { 0}]
}
```
