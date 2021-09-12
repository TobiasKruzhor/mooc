package main

import (
	"fmt"
	"time"

	"imooc.com/TobiasKruzhor/sectionSix/learngo/retriever/mock"
	"imooc.com/TobiasKruzhor/sectionSix/learngo/retriever/real"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

const url = "http://www.imooc.com"

func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post(
		url,
		map[string]string{
			"name":   "ccmouse",
			"course": "golang",
		},
	)
}

type RetrieverPost interface {
	Retriever
	Poster
}

func session(s RetrieverPost) string {
	s.Post(url, map[string]string{
		"contents": "anoteher faked imooc.com",
	})
	return s.Get(url)
}

func main() {
	var r Retriever
	retriever := mock.Retriever{"this is a fake imooc.com"}
	r = &retriever
	inspect(r)
	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	inspect(r)

	//Type assertion
	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}

	fmt.Println("Try a session")
	fmt.Println(session(&retriever))
}

func inspect(r Retriever) {
	fmt.Println("Inspectiong", r)
	fmt.Printf(" > %T %v\n", r, r)
	fmt.Println(" > Type switch:")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgrent:", v.UserAgent)
	}
	fmt.Println()
}
