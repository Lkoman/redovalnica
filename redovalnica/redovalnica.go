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

// Calculates the povprecje of all the ocene in the redovalnica za danega študenta (z dano vpisno številko), če je št. ocen >= stOcen
// returns povprecno oceno, izpise informacije
// PRIVATNA FUNKCIJA
func povprecje(studenti map[string]Student, vpisnaStevilka string, stOcen int) float64 {
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

	if st_ocen < stOcen {
		fmt.Println("Študent/ka ima premalo ocen in zaradi tega NI OPRAVIL/A predmeta.")
		fmt.Println("Št. potrebnih ocen:", stOcen)
		fmt.Println("Št. ocen študenta:", st_ocen)
		return 0.0
	}

	pov = float64(sum_ocen) / float64(st_ocen)
	if pov < 6.0 {
		fmt.Println("Študent/ka ima povprečje ocen", pov, "in zaradi tega predmeta NI OPRAVIL/A.")
		return 0.0
	}

	fmt.Println("Študent/ka", studenti[vpisnaStevilka].Ime, "ima povprečje ocen", pov, "in št. ocen", st_ocen, "in je predmet OPRAVIL/A.")
	return pov
}

// Public funtion, ki pokliče privatno funkcijo povprecje()
func IzpisiKoncniUspeh(studenti map[string]Student, vpisnaStevilka string, stOcen int) float64 {
	return povprecje(studenti, vpisnaStevilka, stOcen)
}

// Danemu študentu (z dano vpisno številko) v redovalnico apppend-amo dano oceno, če je ocena znotraj intervala [minOcena, maxOcena]
// retuns celotno redovalnico
func DodajOceno(studenti map[string]Student, vpisnaStevilka string, ocena int, minOcena int, maxOcena int) map[string]Student {
	if ocena < minOcena || ocena > maxOcena {
		fmt.Println("Na kakem faksu imajo ocene pod", minOcena, "ali nad", maxOcena, "?")
		return studenti
	}

	student, ok := studenti[vpisnaStevilka]
	if ok {
		student.Ocene = append(student.Ocene, ocena)
		studenti[vpisnaStevilka] = student
	} else {
		fmt.Println("Ta student/ka ne obstaja!")
		os.Exit(1)
	}
	return studenti
}
