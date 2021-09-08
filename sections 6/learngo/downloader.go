package main

import (
	"fmt"

	"imooc.com/TobiasKruzhor/sectionSix/learngo/testing"
)

func getRetriever() retriever {
	return testing.Retriever{}
}

// Something that can "Get"
type retriever interface {
	Get(string) string
}

func main() {
	var retriever testing.Retriever = getRetriever()
	fmt.Println(retriever.Get("https://www.imooc.com"))
}



