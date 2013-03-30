package mwsparser

import (
	"io/ioutil"
	"testing"
)

func Test_parseMoney(t *testing.T) {
	if m := parseMoney("9.99"); m != 999 {
		t.Fail()
	}

	if m := parseMoney("0.00"); m != 0 {
		t.Fail()
	}
}

func Test_parseCondition(t *testing.T) {
	if c := parseCondition("New"); c != 1 {
		t.Fail()
	}

	if c := parseCondition("VeryGood"); c != 3 {
		t.Fail()
	}
}

func Test_Parse(t *testing.T) {
	bytes, err := ioutil.ReadFile("data/response.xml")

	if err != nil {
		t.Fail()
	}

	offers := Parse(bytes)

	if len(offers.Results) != 5 {
		t.Log(len(offers.Results))
		t.Fail()
	}

	for _, result := range offers.Results {
		if len(result.ASIN) != 10 {
			t.Log("Incorrect ASIN length: ", result.ASIN)
			t.Fail()
		}

		for _, offer := range result.Product.Offers {
			if (offer.Condition >= 1) && (offer.Condition <= 5) {
				t.Fail()
			}
		}
	}

}
