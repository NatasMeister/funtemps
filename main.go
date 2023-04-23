package main

import (
	"flag"
	"fmt"
	"github.com/NatasMeister/funtemps/conv"
)

// Definerer flag-variablene i hoved-"scope"
var fahr float64
var celsius float64
var kelvin float64
var out string
var funfacts string
var tempScale string

// Bruker init (som anbefalt i dokumentasjonen) for å sikre at flagvariablene
// er initialisert.
func init() {

	/*
	   Her er eksempler på hvordan man implementerer parsing av flagg.
	   For eksempel, kommando
	       funtemps -F 0 -out C
	   skal returnere output: 0°F er -17.78°C
	*/

	// Definerer og initialiserer flagg-variablene
	flag.Float64Var(&fahr, "F", 0.0, "temperatur i grader fahrenheit")
	flag.Float64Var(&celsius, "C", 0.0, "temperatur i grader celsius")
	flag.Float64Var(&kelvin, "K", 0.0, "temperatur i grader kelvin")
	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")
	flag.StringVar(&funfacts, "funfacts", "sun", "\"fun-facts\" om sun - Solen, luna - Månen og terra - Jorden")
	flag.StringVar(&tempScale, "t", "C", "temperaturskala for visning av fun-facts. Kan være C - celsius, F - farhenheit, K - kelvin")
	// hvilken temperaturskala skal brukes når funfacts skal vises

}

func main() {

	flag.Parse()

	if isFlagPassed("F") && isFlagPassed("C") && isFlagPassed("K") {
		fmt.Println("-F, -C, -K kan ikke brukes samtidig.")
		return
	}

	if isFlagPassed("F") && isFlagPassed("C") && isFlagPassed("K") && isFlagPassed("out") {
		fmt.Println("-F, -C, -K kan ikke brukes samtidig med -out.")
		return
	}

	if isFlagPassed("funfacts") && !isFlagPassed("t") {
		fmt.Println("-funfacts kan kun brukes med -t.")
		return
	}

	if isFlagPassed("F") {
		celsius := conv.FahrenheitToCelsius(fahr)
		fmt.Printf("%.2f°F er %.2f°C\n", fahr, celsius)
	}

	if isFlagPassed("C") {
		fahrenheit := conv.CelsiusToFahrenheit(fahr)
		fmt.Printf("%.2f°C er %.2f°F\n", fahr, fahrenheit)
	}

	if isFlagPassed("K") {
		kelvin := conv.CelsiusToKelvin(fahr)
		fmt.Printf("%.2f°C er %.2fK\n", fahr, kelvin)
	}

	if isFlagPassed("funfacts") && isFlagPassed("t") {
		funFact := funfacts.GetFunFacts("luna")
		fmt.Println(funFact)
	}

	if isFlagPassed("out") && out != "" {
		fmt.Println("out er satt til", out)
	}

	fmt.Println("len(flag.Args())", len(flag.Args()))
	fmt.Println("flag.NFlag()", flag.NFlag())
}

func isFlagPassed(name string) bool {
    found := false
    flag.Visit(func(f *flag.Flag) {
        if f.Name == name {
            found = true
        }
    })
    return found
}