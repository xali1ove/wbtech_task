//Реализовать паттерн «адаптер» на любом примере.
package main

import "fmt"

// PaymentProcessor интерфейс определяет метод для проведения платежа
type PaymentProcessor interface {
	Pay(amount float32) bool
}

// BankProcessor определяет метод для проведения платежа в банке
type BankProcessor interface {
	DoPayment(amount float32, cardNumber string) bool
}

// BankAdapter адаптер для использования BankProcessor вместо PaymentProcessor
type BankAdapter struct {
	bank BankProcessor
}

// Pay реализует метод Pay из интерфейса PaymentProcessor
func (ba *BankAdapter) Pay(amount float32) bool {
	return ba.bank.DoPayment(amount, "1234567890")
}

// Bank определяет тип для работы с банком
type Bank struct{}

// DoPayment реализует метод DoPayment из интерфейса BankProcessor
func (b *Bank) DoPayment(amount float32, cardNumber string) bool {
	fmt.Printf("Payment of $%0.2f using card %s processed by bank\n", amount, cardNumber)
	return true
}

func main() {
	bank := &Bank{}
	adapter := &BankAdapter{bank: bank}

	success := adapter.Pay(50.00)
	fmt.Println("Payment successful? ", success)
}
