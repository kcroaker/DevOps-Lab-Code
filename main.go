package main

import (
    "flag"
	"fmt"
	"time"
)

func main() {

	// Get Date that the forecast is wanted
	loc, _ := time.LoadLocation("UTC")
	dateYesterday := fmt.Sprintf("%v",time.Now().In(loc).AddDate(0,0,-1).Format("2006-01-02"))
    var searchDate string
    flag.StringVar(&searchDate, "forecast-date", dateYesterday, "The day to get the forecast (Default is Yesterday)")
	// Get the data source
	var filePath string
	flag.StringVar(&filePath, "input-file", "", "The path to a file that you want to read log lines from")
	
    flag.Parse()

    fmt.Println("forecast-date:", searchDate)
    fmt.Println("input-file:", filePath)
    fmt.Println("tail:", flag.Args())
}
