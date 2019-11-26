package main

import (
	"bufio"
    "flag"
    "fmt"
    "os"
	"time"
)

type CommandArgs struct {  
    SearchDate 		string
    FilePath 		string
}

func (c CommandArgs) InitFlags() {  
	// Get Date that the forecast is wanted
	loc, _ := time.LoadLocation("UTC")
	dateYesterday := time.Now().In(loc).AddDate(0,0,-1)
	flag.StringVar(&c.SearchDate, "forecast-date", fmt.Sprintf("%v",dateYesterday.Format("2006-01-02")), "The day to get the forecast (Default is Yesterday)")
	flag.StringVar(&c.FilePath, "input-file", "os.Stdin", "The path to a file that you want to read log lines from.")
	flag.Parse()
}

func (c CommandArgs) Validate() bool {  
    return true
}

func (c CommandArgs) ToString() string {
	outstring := fmt.Sprintf("forecast-date: %v\ninput-file: %v",
		c.SearchDate,
		c.FilePath)
	return outstring
}

type ParsingModelFactory interface {
    parse(logline string) ParsingModel
}

type ParsingModel interface {
    toJson() string
    parseLogLine(logline string)
}


type WeatherForcasterModelFactory struct { }

func (f WeatherForcasterModelFactory) parse(logline string) ParsingModel {
     var model WeatherForcasterModel
     model.parseLogLine(logline)
     return model
}

type WeatherForcasterModel struct { 
    rawLine string
}

func (m WeatherForcasterModel) toJson() string {
    return m.rawLine
}

func (m WeatherForcasterModel) parseLogLine(logline string) {
    m.rawLine = logline
}

func main() {
	
	var passedArgs CommandArgs
	
	passedArgs.InitFlags()

	fmt.Println("Command Args:\n", passedArgs.ToString())

	var logLines []ParsingModel
	var logLineFactory WeatherForcasterModelFactory

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	logLines = append(logLines, logLineFactory.parse(text))

	for index, element := range logLines {
		fmt.Println("logline %i\t: %v...",index, element.toJson()[0:20])
	}
}
