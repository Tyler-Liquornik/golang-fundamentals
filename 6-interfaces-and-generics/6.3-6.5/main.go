package main

import "fmt"

type Printer interface {
	Print()
}

type Person struct {
	Name string
}

func (p *Person) Print() {
	fmt.Println("Name:", p.Name)
}

func callPrint(p Printer) {
	switch concreteValue := p.(type) {
	case *Person:
		if concreteValue == nil {
			fmt.Println("Cannot print: Person is nil")
			return
		}
		concreteValue.Print()
	default:
		p.Print()
	}
}

func main() {
	validPerson := &Person{Name: "Tyler"}
	callPrint(validPerson) // Works

	var nilPerson *Person = nil
	callPrint(nilPerson) // Works, now Handled safely
}
