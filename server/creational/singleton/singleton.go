package singleton

import (
	"fmt"
	"sync"
)

var once sync.Once

type single struct {
}

var singleInstance *single

func GetSingleTonObject() *single {
	if singleInstance == nil {
		once.Do(
			func() {
				fmt.Println("Creting Single Instance Now")
				singleInstance = &single{}
			})
	} else {
		fmt.Println("Single Instance already created-2")
	}
	return singleInstance
}
