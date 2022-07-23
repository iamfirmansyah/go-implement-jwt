package helper

import "fmt"

func PanicIfError(err error) {
	if err != nil {
		fmt.Println(err)
		// panic(err)
	}
}

func ErrorIfDataEmpty(data interface{}) bool {
	if data == nil {
		return true
	}

	return false
}
