package main

import "fmt"

const USD = 1

func main() {
	const EUR = USD * 1.16
	const RUB = USD * 72.14

	EURfromRUM := RUB / EUR

	fmt.Println(EURfromRUM)

	value := userValue()
	fmt.Println(value)

}

func userValue() float64 {
	var value float64
	fmt.Scan(&value)
	return value
}

func converter(value, evro, rub float64) float64 {
	var USD = value
	var EUR = USD * evro
	var RUB = USD * rub

	var converterEURtoRub = RUB / EUR
	return converterEURtoRub
}
