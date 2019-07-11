package cassango

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

const (
	timeColor    = "\033[1;34m%s\033[0m"
	queryColor   = "\033[1;36m%s\033[0m"
	resultsColor = "\033[1;33m%s\033[0m"
	errorColor   = "\033[1;31m%s\033[0m"
	runtimeColor = "\033[0;36m%s\033[0m"
)

type logger struct {
	format  string `json:"-"`
	Time    string `json:"time"`
	Query   string `json:"query"`
	Results int    `json:"results"`
	Error   error  `json:"error"`
	Runtime int    `json:"runtime"`
}

func (l *logger) printLog() {
	l.Time = time.Now().Format("2006-01-02 03:04:05")

	if l.format == "json" {
		j, _ := json.Marshal(l)
		fmt.Println(string(j))
	} else {
		fmt.Printf(timeColor, "["+l.Time+"]")
		fmt.Printf(" ")
		fmt.Printf(queryColor, "["+l.Query+"]")
		fmt.Printf(" ")
		fmt.Printf(resultsColor, "[results: "+strconv.Itoa(l.Results)+"]")
		fmt.Printf(" ")
		fmt.Printf(runtimeColor, "[runtime: "+strconv.Itoa(l.Runtime)+"]")
		if l.Error != nil {
			fmt.Printf(" ")
			fmt.Printf(errorColor, "["+l.Error.Error()+"]")
		}
		fmt.Println("")
	}
}
