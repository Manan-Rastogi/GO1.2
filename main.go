package main

import "fmt"

func main() {
	err1 := fmt.Errorf("Can not divide by %v.", 0)
	err2 := fmt.Errorf("Number 1 can not be greater than %v.", 10000)

	fmt.Printf("WrapErrors(err1, err2): %v\n", WrapErrors(err1, err2))
}