package parser

import (
	"testing"
	"io/ioutil"
	"fmt"
)

func TestParseCity(t *testing.T) {
	contents, err := ioutil.ReadFile("city_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseCity(contents)

	fmt.Println(result)
}