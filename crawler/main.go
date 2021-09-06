package main

import (
	"github.com/TobiasKruzhor/crawler/engine"
	"github.com/TobiasKruzhor/crawler/info/kjt/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://kjt.shandong.gov.cn/col/col13360/index.html",
		ParserFunc: parser.ParseInfoList,
	})
}
