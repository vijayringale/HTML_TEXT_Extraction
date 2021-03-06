package main

import (
	"fmt"

	"/xtract"
)

func main() {
	v, err := xtract.Page("https://godoc.org/golang.org/x/net/html")

	if err != nil {
		panic(err)
	}
	fmt.Println("Page without limit: ",v)

	v1 ,err := xtract.PageLim("https://godoc.org/golang.org/x/net/html",20)

	if err != nil {
		panic(err)
	}

	fmt.Println("Page With Limit : ", v1)
	htmlVal := "<p>This is a paragraph <a>This is a link inside a paragraph</a></p>"
	v2 := xtract.Value(htmlVal)
	fmt.Println("Value without limit:", v2)

	v3 := xtract.ValueLim(htmlVal, 4)
	fmt.Println("Value with limit:", v3)

}
