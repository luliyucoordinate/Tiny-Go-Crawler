package engine

import (
	"Tiny-Go-Crawler/Crawler/fetcher"
	"log"
	"strings"
	"bytes"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		log.Printf("Fetching %s", r.Url)
		body, err := fetcher.Fetch(r.Url)
		ParseHtmlWithFont(&body)
		if err != nil {
			log.Printf("Fetcher: error fetching url %s: %v", r.Url, err)
			continue
		}

		parseResult := r.ParserFunc(body)
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}

func ParseHtmlWithFont(contents *[]byte) {
	FontToUnicode := map[string]string{
		"0x9476": "5",
		"0x958f": "7",
		"0x993c": "0",
		"0x9a4b": "4",
		"0x9e3a": "9",
		"0x9ea3": "1",
		"0x9f64": "6",
		"0x9f92": "2",
		"0x9fa4": "3",
		"0x9fa5": "8",
	}
	for k, v := range FontToUnicode {
		t := strings.Replace(k, "0x", "&#x", -1) + ";"
		*contents = bytes.Replace(*contents, []byte(t), []byte(v), -1)
	}
	*contents = bytes.Replace(*contents, []byte("&nbsp;"), []byte(""), -1)
}

