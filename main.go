package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
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
	fmt.Println("If US based city, please enter country abbrevation followed by a comma and then US State abbrevation then comma US City")
	input, err := read.ReadString('\n')
	input = strings.ToLower(input)

	if err != nil {
		fmt.Println("Cannot read in input")
	}

	inputSlice := strings.Split(input, ",")
	countryAbbrev, state, city := FormatReport(inputSlice)
	if countryAbbrev == "" && city == "" {
		fmt.Println("Cannot process request due to invalid input")
	} else {
		weather := GetReport(countryAbbrev, state, city)
		temp, err := strconv.ParseFloat(weather.Temp, 64)
		if err != nil {
			fmt.Println("error processing temp to celsius")
		}
		celsius := GetCelsiusTemp(temp)
		fmt.Printf("The temp in %s is %s° F or %d° C and conditions are %s\n", weather.Location, weather.Temp, celsius, weather.Condition)
	}

}

func PrintHead() {
	fmt.Println("TERMINAL WEATHER REPORT")

}

func FormatReport(input []string) (string, string, string) {
	countryAbbrev := ""
	city := ""
	state := ""

	if len(input) != 2 && len(input) != 3 {
		fmt.Println("Error: must have a country abbrev, city or if us a us state abbrev and us city")
		return "", "", ""
	}

	if len(input) == 3 {
		countryAbbrev = input[0]
		state = input[1]
		city = input[2]
		countryAbbrev, state, city = FormatUSCityString(countryAbbrev, state, city)
		return countryAbbrev, state, city
	}

	if len(input) == 2 {
		countryAbbrev = input[0]
		city = input[1]
		countryAbbrev, city = FormatCountryCityString(countryAbbrev, city)
		return countryAbbrev, "", city
	}

	return "", "", ""

}

func FormatCountryCityString(country string, city string) (string, string) {
	re, _ := regexp.Compile("[A-Za-z]+")
	country = strings.TrimSpace(country)
	city = strings.TrimSpace(city)

	if re.MatchString(country) && re.MatchString(city) {
		city = strings.ReplaceAll(city, " ", "-")
		return country, city
	} else {
		return "", ""
	}
}

func FormatUSCityString(country string, state string, city string) (string, string, string) {
	re, _ := regexp.Compile("[A-Za-z]+")
	country = strings.TrimSpace(country)
	city = strings.TrimSpace(city)
	state = strings.TrimSpace(state)

	if re.MatchString(country) && re.MatchString(city) && re.MatchString(state) {
		state = GetUSCity(state)
		if state == "" {
			fmt.Println("State cannot be blank")
			return "", "", ""
		} else {
			city = strings.ReplaceAll(city, " ", "-")
			return country, state, city
		}
	} else {
		return "", "", ""
	}
}

/*
func FormatReport(input []string) (string, string, string) {
	countryAbbrev := ""
	city := ""
	var state string
	var formattedCityString string
	countryAbbrev = input[0]
	if len(input) == 3 {
		// includes a us state
		city = input[2]
		state = GetUSCity(input[1])

		if state == "" {
			fmt.Println("Error: state cannot be blank")
			return "", "", ""
		}

		city = strings.TrimSpace(city)
		//fmt.Println(city)
		matched, err := regexp.MatchString("[A-Za-z]+", city)
		if err != nil {
			fmt.Println("The string doesn't match")
			return "", "", ""
		}
		if matched {

			//check for more than one word
			morethanOne, err := regexp.MatchString("\\s+", city)
			if err != nil {
				fmt.Println("The string is errored")
				return "", "", ""
			}
			if morethanOne {
				formattedCityString = strings.ReplaceAll(city, " ", "-")
			} else {
				formattedCityString = city
			}
		}

		return countryAbbrev, state, formattedCityString
	} else if len(input) == 2 {
		city = input[1]
		state = ""
		city = strings.TrimSpace(city)
		//fmt.Println(city)
		matched, err := regexp.MatchString("[A-Za-z]+", city)
		if err != nil {
			fmt.Println("The string doesn't match")
		}
		if matched {
			//check for more than one word
			morethanOne, err := regexp.MatchString("\\s+", city)
			if err != nil {
				fmt.Println("The string is errored")
				return "", "", ""
			}
			if morethanOne {
				formattedCityString = strings.ReplaceAll(city, " ", "-")
			} else {
				formattedCityString = city
			}
		}
		return countryAbbrev, "", formattedCityString
	} else {
		fmt.Println("Error: must have a country abbrev, city or country abbrev, us state abbrev, and us city.")
		return "", "", ""
	}

}
*/

func GetReport(country string, state string, city string) WeatherReport {
	rep := WeatherReport{}
	url := ""
	if state != "" {
		url = fmt.Sprintf("https://www.wunderground.com/weather/%s/%s/%s", country, state, city)
	} else {
		url = fmt.Sprintf("https://www.wunderground.com/weather/%s/%s", country, city)
	}
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
	return rep
}

func GetUSCity(city string) string {
	states := []string{"al", "ak", "az", "ar", "ca", "co", "ct", "de", "fl", "ga", "hi", "id", "il", "in", "ia", "ks", "ky", "la", "me", "md", "ma", "mi", "mn", "ms", "mo", "mt", "ne", "nv", "nh", "nj", "nm", "ny", "nc", "nd", "oh", "ok", "or", "pa", "ri", "sc", "sd", "tn", "tx", "ut", "vt", "va", "wa", "wi", "wv", "wy", "as", "dc", "gu", "mp", "pr", "vi"}

	for _, state := range states {
		if state == city {
			return city
		}
	}

	fmt.Println("State does not match")
	return ""

}

/*
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
		fmt.Println("Error: not a US State")
	}

	return usCityMatch
}
*/

func GetCelsiusTemp(temp float64) int {

	celsius := math.Ceil((temp - 32) * (5.0 / 9.0))
	return int(celsius)
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
