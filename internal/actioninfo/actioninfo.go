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
	for _, data := range dataset {
		err := dp.Parse(data)
		if err != nil {
			// логируем ошибку
			log.Println(err)
			continue
		}
		info, _ := dp.ActionInfo()
		fmt.Println(info)
	}
}
