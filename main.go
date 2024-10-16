package main

import (
	"fmt"
)

func GetTotalPayment(product string, amount float64, months int) float64 {

	var rate float64
	var maxMonths int

	switch product {
	case "telephone":
		maxMonths = 9
		rate = 0.03
	case "computer":
		maxMonths = 12
		rate = 0.04
	case "tv":
		maxMonths = 18
		rate = 0.05
	default:
		fmt.Println("Вы неправильно ввели название продукеа")
		return 0
	}

	if months < 3 || months > maxMonths {
		fmt.Printf("Неверный срок рассрочки для %s. Допустимые значения: 3-%d месяцев", product, maxMonths)
		return 0
	}

	if months > 3 {
		extraMonths := months - 3
		amount += amount * rate * float64(extraMonths)
	}

	return amount
}

func SendSMS(phone string, product string, amount float64, months int) {
	message := fmt.Sprintf("Спасибо за покупку %s в рассрочку на %d месяцев. Общая сумма: %.2f сомони.", product, months, amount)
	fmt.Printf("Отправка SMS на номер %s: %s\n", phone, message)
}

func main() {
	var (
		product string

		amount float64
		months int
		phone  string
	)

	fmt.Print("Введите название продукта (telephone, computer, tv): ")
	fmt.Scan(&product)

	fmt.Print("Введите сумму продукта: ")
	fmt.Scan(&amount)

	fmt.Print("На какйо срок вы хотите взять рассрочку: ")
	fmt.Scan(&months)

	fmt.Print("Введите номер телефона для SMS: ")
	fmt.Scan(&phone)

	totalAmount := GetTotalPayment(product, amount, months)

	if totalAmount > 0 {
		fmt.Printf("Общая сумма платежа для %s: %.2f сомони\n", product, totalAmount)
		SendSMS(phone, product, totalAmount, months)
	}
}
