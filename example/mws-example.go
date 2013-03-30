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
