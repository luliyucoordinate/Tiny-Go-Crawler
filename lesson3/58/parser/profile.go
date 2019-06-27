package parser

import (
	"Tiny-Go-Crawler/Crawler/engine"
	"regexp"
	"Tiny-Go-Crawler/Crawler/model"
	"strconv"
	"strings"
)


var rentRe = regexp.MustCompile(`<span class="c_ff552e">[^>]*>(\d*)`)
var rentalMethodRe = regexp.MustCompile(`<span[^>]*>租赁方式：</span><span>(.*)</span>`)
var propertyTypeRe = regexp.MustCompile(`<span[^>]*>房屋类型：</span><span[^>]*>(.*)</span>`)
var imageUrlRe = regexp.MustCompile(`<img lazy_src="([^"]*)"`)
var floorOrientedRe = regexp.MustCompile(`<span[^>]*>朝向楼层：</span><span[^>]*>(.*)</span>`)
var addressRe = regexp.MustCompile(`<span[^"]*"dz"[^>]*>[\n|\s]*([^\s]*)[\n|\s]*</span>`)
var communityRe = regexp.MustCompile(`<a class="c_333 ah"[^>]*>(.*)</a>`)
var propertyCompanyRe = regexp.MustCompile(`<span[^>]*>物业公司：</span>[^>]*>(.*)</span>`)
var propertyCostsRe = regexp.MustCompile(`<span[^>]*>物业费用：</span>[^>]*>(.*)</span>`)



func ParseProfile(contents []byte, name string) engine.ParseResult {


	profile := model.Profile{}
	profile.Name = name
	rent, err := strconv.Atoi(extractString(contents, rentRe))
	if err == nil {
		profile.Rent = rent
	}

	profile.RentalMethod = extractString(contents, rentalMethodRe)
	profile.PropertyType = strings.Replace(extractString(contents, propertyTypeRe), " ", "", -1)
	profile.ImageUrl = strings.Replace(extractString(contents, imageUrlRe), "&amp", "&", -1)

	floorOriented := extractString(contents, floorOrientedRe)
	profile.Floor = string([]rune(floorOriented)[1:])
	profile.Oriented = string([]rune(floorOriented)[:1])
	profile.Address = extractString(contents, addressRe)
	profile.Community = extractString(contents, communityRe)
	profile.PropertyCompany = extractString(contents, propertyCompanyRe)
	profile.PropertyCosts = extractString(contents, propertyCostsRe)

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}

