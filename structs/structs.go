package structs

import "fmt"

type Student struct {
	name, matric_no string
	cgpa float64
}

func main() {
	stud1 := Student{"Blessing", "GLY/2018/030", 4.34}
	stud1.name = "Ewuola Blessing"
	stud2 := Student{cgpa: 3.56, name: "Ocan", matric_no: "GLY/1976/016"}
	stud_new := Student{}
	p_stud := &Student{"Tolu", "GLY/2018/027", 4.16}
	p_stud.cgpa = 4.99
	fmt.Println(stud1.name, stud2, stud_new, p_stud.cgpa)
}