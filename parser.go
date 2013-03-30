package mwsparser

import (
	"encoding/xml"
	"log"
	"regexp"
	"strconv"
)

type Document struct {
	XMLName xml.Name `xml:"GetLowestOfferListingsForASINResponse"`
	Results []Result `xml:"GetLowestOfferListingsForASINResult"`
}

type Result struct {
	XMLName xml.Name `xml:"GetLowestOfferListingsForASINResult"`
	ASIN    string   `xml:"ASIN,attr"`
	Status  string   `xml:"status,attr"`
	Product *Product
}

type Product struct {
	XMLName xml.Name `xml:"Product"`
	Offers  []Offer  `xml:"LowestOfferListings>LowestOfferListing"`
}

type Offer struct {
	XMLName              xml.Name `xml:"LowestOfferListing"`
	ConditionString      string   `xml:"Qualifiers>ItemSubcondition"`
	Condition            int
	DomesticString       string `xml:"Qualifiers>ShipsDomestically"`
	Domestic             bool
	ShippingTimeString   string `xml:"Qualifiers>ShippingTime>Max"`
	ShippingTime         int
	FeedbackRatingString string `xml:"Qualifiers>SellerPositiveFeedbackRating"`
	FeedbackRating       int
	FeedbackCount        int    `xml:"SellerFeedbackCount"`
	ListingPriceString   string `xml:"Price>ListingPrice>Amount"`
	ShippingPriceString  string `xml:"Price>Shipping>Amount"`
	ListingPrice         int
	ShippingPrice        int
}

func parseMoney(priceStr string) (priceInt int) {
	if priceStr == "0.00" {
		priceInt = 0
		return
	}

	p, err := strconv.ParseFloat(priceStr, 64)

	if err != nil {
		log.Fatal(err)
	}

	priceInt = int(p * 100.0)

	return
}

func parseCondition(condStr string) int {
	switch condStr {
	case "New":
		return 1
	case "Mint":
		return 2
	case "VeryGood":
		return 3
	case "Good":
		return 4
	}
	return 5
}

func parseDomestic(domStr string) bool {
	switch domStr {
	case "True":
		return true
	case "False":
		return false
	}
	return false
}

func parseFeedbackRating(fbStr string) int {
	ratingRegex, err := regexp.Compile(`([\d]+)%$`)

	if err != nil {
		log.Fatal(err)
	}

	if ratingRegex.Match([]byte(fbStr)) {
		m := ratingRegex.FindAllStringSubmatch(fbStr, 1)
		i, err := strconv.Atoi(m[0][1])

		if err != nil {
			log.Fatal("Couldn't parse Feedback Rating: ", err)
		}

		return i
	}

	return -1
}

func parseMaxShipping(shipStr string) int {
	maxRegex, err := regexp.Compile(`([\d]+) (?:or more)?days$`)

	if err != nil {
		log.Fatal(err)
	}

	if maxRegex.Match([]byte(shipStr)) {
		m := maxRegex.FindAllStringSubmatch(shipStr, 1)
		i, err := strconv.Atoi(m[0][1])
		if err != nil {
			log.Fatal("Couldn't parse Max Shipping: ", err)
		}
		return i
	}

	return 100
}

func Parse(body []byte) (mws Document) {
	mws = Document{}

	if err := xml.Unmarshal(body, &mws); err != nil {
		log.Fatal(err)
	}

	for _, result := range mws.Results {
		for k, o := range result.Product.Offers {
			result.Product.Offers[k].ListingPrice = parseMoney(o.ListingPriceString)
			result.Product.Offers[k].ShippingPrice = parseMoney(o.ShippingPriceString)

			result.Product.Offers[k].Condition = parseCondition(o.ConditionString)
			result.Product.Offers[k].Domestic = parseDomestic(o.DomesticString)
			result.Product.Offers[k].ShippingTime = parseMaxShipping(o.ShippingTimeString)
		}
	}

	return mws
}
