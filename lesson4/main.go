package main

import (
	"Tiny-Go-Crawler/Crawler/engine"
	"Tiny-Go-Crawler/Crawler/58/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:"https://www.58.com/changecity.html",
		ParserFunc: parser.ParseCityList,
	})
}