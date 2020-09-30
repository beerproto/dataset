package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gocarina/gocsv"
	"github.com/gocolly/colly/v2"
)

func main() {
	clientsFile, err := os.OpenFile("barthhaas.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
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

	hopMap := map[string]*Hop{}

	c := colly.NewCollector()
	hopColl := colly.NewCollector()

	hopColl.OnHTML("h1", func(e *colly.HTMLElement) {
		hopMap[e.Request.URL.String()].Name = strings.TrimSpace(e.Text)
	})

	hopColl.OnHTML(".pt-3 > div:nth-child(1) > div:nth-child(2) > div:nth-child(1)", func(e *colly.HTMLElement) {
		description := strings.TrimSpace(e.Text)
		description = strings.TrimSpace(strings.TrimLeft(description, "AROMA PROFILE"))
		hopMap[e.Request.URL.String()].Description = description
	})

	hopColl.OnHTML("span.font-type-3", func(e *colly.HTMLElement) {
		hopMap[e.Request.URL.String()].Country = strings.TrimSpace(e.Text)
	})

	hopColl.OnHTML(".key-flavors", func(e *colly.HTMLElement) {

		items := e.DOM.Find("li")
		flavors := []string{}

		for _,  item := range items.Nodes {
			flavors = append(flavors, strings.TrimSpace(item.FirstChild.Data))
		}
		hopMap[e.Request.URL.String()].Flavors = strings.Join(flavors, ", ")
	})

	setValue := func(s string) {

	}


	hopColl.OnHTML("div.col-lg-5:nth-child(1) > dl dt, div.col-lg-5:nth-child(1) > dl dd", func(e *colly.HTMLElement) {

		if e.Name == "dt" {
			switch e.Text {
			case "CULTIVATION AREA":
				setValue = func(s string) {
				}
				break
			case "ALPHA-ACIDS*":
				setValue = func(s string) {
					low, high := split(s)
					hopMap[e.Request.URL.String()].AlphaAcidsLow = low
					hopMap[e.Request.URL.String()].AlphaAcidsHigh = high
				}
				break
			case "BETA-ACIDS":
				setValue = func(s string) {
					low, high := split(s)
					hopMap[e.Request.URL.String()].BetaAcidsLow = low
					hopMap[e.Request.URL.String()].BetaAcidsHigh = high
				}
				break
			case "TOTAL POLYPHENOLS":
				setValue = func(s string) {
					low, high := split(s)
					hopMap[e.Request.URL.String()].TotalPolyphenolsLow = low
					hopMap[e.Request.URL.String()].TotalPolyphenolsHigh = high
				}
				break
			case "TOTAL OIL":
				setValue = func(s string) {
					low, high := split(strings.TrimRight(s, "ML/100G"))
					hopMap[e.Request.URL.String()].TotalOilLow = low
					hopMap[e.Request.URL.String()].TotalOilHigh = high
				}
				break
			case "MYRCENE":
				setValue = func(s string) {
					low, high := split(s)
					hopMap[e.Request.URL.String()].MyrceneLow = low
					hopMap[e.Request.URL.String()].TotalOilLow = high
				}
				break
			case "LINALOOL":
				setValue = func(s string) {
					low, high := split(s)
					hopMap[e.Request.URL.String()].LinaloolLow = low
					hopMap[e.Request.URL.String()].LinaloolHigh = high
				}
				break
			case "SUM OF TERPENE ALCOHOLS":
				setValue = func(s string) {
				}
				break
			case "TOTAL OIL MINUS MYRCENE":
				setValue = func(s string) {
				}
				break
			default:
				fmt.Println(e.Text)
				setValue = func(s string) {
				}
			}
		}

		if e.Name == "dd" {
			setValue(e.Text)
		}
	})

	hopColl.OnHTML("div.col-lg-5:nth-child(1) > dl:nth-child(4) > dd:nth-child(6)", func(e *colly.HTMLElement) {
		low, high := split(e.Text)
		hopMap[e.Request.URL.String()].BetaAcidsHigh = high
		hopMap[e.Request.URL.String()].BetaAcidsLow = low
	})

	hopColl.OnHTML("div.col-lg-5:nth-child(1) > dl:nth-child(4) > dd:nth-child(6)", func(e *colly.HTMLElement) {
		low, high := split(e.Text)
		hopMap[e.Request.URL.String()].BetaAcidsHigh = high
		hopMap[e.Request.URL.String()].BetaAcidsLow = low
	})

	hopColl.OnRequest(func(r *colly.Request) {
		hopMap[r.URL.String()] = &Hop{}
		fmt.Println("Visiting hop", r.URL)
	})
	c.OnHTML("div.hn__results__item a[href]", func(e *colly.HTMLElement) {
		hopColl.Visit("https://www.barthhaas.com" + e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://www.barthhaas.com/en/hopfen/hopfensorten/find-hop-varieties")

	arr := []*Hop{}
	for _, hop := range hopMap {
		arr = append(arr, hop)
	}

	err = gocsv.MarshalFile(&arr, clientsFile)
	if err != nil {
		panic(err)
	}
}

type Hop struct {
	Name                 string `csv:"Name"`
	Flavors              string `csv:"Flavors"`
	AlphaAcidsLow        string `csv:"Alpha-Acids Low (%)"`
	AlphaAcidsHigh       string `csv:"Alpha-Acids High (%)"`
	BetaAcidsLow         string `csv:"Beta-Acids Low (%)"`
	BetaAcidsHigh        string `csv:"Beta-Acids High (%)"`
	TotalOilLow          string `csv:"Total Oil Low (ML/100G)"`
	TotalOilHigh         string `csv:"Total Oil High (ML/100G)"`
	MyrceneLow           string `csv:"Myrcene Low (%)"`
	MyrceneHigh          string `csv:"Myrcene High (%)"`
	LinaloolLow          string `csv:"Linalool Low (%)"`
	LinaloolHigh         string `csv:"Linalool High (%)"`
	TotalPolyphenolsLow  string `csv:"Total Polyphenols Low (%)"`
	TotalPolyphenolsHigh string `csv:"Total Polyphenols High (%)"`
	SumOfTerpeneAlcohols string `csv:"Sum Of Terpene Alcohols"`
	TotalOilMinusMyrcene string `csv:"Total Oil Minus Myrcene"`
	Country              string `csv:"Country"`
	Description          string `csv:"Description"`
}
