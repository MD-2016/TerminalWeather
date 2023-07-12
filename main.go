package main

import (
	"bufio"
	"fmt"
	"os"

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
	fmt.Println("Enter your zipcode to receive your weather report")
	zipcode, err := read.ReadString('\n')

	if err != nil {
		fmt.Println("Cannot read in zipcode")
	}

	weather := GetReport(zipcode)
	fmt.Printf("The temp in %s is %s %s and %s.", weather.Location, weather.Temp, weather.Scale, weather.Condition)
}

func PrintHead() {
	fmt.Println("TERMINAL WEATHER REPORT")

}

func GetReport(zipcode string) WeatherReport {
	url := fmt.Sprintf("https://www.wunderground.com/weather-forecast/%s", zipcode)
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

	c.Visit(url)
	return rep
}
