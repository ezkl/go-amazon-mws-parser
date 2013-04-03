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

	if c := parseCondition("Nonsensical"); c != 5 {
		t.Fail()
	}
}

func Test_parseFeedbackRating(t *testing.T) {
	if c := parseFeedbackRating("Just Launched"); c != -1 {
		t.Fail()
	}

	if c := parseFeedbackRating("98-100%"); c != 100 {
		t.Fail()
	}
}

func Test_parseMaxShipping(t *testing.T) {
        if c:= parseMaxShipping("10 or more days"); c != 10 {
                t.Log("Incorrect shipping days: ", c)
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
			if !(offer.Condition >= 1) && !(offer.Condition <= 5) {
				t.Log("Condition Value: ", offer.Condition)
				t.Fail()
			}
		}
	}

}
