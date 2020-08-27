package util

import (
	"fmt"
)

func Subtract(a, b int) (err error) {
	err = validation(a, b)
	if err != nil {
		fmt.Println(err)
		return err
	}
	var z int = a - b
	fmt.Printf("a - b = %d", z)
	fmt.Println()
	return nil

}
