package main

import (
	"fmt"
	"github.com/beerproto/dataset/hops/hopslist"
	"os"
	"strings"

	"github.com/gocarina/gocsv"
	"github.com/gocolly/colly/v2"
)

func main() {
	clientsFile, err := os.OpenFile("hopslist.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer clientsFile.Close()

	split := func(s string) (low, high string) {
		s = strings.TrimSpace(s)
		arr := strings.Split(s, "-")
		// En Dash (it's bigger)
		if strings.ContainsAny(s, "–") {
			arr = strings.Split(s, "–")
		}

		low = strings.TrimSpace(strings.TrimLeft(strings.TrimRight(strings.TrimSpace(arr[0]), "%"), "<"))

		if low == "Trace Amounts" || low == "None" || low == "Trace" {
			low = ""
		}

		if len(arr) > 1 {
			high = strings.TrimRight(strings.TrimSpace(arr[1]), "%")
		}
		return
	}

	hopMap := map[string]*hopslist.Hop{}

	c := colly.NewCollector()
	hopColl := colly.NewCollector()

	hopColl.OnHTML("h1", func(e *colly.HTMLElement) {
		hopMap[e.Request.URL.String()].Name = strings.TrimSpace(e.Text)
	})

	hopColl.OnHTML(".entry-content > p:nth-child(1)", func(e *colly.HTMLElement) {
		hopMap[e.Request.URL.String()].Description = strings.TrimSpace(e.Text)
	})

	hopColl.OnHTML(".entry-content table > tbody", func(e *colly.HTMLElement) {
		title := e.DOM.Find("tbody:nth-child(1) > tr:nth-child(1) > td:nth-child(1)").Text()
		if strings.TrimSpace(strings.ToLower(title)) != "also known as" {
			return
		}

		alsoKnownAs := e.DOM.Find("tbody > tr:nth-child(1) > td:nth-child(2)").Text()
		hopMap[e.Request.URL.String()].AlsoKnownAs = strings.TrimSpace(alsoKnownAs)

		characteristics := e.DOM.Find("tbody > tr:nth-child(2) > td:nth-child(2)").Text()
		hopMap[e.Request.URL.String()].Characteristics = strings.TrimSpace(characteristics)

		purpose := e.DOM.Find("tbody > tr:nth-child(3) > td:nth-child(2)").Text()
		arr := strings.Split(strings.TrimSpace(purpose), "&")
		for i, s := range arr {
			arr[i] = strings.TrimSpace(s)
		}
		elements := strings.Join(arr, ", ")
		hopMap[e.Request.URL.String()].Purpose = elements

		alphaAcid := e.DOM.Find("tbody > tr:nth-child(4) > td:nth-child(2)").Text()
		low, high := split(alphaAcid)
		hopMap[e.Request.URL.String()].AlphaAcidHigh = high
		hopMap[e.Request.URL.String()].AlphaAcidLow = low

		betaAcid := e.DOM.Find("tbody > tr:nth-child(5) > td:nth-child(2)").Text()
		low, high = split(betaAcid)
		hopMap[e.Request.URL.String()].BetaAcidHigh = high
		hopMap[e.Request.URL.String()].BetaAcidLow = low

		coHumulone := e.DOM.Find("tbody > tr:nth-child(6) > td:nth-child(2)").Text()
		low, high = split(coHumulone)
		hopMap[e.Request.URL.String()].CoHumuloneHigh = high
		hopMap[e.Request.URL.String()].CoHumuloneLow = low

		country := e.DOM.Find("tbody > tr:nth-child(7) > td:nth-child(2)").Text()
		hopMap[e.Request.URL.String()].Country = strings.TrimSpace(country)

		storability := e.DOM.Find("tbody > tr:nth-child(15) > td:nth-child(2)").Text()
		hopMap[e.Request.URL.String()].Storability = strings.TrimSpace(storability)

		totalOil := e.DOM.Find("tbody > tr:nth-child(17) > td:nth-child(2)").Text()
		totalOil = strings.TrimRight(strings.TrimRight(strings.TrimSpace(totalOil), "mL/100g"), "mls/100 grams")
		low, high = split(totalOil)
		hopMap[e.Request.URL.String()].TotalOilHigh = high
		hopMap[e.Request.URL.String()].TotalOilLow = low

		myrceneOil := e.DOM.Find("tbody > tr:nth-child(18) > td:nth-child(2)").Text()
		low, high = split(myrceneOil)
		hopMap[e.Request.URL.String()].MyrceneOilHigh = high
		hopMap[e.Request.URL.String()].MyrceneOilLow = low

		humuleneOil := e.DOM.Find("tbody > tr:nth-child(19) > td:nth-child(2)").Text()
		low, high = split(humuleneOil)
		hopMap[e.Request.URL.String()].HumuleneOilHigh = high
		hopMap[e.Request.URL.String()].HumuleneOilLow = low

		caryophylleneOil := e.DOM.Find("tbody > tr:nth-child(20) > td:nth-child(2)").Text()
		low, high = split(caryophylleneOil)
		hopMap[e.Request.URL.String()].CaryophylleneOilHigh = high
		hopMap[e.Request.URL.String()].CaryophylleneOilLow = low

		farneseneOil := e.DOM.Find("tbody > tr:nth-child(21) > td:nth-child(2)").Text()
		low, high = split(farneseneOil)
		hopMap[e.Request.URL.String()].FarneseneOilHigh = high
		hopMap[e.Request.URL.String()].FarneseneOilLow = low

		substitutes := e.DOM.Find("tbody > tr:nth-child(22) > td:nth-child(2)").Text()
		hopMap[e.Request.URL.String()].Substitutes = strings.TrimSpace(substitutes)

		styleGuide := e.DOM.Find("tbody > tr:nth-child(23) > td:nth-child(2)").Text()
		hopMap[e.Request.URL.String()].StyleGuide = strings.TrimSpace(styleGuide)
	})

	hopColl.OnRequest(func(r *colly.Request) {
		hopMap[r.URL.String()] = &hopslist.Hop{}
		fmt.Println("Visiting hop", r.URL)
	})

	c.OnHTML("div.x-column ul li a[href]", func(e *colly.HTMLElement) {
		hopColl.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("http://www.hopslist.com/hops/")

	arr := []*hopslist.Hop{}
	for _, hop := range hopMap {
		arr = append(arr, hop)
	}

	err = gocsv.MarshalFile(&arr, clientsFile)
	if err != nil {
		panic(err)
	}
}
