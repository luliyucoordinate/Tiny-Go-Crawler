package parser

import (
	"testing"
	"Tiny-Go-Crawler/Crawler/fetcher"
)

func TestParseCityList(t *testing.T) {
	contents, err := fetcher.Fetch("https://www.58.com/changecity.html")
	if err != nil {
		panic(err)
	}

	result := ParseCityList(contents)

	const resultSize = 689
	expectedUrls := []string {
		"https://bj.58.com", "https://sh.58.com", "https://tj.58.com",
	}
	expectedCities := []string {
		"北京", "上海", "天津",
	}
	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d requests; but had %d", resultSize, len(result.Requests))
	}
	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s; but was %s", i, url, result.Requests[i].Url)
		}
	}

	if len(result.Items) != resultSize {
		t.Errorf("result should have %d requests; but had %d", resultSize, len(result.Items))
	}
	for i, city := range expectedCities {
		if result.Items[i].(string) != city {
			t.Errorf("expected city #%d: %s; but was %s", i, city, result.Items[i].(string))
		}
	}
}