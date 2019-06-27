package parser

import (
	"testing"
	"io/ioutil"
	"fmt"
)

func TestRentHouse(t *testing.T) {
	contents, err := ioutil.ReadFile("renthouse_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseRentHouse(contents)
	fmt.Println(result)
}