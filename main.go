package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

type WeatherReport struct {
	Location  string
	Temp      string
	Scale     string
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
	fmt.Println(inputSlice)

	weather := GetReport(input)
	fmt.Printf("The temp in %s is %s %s and %s.", weather.Location, weather.Temp, weather.Scale, weather.Condition)
}

func PrintHead() {
	fmt.Println("TERMINAL WEATHER REPORT")

}

func GetReport(zipcode string) WeatherReport {
	//url := fmt.Sprintf("https://www.wunderground.com/weather/%s/%s")
	rep := WeatherReport{}
	c := colly.NewCollector()

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Error processing request", err)
	})

	c.OnHTML(".city-header > h1:nth-child(2) > span:nth-child(1)", func(e *colly.HTMLElement) {
		rep.Location = e.Text
	})

	c.OnHTML(".condition-icon > p:nth-child(2)", func(e *colly.HTMLElement) {
		rep.Condition = e.Text
	})

	c.OnHTML(".current=temp > lib-display-unit:nth-child(1) > span:nth-child(1) > span:nth-child(1)", func(e *colly.HTMLElement) {
		rep.Temp = e.Text
	})

	c.OnHTML(".current-temp > lib-display-unit:nth-child(1) > span:nth-child(1) > span:nth-child(2) > span:nth-child(1)", func(e *colly.HTMLElement) {
		rep.Scale = e.Text
	})

	//c.Visit(url)
	return rep
}
