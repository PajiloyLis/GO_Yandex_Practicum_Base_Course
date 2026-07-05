package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Person struct {
	Name        string    `json:"Имя"`
	Email       string    `json:"Почта"`
	DateOfBirth time.Time `json:"-"`
}

func NewPerson(name, email string, dobYear, dobMonth, dobDay int) Person {
	return Person{
		Name:        name,
		Email:       email,
		DateOfBirth: time.Date(dobYear, time.Month(dobMonth), dobDay, 0, 0, 0, 0, time.UTC),
	}
}

func main() {
	p := NewPerson("Вася", "vasya@example.com", 1998, 1, 6)
	data, _ := json.Marshal(p)
	fmt.Println(string(data))
}
