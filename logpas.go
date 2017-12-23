package main

import (
	"flag"
	"github.com/randomtask1155/logpas/formatter"
	"github.com/randomtask1155/logpas/formatter/lager"
	"os"
	"io"
	"fmt"
)

var (
	logFile = flag.String("l", "", "Specify path to logfile to parse.  If not defined the default is stdin")
	formatType = flag.String("f", "default", "What log component are you reading.  Defaults assumes log formatter is '{}'.\nAvailable components include:")
)

func main() {
	flag.Parse()
	var reader io.Reader
	var err error
	if *logFile == "" {
		reader = os.Stdin
	} else {
		reader, err = os.Open(*logFile)
		if err != nil {
			panic(err)
		}
	}

	var logFormatter interface{}
	if  *formatType == "default" {
		logFormatter = formatter.DefaultFormat{reader, os.Stdout}
	} else if *formatType == "lager" {
		logFormatter = lager.Lager{reader, os.Stdout}
	}
	err = formatter.ParseLog(logFormatter.(formatter.Formatter))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}