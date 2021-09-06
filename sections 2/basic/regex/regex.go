package main

import (
	"fmt"
	"regexp"
)

const text = `
My email is ccmouse@gmail.com.@aaads
email2 is aaaaa@gmail.com
email3 is cccccce@gmail.com


`

func main() {
	re, err := regexp.Compile(`([a-zA-Z0-9]+)@([a-zA-Z0-9.]+)\.([a-zA-Z0-9]+)`)
	if err != nil {
		panic(err)
	}
	match := re.FindAllStringSubmatch(text, -1)
	for _, m := range match {
		fmt.Println(m)
	}
	// fmt.Println(match)

}
