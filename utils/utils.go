package utils

import "fmt"

func IsErr(msg string, err error) {
	if err != nil {
		fmt.Println("ERROR:"+msg, err)
		return
	}
}
