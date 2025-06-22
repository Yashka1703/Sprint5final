package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(data string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, i := range dataset {
		err := dp.Parse(i)
		if err != nil {
			log.Printf("parsing error: %v", err)
		}
		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("information error: %v", err)
		}
		fmt.Println(info)
	}
}
