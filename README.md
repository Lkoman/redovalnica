*Modul redovalnica je preprosta Go knjižnica za delo s podatki o ocenah študentov.*

# Opis
**Struktura Student:** hrani podatke o študentih:
- Ime string
- Priimek string
- Ocene []int

**IzpisVsehOcen(map[string]Student):** izpiše vse študente in njihove ocene.

**DodajOceno(map[string]Student, string, int) map[string]Student:** doda oceno študentu z dano vpisno številko, če je ocena znotraj danega intervala.

**IzpisiKoncniUspeh(map[string]Student, string) float64:** vrne končno povprečje ocen študenta (uporablja interno privatno funkcijo).

**povprecje(...) float64:** privatna funkcija, ki izračuna povprečje ocen. Če je št. ocen manjše od danega minimalnega števila ocen ali pa je povprečje ocen manjše od 6.0, vrne 0.0.

# Struktura projketa
```
redovalnica/
│   go.mod
│   README.txt
│
├── cmd/
│   └── main.go
│
└── redovalnica/
    └── redovalnica.go
```

# Uporaba
## Kloniraj repozitorij
```
git clone https://github.com/lkoman/redovalnica.git
cd redovalnica
```
## Namesti CLI knjižnico
```
go get github.com/urfave/cli/v3
```
## Prevedi aplikacijo
```
go build ./cmd
```
## Zaženi aplikacijo
```
go run ./cmd/main.go --stOcen 3 --minOcena 5 --maxOcena 10 // če še ni buildan-a
./cmd --stOcen 3 --minOcena 5 --maxOcena 10 // po build-u
```
