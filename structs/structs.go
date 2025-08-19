package main

import "fmt"

type Person struct {
	name string
	age int
	job string
	salary int
}

func (p Person) getDesc() (desc string) {
	desc = fmt.Sprintf("Name: %s | Age: %d | Job: %s | Monthly salary: $%d", p.name, p.age, p.job, p.salary)
	return
}

func main() {
	var person1 = Person{name: "Tiskae", age: 23, job: "Software Developer", salary: 2300}
	person2 := Person{"Tiwu", 23, "Backend Engineer", 1500}

	fmt.Println(person1.getDesc(), person2.getDesc())
}
