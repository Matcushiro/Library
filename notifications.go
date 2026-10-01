package main

import "fmt"

type Notifier interface {
	Notify(message string)
}

type EmailNotifier struct {
	EmailAddress string
}

type SMSNotifier struct {
	PhoneNumber string
}

func (e EmailNotifier) Notify(message string) {
	fmt.Printf("Отправляю email на адрес %s: %s\n", e.EmailAddress, message)
}

func (s SMSNotifier) Notify(message string) {
	fmt.Printf("Отправляю SMS на номер %s: %s\n", s.PhoneNumber, message)
}
