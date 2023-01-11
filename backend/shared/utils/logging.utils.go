package utils

import (
	"fmt"
	"log"

	"github.com/fatih/color"
)

const (
	DEFAULT_LOG= iota
	WARNING_LOG= iota
	ERROR_LOG= iota
)

type LogDetails struct {

	Type int
	Message interface{ }
	Method string
}

var boldYellowText= color.New(color.FgHiYellow).Add(color.Bold)

func Log(details LogDetails) {
	var message string

	switch inputMessage := details.Message.(type) {
		case error:
			message= inputMessage.Error( )

		case string:
			message= inputMessage

		default:
	}

	log.Printf("in %s", boldYellowText.Sprint(details.Method, "( )"))

	switch details.Type {

		case WARNING_LOG:
			color.HiYellow(message)

		case DEFAULT_LOG:
			color.HiBlue(message)

		//* default log type is error
		default:
			color.HiRed(message)
	}

	fmt.Println( )
}