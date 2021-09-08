package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) introduce() {
	fmt.Printf("Hello everyone my name is %s \n", p.Name)
}

type Student struct {
	*Person
	GPA float64
}

func main() {
	joe := Student{
		Person: &Person{"Jane Doe"},
		GPA:    3.5567,
	}

	joe.introduce()
}
