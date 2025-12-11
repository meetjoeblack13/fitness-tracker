package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		err := dp.Parse(data)
		if err != nil {
			log.Println("Ошибка парсинга:", err)
			continue
		}
		information, err := dp.ActionInfo()
		if err != nil {
			log.Println("Ошибка формирования информации:", err)
			continue
		}
		fmt.Println(information)
	}
}
