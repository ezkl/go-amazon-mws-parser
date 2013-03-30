# Amazon MWS Products API Parser

This library will quickly parse the XML response body of an Amazon MWS Products
API request and cast it to Go data structures.

# Usage

```go
package main

import (
	"fmt"
	parser "github.com/ezkl/go-amazon-mws-parser"
	"io/ioutil"
	"os"
)

func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		fmt.Println(err)
	}

	doc := parser.Parse(bytes)

	for _, result := range doc.Results {
		fmt.Println("============", result.ASIN, "============")
		for _, offer := range result.Product.Offers {
			fmt.Println(offer)
		}
	}
}
```

## Quickstart Example

Quickstart: `cat ./data/response.xml | go run ./example/mws-example.go`

## TODO

Code review from someone with more Go experience.
