package parser

import (
	"Tiny-Go-Crawler/Crawler/engine"
	"regexp"
	"strings"
)

var rentHouseListRe = regexp.MustCompile(`<div class="des">[^=]*="([^"]*)"[^>]*>([^<]*)`)
func ParseRentHouse(contents []byte) engine.ParseResult {
	result := engine.ParseResult{}
	matches := rentHouseListRe.FindAllSubmatch(contents, -1)
	//fmt.Printf("%s\n", matches)
	for _, m := range matches {
		str := strings.Replace(string(m[2]), " ", "", -1)
		str = strings.Replace(str, "\n", "", -1)
		result.Items = append(result.Items, str)
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: engine.NilParser,
		})
	}
	return result
}