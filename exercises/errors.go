package main

import (
	"errors"
	"fmt"
)

type argError struct {
	arg int
	err string
}

func (a *argError) Error() string {
	return fmt.Sprintf("%d - %s", a.arg, a.err)
}

func New(errorString string) error {
	return &argError{0, errorString}
}

func func1(num int) (bool, error) {
	if num < 0 {
		return false, errors.New("can't work with that")
	}
	return true, nil
}

func func2(num int) (bool, error) {
	if num < 0 {
		return false, errors.New("can't work with -ve values")
	}
	return true, nil
}

func main() {
	fmt.Println(func1(1))
	fmt.Println(func1(-5))
	fmt.Println(func2(1))
	fmt.Println(func2(-5))
}
