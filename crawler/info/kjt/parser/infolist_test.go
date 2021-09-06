package parser

import (
	"io/ioutil"
	"testing"

	
)

func TestParseInfoList(t *testing.T) {
	contents, err := ioutil.ReadFile("infolist_test_data.html")

	if err != nil {
		panic(err)
	}
	result := ParseInfoList(contents)

	const resultSize = 75
	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d "+"requeste; but had %d",
			resultSize, len(result.Requests))
	}
	if len(result.Items) != resultSize {
		t.Errorf("result should have %d "+"requeste; but had %d",
			resultSize, len(result.Items))
	}

	// verify result

}
