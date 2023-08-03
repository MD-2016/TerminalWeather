package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/MD-2016/TerminalWeather/formatweatherreport"
	"github.com/gocolly/colly"
)

type WeatherReport struct {
	Location  string
	Temp      string
	Condition string
}

const WEBSITE_URL = "https://www.wunderground.com/weather/"

func main() {
	fmt.Println("TERMINAL WEATHER REPORT")

	input, err := handleInput()
	if err != nil {
		panic(err.Error())
	}
	re, _ := regexp.Compile("[A-Za-z,]+")
	if inputValid := re.MatchString(input); !inputValid {
		panic("Input must only contain letters")
	}

	inputSlice := strings.Split(input, ",")
	data, err := FormatReport(inputSlice)
	if err != nil {
		panic(err.Error())
	}

	if err := formatweatherreport.ValidateInput(data); err != nil {
		panic(err.Error())
	}

	weatherReport := GetReport(data)
	if weatherReport.Location == "undefined" {
		fmt.Println("the location either doesn't exist or isn't supported by the site \n")
		fmt.Println("check wunderground site for a PWS closest to that city")
		os.Exit(0)
	}
	temp, err := strconv.ParseFloat(weatherReport.Temp, 64)
	if err != nil {
		fmt.Println("error processing temp")
	}
	celsius := GetCelsiusTemp(temp)
	fmt.Printf("The temp in %s is %s° F or %d° C and conditions are %s\n", weatherReport.Location, weatherReport.Temp, celsius, weatherReport.Condition)

}

func handleInput() (string, error) {
	read := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your country abbrevation followed by a comma and then city")
	fmt.Println("US Cities require 3 inputs which are us, state, city")
	input, err := read.ReadString('\n')
	input = strings.ToLower(input)

	if err != nil {
		return input, fmt.Errorf("Input cannot be processed")
	}

	return input, err
}

func FormatReport(input []string) (formatweatherreport.ReportData, error) {
	var rd = formatweatherreport.ReportData{}

	if len(input) == 3 {
		rd.CountryAbbrev = input[0]
		rd.State = input[1]
		rd.City = input[2]
	} else if len(input) == 2 {
		rd.CountryAbbrev = input[0]
		rd.City = input[1]
		if rd.CountryAbbrev == "us" {
			return rd, fmt.Errorf("US requires 3 inputs")
		}
	} else {
		return rd, fmt.Errorf("must have country abbrev, city or us, state, city")
	}

	formatweatherreport.FormatDataForReport(&rd)
	return rd, nil
}

func GetReport(data formatweatherreport.ReportData) WeatherReport {
	url := ""
	report := WeatherReport{}
	if data.State != "" {
		url = fmt.Sprintf(WEBSITE_URL+"/%s/%s/%s", data.CountryAbbrev, data.State, data.City)
	} else {
		url = fmt.Sprintf(WEBSITE_URL+"/%s/%s", data.CountryAbbrev, data.City)
	}
	c := colly.NewCollector()

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Error processing request", err)
	})

	c.OnHTML(".city-header > h1:nth-child(2) > span:nth-child(1)",
		func(e *colly.HTMLElement) {

			res := strings.Split(e.Text, " Weather Conditions")
			//fmt.Println(res)
			report.Location = strings.Join(res, " ")
		})

	c.OnHTML(".condition-icon > p:nth-child(2)", func(e *colly.HTMLElement) {
		report.Condition = e.Text
	})

	c.OnHTML(".current-temp > lib-display-unit:nth-child(1) > span:nth-child(1) > span:nth-child(1)",
		func(e *colly.HTMLElement) {
			report.Temp = e.Text
		})

	c.Visit(url)
	return report
}

func GetCelsiusTemp(temp float64) int {
	celsius := math.Ceil((temp - 32) * (5.0 / 9.0))
	return int(celsius)
}
