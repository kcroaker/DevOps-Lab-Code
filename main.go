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

func main() {
	
	var passedArgs CommandArgs
	
	passedArgs.initFlags()

	fmt.Println("Command Args:\n", passedArgs.ToString())

	reader := bufio.NewReader(os.Stdin)
    text, _ := reader.ReadString('\n')
    fmt.Println(text)
}
