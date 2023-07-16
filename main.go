package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
)

type WeatherReport struct {
	Location  string
	Temp      string
	Condition string
}

func main() {
	PrintHead()

	read := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your country abbrevation followed by a comma and then the city")
	fmt.Println("If US based city, please enter country abbrevation followed by a comma and then US State then comma US City")
	input, err := read.ReadString('\n')
	input = strings.ToLower(input)
	fmt.Println(input)

	if err != nil {
		fmt.Println("Cannot read in country abbrevation and city")
	}

	inputSlice := strings.Split(input, ",")
	//fmt.Println(inputSlice)

	//testStrings()

	weather := FormatReport(inputSlice)
	fmt.Printf("The temp in %s is %s and %s.", weather.Location, weather.Temp, weather.Condition)
}

func PrintHead() {
	fmt.Println("TERMINAL WEATHER REPORT")

}

func FormatReport(input []string) WeatherReport {
	countryAbbrev := ""
	fmt.Println(countryAbbrev)
	city := ""
	rep := WeatherReport{}
	var state string
	if len(input) == 3 {
		// includes a us state
		countryAbbrev = input[0]
		state = input[1]
		city = input[2]
		res := GetUSCity(state)

	} else if len(input) == 2 {

	}
	return rep
}

func GetReport() {
	city = strings.TrimSpace(city)
	cityByte := []byte(city)
	//fmt.Println(city)
	matched, err := regexp.Match("[^A-Za-z]", cityByte)
	formattedCityString := ""
	var citySlice []string
	if err != nil {
		fmt.Println("The string doesn't match")
	}
	if matched {
		citySlice = strings.Split(city, " ")
		formattedCityString = strings.Join(citySlice, "-")
	}
	url := fmt.Sprintf("https://www.wunderground.com/weather/%s/%s", countryAbbrev, formattedCityString)
	c := colly.NewCollector()

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Error processing request", err)
	})

	c.OnHTML(".city-header > h1:nth-child(2) > span:nth-child(1)", func(e *colly.HTMLElement) {

		res := strings.Split(e.Text, " Weather Conditions")
		//fmt.Println(res)
		rep.Location = strings.Join(res, " ")
	})

	c.OnHTML(".condition-icon > p:nth-child(2)", func(e *colly.HTMLElement) {
		rep.Condition = e.Text
	})

	c.OnHTML(".current-temp > lib-display-unit:nth-child(1) > span:nth-child(1) > span:nth-child(1)", func(e *colly.HTMLElement) {
		rep.Temp = e.Text
	})

	c.Visit(url)
}

func GetUSCity(city string) string {
	usCityMatch := ""

	switch city {
	case "al":
		usCityMatch = "al"
	case "ak":
		usCityMatch = "ak"
	case "az":
		usCityMatch = "az"
	case "ar":
		usCityMatch = "ar"
	case "as":
		usCityMatch = "as"
	case "ca":
		usCityMatch = "ca"
	case "co":
		usCityMatch = "co"
	case "ct":
		usCityMatch = "ct"
	case "de":
		usCityMatch = "de"
	case "dc":
		usCityMatch = "dc"
	case "fl":
		usCityMatch = "fl"
	case "ga":
		usCityMatch = "ga"
	case "gu":
		usCityMatch = "gu"
	case "hi":
		usCityMatch = "hi"
	case "id":
		usCityMatch = "id"
	case "il":
		usCityMatch = "il"
	case "in":
		usCityMatch = "in"
	case "ia":
		usCityMatch = "ia"
	case "ks":
		usCityMatch = "ks"
	case "ky":
		usCityMatch = "ky"
	case "la":
		usCityMatch = "la"
	case "me":
		usCityMatch = "me"
	case "md":
		usCityMatch = "md"
	case "ma":
		usCityMatch = "ma"
	case "mi":
		usCityMatch = "mi"
	case "mn":
		usCityMatch = "mn"
	case "ms":
		usCityMatch = "ms"
	case "mo":
		usCityMatch = "mo"
	case "mt":
		usCityMatch = "mt"
	case "ne":
		usCityMatch = "ne"
	case "nv":
		usCityMatch = "nv"
	case "nh":
		usCityMatch = "nh"
	case "nj":
		usCityMatch = "nj"
	case "nm":
		usCityMatch = "nm"
	case "ny":
		usCityMatch = "ny"
	case "nc":
		usCityMatch = "nc"
	case "nd":
		usCityMatch = "nd"
	case "mp":
		usCityMatch = "mp"
	case "oh":
		usCityMatch = "oh"
	case "ok":
		usCityMatch = "ok"
	case "or":
		usCityMatch = "or"
	case "pa":
		usCityMatch = "pa"
	case "pr":
		usCityMatch = "pr"
	case "ri":
		usCityMatch = "ri"
	case "sc":
		usCityMatch = "sc"
	case "sd":
		usCityMatch = "sd"
	case "tn":
		usCityMatch = "tn"
	case "tx":
		usCityMatch = "tx"
	case "tt":
		usCityMatch = "tt"
	case "ut":
		usCityMatch = "ut"
	case "vt":
		usCityMatch = "vt"
	case "va":
		usCityMatch = "va"
	case "vi":
		usCityMatch = "vi"
	case "wa":
		usCityMatch = "wa"
	case "wv":
		usCityMatch = "wv"
	case "wi":
		usCityMatch = "wi"
	case "wy":
		usCityMatch = "wy"
	default:
		usCityMatch = ""
		fmt.Println("Error: not a US City")
	}

	return usCityMatch
}

/*
func testStrings() {
	test := "New Delhi, Delhi, India Weather Conditions"
	fmt.Println(test)
	split := strings.Split(test, " Weather Conditions")
	fmt.Println(split)
	res := strings.Join(split, " ")
	fmt.Println(res)

}
*/
