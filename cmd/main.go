package main

import (
	"fmt"

	"github.com/lkoman/redovalnica"
)

func main() {
	var studenti map[string]redovalnica.Student
	studenti = make(map[string]redovalnica.Student)

	var s1, s2 redovalnica.Student
	s1.Ime = "Laraana"
	s1.Priimek = "Koman"
	s1.Ocene = []int{9, 10, 7}

	s2 = redovalnica.Student{
		Ime:     "Janez",
		Priimek: "Novak",
		Ocene:   []int{5, 5, 5},
	}

	s3 := redovalnica.Student{
		Ime:     "John",
		Priimek: "Wick",
		Ocene:   []int{10, 10, 10},
	}

	studenti["63210152"] = s1
	studenti["63210043"] = s2
	studenti["63210067"] = s3

	redovalnica.IzpisVsehOcen(studenti)

	studenti = redovalnica.DodajOceno(studenti, "63210152", 9)

	redovalnica.IzpisVsehOcen(studenti)

	fmt.Println("Povprecje ocen za studenta", studenti["63210152"].Ime, ":", redovalnica.IzpisiKoncniUspeh(studenti, "63210152"))

}
