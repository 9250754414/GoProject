package main

import (
	"errors"
	"fmt"
	"unicode"
)

type IUser interface {
	GetAge() int
	GetName() string
	GetSurname() string
}

type Person struct {
	age     int
	name    string
	surname string
}

func main() {

	validatedPerson := validateWithRetry()

	fmt.Println("Валидация успешна")
	fmt.Println("Создан пользователь:")
	fmt.Printf("Имя: %s\nФамилия: %s\nВозраст: %d\n",
		validatedPerson.GetName(), validatedPerson.GetSurname(), validatedPerson.GetAge())
}

func NewPerson(age int, name string, surname string) Person {
	return Person{
		age:     age,
		name:    name,
		surname: surname,
	}
}

func validateWithRetry() Person {
	var age int
	var name, surname string

	// Валидация возраста
	for {
		fmt.Print("Введите возраст: ")
		fmt.Scan(&age)

		tempPerson := NewPerson(age, "", "")
		err := validate(tempPerson)
		if err != nil {
			fmt.Printf("Ошибка в возрасте: %v\n\n", err)
			continue
		}
		break
	}

	// Валидация имени
	for {
		fmt.Print("Введите имя: ")
		fmt.Scan(&name)

		tempPerson := NewPerson(0, name, "")
		err := validate(tempPerson)
		if err != nil {
			fmt.Printf("Ошибка в имени: %v\n\n", err)
			continue
		}
		break
	}

	for {
		fmt.Print("Введите фамилию: ")
		fmt.Scan(&surname)

		tempPerson := NewPerson(0, "", surname)
		err := validate(tempPerson)
		if err != nil {
			fmt.Printf("Ошибка в фамилии: %v\n\n", err)
			continue
		}
		break
	}

	return NewPerson(age, name, surname)
}

func validate(user IUser) error {
	age := user.GetAge()

	if age != 0 {
		if age < 1 {
			return errors.New("возраст не может быть меньше 1 года")
		}
		if age > 120 {
			return errors.New("возраст не может быть больше 120 лет")
		}
	}

	name := user.GetName()

	if name != "" && containsDigits(name) {
		return errors.New("имя не может содержать цифры")
	}

	surname := user.GetSurname()

	if surname != "" && containsDigits(surname) {
		return errors.New("фамилия не может содержать цифры")
	}

	return nil
}

func containsDigits(s string) bool {
	for _, char := range s {
		if unicode.IsDigit(char) {
			return true
		}
	}
	return false
}

func (p Person) GetAge() int {
	return p.age
}

func (p Person) GetName() string {
	return p.name
}

func (p Person) GetSurname() string {
	return p.surname
}
