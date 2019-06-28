package parser

import (
	"Tiny-Go-Crawler/Crawler/engine"
	"regexp"
)

var localRe = regexp.MustCompile(`locallist:.*listname: '([^']*)'}`)
var rentHouseRe = regexp.MustCompile(`<a href="([^"]*)"[^>]*>租房</a>`)

func ParseCity(contents []byte) engine.ParseResult {
	result := engine.ParseResult{}
	local := localRe.FindSubmatch(contents)
	rent := rentHouseRe.FindSubmatch(contents)
	parseCityUrl := "https://" + string(local[1]) + ".58.com" + string(rent[1])
	result.Items = append(result.Items, string(rent[1]))
	result.Requests = append(result.Requests, engine.Request{
		Url: parseCityUrl,
		ParserFunc: ParseRentHouse,
	})
	return result
}
