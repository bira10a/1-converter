package main

import "fmt"

const USD = 1
const EUR = USD * 1.16
const RUB = USD * 72.14

func main() {

	EURfromRUM := RUB / EUR

	fmt.Println(EURfromRUM)

	sum := converter()
	fmt.Println(sum)

}

func converter() (sum float64) {

	var firstVal string
	var lastVal string
	var value float64

	fmt.Print("Выберите исходную валюту:")
	firstVal = getCurrency()

	value = getAmount()

	fmt.Print("Выберите вторую валюту:")
	lastVal = getCurrency()
	if firstVal == lastVal {
		fmt.Println("Выберите вторую валюту, отличную от первой:")
		lastVal = getCurrency()
	}

	sum = summConverter(firstVal, value, lastVal)

	return sum

}

func getCurrency() string {
	var value string

	for {
		fmt.Scan(&value)

		switch value {
		case "USD":
			fmt.Println("Вы выбрали доллары")
		case "EUR":
			fmt.Println("Вы выбрали евро")
		case "RUB":
			fmt.Println("Вы выбрали рубли")
		default:
			continue
		}

		break
	}

	return value
}

func getAmount() float64 {
	var value float64
	for {
		fmt.Print("Выберите число:")
		fmt.Scan(&value)

		if value == 0 || value < 0 {
			continue
		}
		fmt.Println("Вы выбрали", value)
		break
	}
	return value
}

func summConverter(firstVal string, value float64, lastVal string) float64 {
	var sum float64

	if firstVal == "USD" && lastVal == "EUR" {
		sum = (USD / 1.16) * value
	} else if firstVal == "USD" && lastVal == "RUB" {
		sum = (USD * 72.14) * value
	} else if firstVal == "EUR" && lastVal == "USD" {
		sum = (USD * 1.16) * value
	} else if firstVal == "EUR" && lastVal == "RUB" {
		sum = (USD * 72.14 * 1.16) * value
	} else if firstVal == "RUB" && lastVal == "USD" {
		sum = (USD / 72.14) * value
	} else if firstVal == "RUB" && lastVal == "EUR" {
		sum = (USD / (72.14 * 1.16)) * value
	}

	return sum
}
