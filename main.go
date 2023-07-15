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
	input, err := read.ReadString('\n')
	fmt.Println(input)

	if err != nil {
		fmt.Println("Cannot read in country abbrevation and city")
	}

	inputSlice := strings.Split(input, ",")
	//fmt.Println(inputSlice)

	//testStrings()

	weather := GetReport(inputSlice)
	fmt.Printf("The temp in %s is %s and %s.", weather.Location, weather.Temp, weather.Condition)
}

func PrintHead() {
	fmt.Println("TERMINAL WEATHER REPORT")

}

func GetReport(input []string) WeatherReport {
	countryAbbrev := input[0]
	fmt.Println(countryAbbrev)
	city := input[1]
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
	rep := WeatherReport{}
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


func 