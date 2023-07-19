# TerminalWeather
A terminal application where you can request weather for an area of choice. This project scrapes data from a weather source [wunderground.com](https://www.wunderground.com) which means there are limitations to running the application that will be discussed below. The overall goal was to try out the [colly framework](https://go-colly.org/) for web scraping and using golang for a meteorology based project. 

## Table of Contents
1. [Instructions](#instructions)
2. [Technologies Used](#technologies-used)
3. [Project Goals](#project-goals)
4. [Findings](#findings)
5. [Reported Errors](#reported-errors)
5. [Conclusions](#conclusions)

## Instructions 
## Technologies Used
1. [Go](https://go.dev)
2. [Colly](https://go-coloy.org)

## Project Goals
- [ ] Properly format the URL to be used in the GET request for the weather data
- [ ] Target the correct parts of the web page to scrape the data
- [ ] Display the returned results to the user in the terminal
- [ ] Handle one special case for US where the url has to be formatted to handle US States and their cities
- [ ] Convert the temperature result from Fahrenheit to Celsius and display

## Findings
One finding during this project so far was the issue of not every city being recognized. After searching on the site myself for the reason, it seems the data comes from weather stations that not every city has or has register to the site. This means you need an alternative way of getting weather for your city. The trick I found is to look up the airport code for the closest airport and substitute that for the respective city.

## Reported Errors
## Conclusions
