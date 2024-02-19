package main

import (
	"errors"
	"fmt"
)

func pembagian(a, b int) (int, error) {
	if a == 0 {
		return 0, errors.New("Ups pembagian tidak bisa dilakukan")
	} else {
		return a / b, nil
	}
}

func main() {
	hasil, err := (pembagian(5, 2))
	if err == nil {
		fmt.Println(hasil)
	} else {
		fmt.Println("Error", err.Error())
	}
}
