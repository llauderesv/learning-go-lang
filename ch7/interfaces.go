package main

import (
	"fmt"
	"os"
)

type PaymentProcessor interface {
	ProcessPayment(amount float64) (bool, error)
}

type CreditCardProcessor struct {
	CardNumber string
	CVV        string
	ExpiryDate string
}

func (cc CreditCardProcessor) ProcessPayment(amount float64) (bool, error) {
	// Logic to process payment through a credit card
	fmt.Println("Processing credit card payment...")

	// Assume the payment is processed
	return true, nil
}

type PayPalProcessor struct {
	Email string
}

func (pp PayPalProcessor) ProcessPayment(amount float64) (bool, error) {
	fmt.Println("Processing PayPal payment...")

	return true, nil
}

func checkout(processor PaymentProcessor, amount float64) {
	success, err := processor.ProcessPayment(amount)
	if err != nil {
		fmt.Println("Failed to processed payment:", err)
		return
	}

	if success {
		fmt.Println("Payment processed successfully")
	} else {
		fmt.Println("Payment failed")
	}
}

func main() {
	var processor PaymentProcessor
	var paymentMethod string

	if len(os.Args[1:]) > 0 {
		paymentMethod = os.Args[1:][0]
	} else {
		panic("Error processing your payment")
	}

	switch paymentMethod {
	case "PayPal":
		processor = PayPalProcessor{Email: "llauderesv@gmail.com"}
	case "CreditCard":
		processor = CreditCardProcessor{CardNumber: "411121111111111", CVV: "123", ExpiryDate: "12/30"}
	default:
		fmt.Println("Payment not supported")
		return
	}

	checkout(processor, 99.99)
}
