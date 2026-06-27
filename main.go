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

	for {
		fmt.Print("Выберите исходную валюту:")
		fmt.Scan(&firstVal)

		switch firstVal {
		case "USD":
			fmt.Println("Вы выбрали доллары")
		case "EUR":
			fmt.Println("Вы выбрали евро")
		case "RUB":
			fmt.Println("Вы выбрали рубли")
		default:
			fmt.Println("Выберите исходную валюту:")
			fmt.Scan(&firstVal)
		}

		break
	}

	for {
		fmt.Print("Выберите число:")
		fmt.Scan(&value)

		if value == 0 || value < 0 {
			fmt.Print("Выберите число pliz:")
			fmt.Scan(&value)
		}
		fmt.Println("Вы выбрали", value)
		break
	}

	for {
		fmt.Print("Выберите вторую валюту:")
		fmt.Scan(&lastVal)

		if firstVal == lastVal {
			fmt.Println("Выберите вторую валюту, отличную от первой:")
			fmt.Scan(&lastVal)
		}

		switch lastVal {
		case "USD":
			fmt.Println("Вы выбрали второй валютой доллары")
		case "EUR":
			fmt.Println("Вы выбрали второй валютой евро")
		case "RUB":
			fmt.Println("Вы выбрали второй валютой рубли")
		default:
			fmt.Println("Выберите вторую валюту:")
			fmt.Scan(&firstVal)
		}

		break
	}

	if firstVal == "USD" && lastVal == "EUR" {
		sum = (USD / 1.16) * value
		return
	} else if firstVal == "USD" && lastVal == "RUB" {
		sum = (USD * 72.14) * value
		return
	} else if firstVal == "EUR" && lastVal == "USD" {
		sum = (USD * 1.16) * value
		return
	} else if firstVal == "EUR" && lastVal == "RUB" {
		sum = (USD * 72.14 * 1.16) * value
		return
	} else if firstVal == "RUB" && lastVal == "USD" {
		sum = (USD / 72.14) * value
		return
	} else if firstVal == "RUB" && lastVal == "EUR" {
		sum = (USD / (72.14 * 1.16)) * value
		return
	}

	return

}
