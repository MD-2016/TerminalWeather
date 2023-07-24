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

	if len(input) == 3 {
		countryAbbrev = input[0]
		state = input[1]
		city = input[2]
		countryAbbrev, state, city = FormatUSCityString(countryAbbrev, state, city)
		return countryAbbrev, state, city
	}

	if len(input) == 2 {
		countryAbbrev = input[0]
		if countryAbbrev == "us" {
			fmt.Println("US requires 3 inputs")
			countryAbbrev = ""
			return countryAbbrev, state, city
		}
		city = input[1]
		countryAbbrev, city = FormatCountryCityString(countryAbbrev, city)
		return countryAbbrev, state, city
	}

	fmt.Println("Error: must have a country abbrev, city or if us a us state abbrev and us city")
	return countryAbbrev, city, state

}

func FormatCountryCityString(country string, city string) (string, string) {
	re, _ := regexp.Compile("[A-Za-z]+")
	country = strings.TrimSpace(country)
	country = CheckCountryAbbrev(country)
	city = strings.TrimSpace(city)

	if re.MatchString(country) && re.MatchString(city) {
		city = strings.ReplaceAll(city, " ", "-")
		return country, city
	} else {
		country = ""
		return country, city
	}
}

func FormatUSCityString(country string, state string, city string) (string, string, string) {
	re, _ := regexp.Compile("[A-Za-z]+")
	country = strings.TrimSpace(country)
	country = CheckCountryAbbrev(country)
	city = strings.TrimSpace(city)
	state = strings.TrimSpace(state)

	if re.MatchString(country) && re.MatchString(city) && re.MatchString(state) {
		state = GetUSCity(state)
		if state == "" {
			fmt.Println("State cannot be blank")
			country = ""
			return country, state, city
		} else {
			city = strings.ReplaceAll(city, " ", "-")
			return country, state, city
		}
	} else {
		country = ""
		return country, state, city
	}
}

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
	city = ""
	return city

}

func CheckCountryAbbrev(country string) string {
	countries := []string{"ad", "ae", "af", "ag", "ai", "al", "am", "ao", "aq", "ar", "as", "at", "au", "aw", "ax", "az", "ba", "bb", "bd", "be", "bf", "bg", "bh", "bi", "bj", "bl", "bm", "bn", "bo", "bq", "br", "bs", "bt", "bv", "bw", "by", "bz", "ca", "cc", "cd", "cf", "cg", "ch", "ci", "ck", "cl", "cm", "cn", "co", "cr", "cu", "cv", "cw", "cx", "cy", "cz", "de", "dj", "dk", "dm", "do", "dz", "ec", "ee", "eg", "eh", "er", "es", "et", "fi", "fj", "fk", "fm", "fo", "fr", "ga", "gb", "gd", "ge", "gf", "gg", "gh", "gi", "gl", "gm", "gn", "gp", "gq", "gr", "gs", "gt", "gu", "gw", "gy", "hk", "hm", "hn", "hr", "ht", "hu", "id", "ie", "il", "im", "in", "io", "iq", "ir", "is", "it", "je", "jm", "jo", "jp", "ke", "kg", "kh", "ki", "km", "kn", "kp", "kr", "kw", "ky", "kz", "la", "lb", "lc", "li", "lk", "lr", "ls", "lt", "lu", "lv", "ly", "ma", "mc", "md", "me", "mf", "mg", "mh", "mk", "ml", "mm", "mn", "mo", "mp", "mq", "mr", "ms", "mt", "mu", "mv", "mw", "mx", "my", "mz", "na", "nc", "ne", "nf", "ng", "ni", "nl", "no", "np", "nr", "nu", "nz", "om", "pa", "pe", "pf", "pg", "ph", "pk", "pl", "pm", "pn", "pr", "ps", "pt", "pw", "py", "qa", "re", "ro", "rs", "ru", "rw", "sa", "sb", "sc", "sd", "se", "sg", "sh", "si", "sj", "sk", "sl", "sm", "sn", "so", "sr", "ss", "st", "sv", "sx", "sy", "sz", "tc", "td", "tf", "tg", "th", "tj", "tk", "tl", "tm", "tn", "to", "tr", "tt", "tv", "tw", "tz", "ua", "ug", "um", "us", "uy", "uz", "va", "vc", "ve", "vg", "vi", "vn", "vu", "wf", "ws", "ye", "yt", "za", "zm", "zw"}

	for _, selectedCountry := range countries {
		if selectedCountry == country {
			return country
		}
	}

	fmt.Println("Country doesn't match the two letter abbrevation for countries in the world")
	country = ""
	return country
}

func GetCelsiusTemp(temp float64) int {

	celsius := math.Ceil((temp - 32) * (5.0 / 9.0))
	return int(celsius)
}
