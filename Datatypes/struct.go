package main

import "fmt"

type student struct {
	students []studentid
}

type studentid struct {
	name  string
	grade int
}

func dostruct() {
	var s0 studentid
	s0.name = "Arga"
	s0.grade = 3
	fmt.Println(s0)

	Arya := studentid{"Arya", 4}
	Sindi := studentid{"Sindi", 3}

	studentids := []studentid{s0, Arya, Sindi}
	student := student{studentids}

	fmt.Println(student)
}
