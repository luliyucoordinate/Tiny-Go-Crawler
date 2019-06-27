package parser

import (
	"Tiny-Go-Crawler/Crawler/engine"
	"regexp"
	"strings"
)

func ParseCityList(contents []byte) engine.ParseResult {
	provinceList := make([]string, 0)

	re := regexp.MustCompile(`provinceList = {([^}]*)`)
	match := re.FindSubmatch(contents)
	str := strings.Replace(string(match[1]), " ", "", -1)
	str = strings.Replace(str, "\n", "", -1)
	for _, sub := range strings.FieldsFunc(str, splitByComma) {
		provinceList = append(provinceList, strings.Trim(strings.FieldsFunc(sub, splitBySemi)[0], `"`))
	}


	result := engine.ParseResult{}

	re = regexp.MustCompile(`independentCityList = {([^}]*)`)
	match = re.FindSubmatch(contents)
	str = strings.Replace(string(match[1]), " ", "", -1)
	str = strings.Replace(str, "\n", "", -1)
	for _, sub := range strings.FieldsFunc(str, splitByComma) {
		independentCityStr := strings.FieldsFunc(sub, splitBySemi)
		independentCity := strings.Trim(independentCityStr[0], `"`)
		independentCityAb := strings.FieldsFunc(strings.Trim(independentCityStr[1], `"`), splitByVertical)[0]
		independentCityAbUrl := "https://" + independentCityAb + ".58.com"
		result.Items = append(result.Items, independentCity)
		result.Requests = append(result.Requests, engine.Request{
			Url:independentCityAbUrl,
			ParserFunc: engine.NilParser,
		})
	}

	for _, province := range provinceList {
		re := regexp.MustCompile(province + `":{([^}]*)`)
		match = re.FindSubmatch(contents)
		str = strings.Replace(string(match[1]), " ", "", -1)
		str = strings.Replace(str, "\n", "", -1)
		for _, sub := range strings.FieldsFunc(str, splitByComma) {
			cityStr := strings.FieldsFunc(sub, splitBySemi)
			city := strings.Trim(cityStr[0], `"`)
			cityAb := strings.FieldsFunc(strings.Trim(cityStr[1], `"`), splitByVertical)[0]
			cityAbUrl := "https://" + cityAb + ".58.com"
			result.Items = append(result.Items, city)
			result.Requests = append(result.Requests, engine.Request{
				Url:cityAbUrl,
				ParserFunc: engine.NilParser,
			})
		}
	}
	return result
}

func splitByComma(s rune) bool {
	if s == ',' {
		return true
	}
	return false
}

func splitBySemi(s rune) bool {
	if s == ':' {
		return true
	}
	return false
}

func splitByVertical(s rune) bool {
	if s == '|' {
		return true
	}
	return false
}