# Simple Terminal Weather Application
A terminal application where you can request weather for an area of choice. This project scrapes data from a weather source [wunderground.com](https://www.wunderground.com) which means there are limitations to running the application that will be discussed below. The overall goal was to try out the [colly framework](https://go-colly.org/) for web scraping and using golang for a meteorology based project. 

## Table of Contents
1. [Instructions](#instructions)
2. [Technologies Used](#technologies-used)
3. [Project Goals](#project-goals)
4. [Findings](#findings)
5. [Reported Errors](#reported-errors)
5. [Conclusions](#conclusions)

## Instructions 
1. Make sure you have the Go programming language installed from [Go](https://go.dev). Note that this application has only been tested on linux and not windows. It's possible that it could have issues on windows or possibly MacOS, but you're welcome try.
2. Install [git](https://git-scm.com) from your package manager on linux and use `git clone` this link [TerminalWeather](https://github.com/MD-2016/TerminalWeather.git), or proceed to this link here and download the zip [TerminalWeather](https://github.com/MD-2016/TerminalWeather) and extract it to a directory of your choice.
3. After unzipping or cloning the repo, change directories in your respective terminal to the project directory. 
4. Once in the `TerminalWeather` directory, run the `terminalweather` executable by typing `./terminalweather` in your respective terminal.
5. After running the `./terminalweather` in your terminal, you'll get the following screen ![Screenshot of terminal](./Screenshot%20at%202023-07-19%2019-33-31.png). 
6. Once here, you can do either of the following...
    - Enter a country abbrevation followed by a comma, then enter the name of the city of said country and hit enter.
    - Or if you have a US city, then you enter the country abbrevation (in this case `us`) followed by a comma, then the US state abbrevation (for example `oh` for Ohio), followed by another comma and your city (example `Columbus`)

    Here's two examples of these above steps in action...
    ![Screenshot of country abbrevation followed by city](./Terminal%20Weather%20Screenshot%202.png)
    ![Screenshot of country abbrevation followed by US state and US city](./Terminal%20Weather%20Screenshot%203.png)

Please keep in mind that proper format in either case is as follows `[country abbrevation],[city]` or `[country abbrevation],[US state abbrevation],[US city]` as in the above example screenshots.

*NOTE*: It is possible that you get incorrect data or no data at all because the city of interest doesn't actually exist on their site. This mean that the weather station isn't online or the respective city doesn't have an active weather station (this site gets data from weather stations). This means you may need to research your city of interest by looking for an active PWS or Personal Weather Station on their site. Unfortunately, this is an issue that cannot be resolved currently. It is also possible that timeouts occur.  

7. You get weather reported to the screen in both fahrenheit and celsius and the program ends.

**DISCLAIMER**: This program scrapes weather data from the site mentioned in the introduction. This means too many requests could result in your ip blacklisted. The developer of this program is not responsible if you abuse the program. You claim all responsibiilty if your ip gets banned from using the site. 




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
One finding during this project so far was the issue of not every city being recognized. After searching on the site myself for the reason, it seems the data comes from weather stations that not every city has or has register to the site. This means you need an alternative way of getting weather for your city. This means you may need to research the city for the closest PWS.
## Reported Errors
## Conclusions
This project was fun and it was a good introduction to web scraping with the go programming language. Helped understand a bit more about how the go language works and how to use it for web scraping. 
