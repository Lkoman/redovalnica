package redovalnica

import (
	"fmt"
	"os"
)

// Student structure
type Student struct {
	Ime     string
	Priimek string
	Ocene   []int
}

// Prints data from redovalnica for each student in the redovalnica
// brez returna, samo izpiše redovalnico
func IzpisVsehOcen(studenti map[string]Student) {
	fmt.Println("REDOVALNICA:")

	for key, value := range studenti {
		fmt.Println(key, "-", value.Ime, value.Priimek, ":", value.Ocene)
	}
	fmt.Println()
}

// Calculates the povprecje of all the ocene in the redovalnica za danega študenta (z dano vpisno številko)
// returns povprečje, brez izpisov (izpisi se naredijo kjer se pokliče funkcija)
// PRIVATNA FUNKCIJA
func povprecje(studenti map[string]Student, vpisnaStevilka string) float64 {
	student, ok := studenti[vpisnaStevilka]

	if !ok {
		return 0.0
	}

	var pov float64
	sum_ocen := 0
	st_ocen := len(student.Ocene)

	i := 0
	for i < st_ocen {
		sum_ocen += student.Ocene[i]
		i++
	}

	pov = float64(sum_ocen) / float64(st_ocen)
	if pov < 6.0 {
		return 0.0
	}
	return pov
}

// Public funtion, ki pokliče privatno funkcijo povprecje()
func IzpisiKoncniUspeh(studenti map[string]Student, vpisnaStevilka string) float64 {
	return povprecje(studenti, vpisnaStevilka)
}

// Danemu študentu (z dano vpisno številko) v redovalnico apppend-amo dano oceno
// retuns celotno redovalnico
func DodajOceno(studenti map[string]Student, vpisnaStevilka string, ocena int) map[string]Student {
	if ocena < 5 || ocena > 10 {
		fmt.Println("Na kakem faksu imajo ocene pod 5 ali nad 10?")
		os.Exit(1)
	}

	student, ok := studenti[vpisnaStevilka]
	if ok {
		student.Ocene = append(student.Ocene, ocena)
		studenti[vpisnaStevilka] = student
	} else {
		fmt.Println("Ta student ne obstaja!")
		os.Exit(1)
	}
	return studenti
}
