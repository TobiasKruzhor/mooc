package main

import (
	"fmt"

	"imooc.com/TobiasKruzhor/sectionSix/learngo/infra"
)

func getRetriever() retriever {
	return infra.Retriever{}
}

// Something that can "Get"
type retriever interface {
	Get(string) string
}

func main() {
	var retriever = getRetriever()
	fmt.Println(retriever.Get("https://www.imooc.com"))
}
