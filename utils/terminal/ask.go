package terminal

import "fmt"

func Ask(question string) string {
	fmt.Print(question)
	var str string
	_, err := fmt.Scanln(&str)
	if err != nil {
		panic("Error getting value:" + err.Error())
	}
	return str
}
