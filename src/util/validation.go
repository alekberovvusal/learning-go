package util

import (
	"errors"
	"fmt"
)

func validation(a, b int) error {

	if (a > 0) && (b > 0) {

		return nil

	}

	fmt.Printf(" a = %d ; b = %d ", a, b)

	return errors.New("ERROR: Numbers must be positive!")

}
