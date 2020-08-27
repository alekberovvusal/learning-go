package util

import (
	"fmt"
)

func Divide(a, b int) (err error) {
	err = validation(a, b)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Printf("a / b = %f", float32(a/b))
	fmt.Println()
	return nil

}
