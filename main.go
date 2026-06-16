package main

import "fmt"

func main() {
	const USD = 1
	const EUR = USD * 1.16
	const RUB = USD * 72.14

	EURfromRUM := RUB / EUR

	fmt.Println(EURfromRUM)
}
