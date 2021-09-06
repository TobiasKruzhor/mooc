package parser

import (
	"regexp"

	"github.com/TobiasKruzhor/crawler/engine"
)

const infoListRe = `<li><span.*>(.*)</span><a.*href="([^"]+)".*</s>(.*)</a></li>`

func ParseInfoList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(infoListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(
			result.Items, string(m[3]))
		result.Requests = append(
			result.Requests, engine.Request{
				Url:        string(m[2]),
				ParserFunc: engine.NilParser,
			})
		result.Time = append(
			result.Time, string(m[1]))
	}
	return result
}
