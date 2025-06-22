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
		err := dp.Parse(i) // распарсинг значения
		if err != nil {
			log.Printf("parsing error: %v", err) // вывод ошибки в логи
		}
		info, err := dp.ActionInfo() // формирование и вывод информации
		if err != nil {
			log.Printf("information error: %v", err) // вывод ошибки в логи
		}
		fmt.Println(info)
	}
}
