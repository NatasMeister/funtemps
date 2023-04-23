package main

import (
	"flag"
	"fmt"
	"github.com/NatasMeister/funtemps/conv"
	"github.com/NatasMeister/funtemps/funfacts"
)

// Definerer flag-variablene i hoved-"scope"
var fahrenheit float64
var celsius float64
var kelvin float64
var out string
var ff string
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
	flag.Float64Var(&fahrenheit, "F", 0.0, "temperatur i grader fahrenheit")
	flag.Float64Var(&celsius, "C", 0.0, "temperatur i grader celsius")
	flag.Float64Var(&kelvin, "K", 0.0, "temperatur i grader kelvin")
	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")
	flag.StringVar(&ff, "ff", "sun", "\"fun-facts\" om sun - Solen, luna - Månen og terra - Jorden")
	flag.StringVar(&tempScale, "t", "C", "temperaturskala for visning av fun-facts. Kan være C - celsius, F - farhenheit, K - kelvin")

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

	if isFlagPassed("ff") && !isFlagPassed("t") {
		fmt.Println("-ff kan kun brukes med -t.")
		return
	}

	if isFlagPassed("F") {
		if out == "C" {
			celsius := conv.FahrenheitToCelsius(fahrenheit)
			fmt.Printf("%.2f°F is %.2f°C\n", fahrenheit, celsius)
		} else if out == "K" {
			kelvin := conv.FahrenheitToKelvin(fahrenheit)
			fmt.Printf("%.2f°F is %.2fK\n", fahrenheit, kelvin)
		} else {
			fmt.Println("Invalid temperature scale. Please use C, F, or K.")
		}
	}

	if isFlagPassed("C") {
		if out == "F" {
			fahrenheit := conv.CelsiusToFahrenheit(celsius)
			fmt.Printf("%.2f°C is %.2f°F\n", celsius, fahrenheit)
		} else if out == "K" {
			kelvin := conv.CelsiusToKelvin(celsius)
			fmt.Printf("%.2f°C is %.2fK\n", celsius, kelvin)
		} else {
			fmt.Println("Invalid temperature scale. Please use C, F, or K.")
		}
	}

	if isFlagPassed("K") {
		if out == "C" {
			celsius := conv.KelvinToCelsius(kelvin)
			fmt.Printf("%.2fK is %.2f°C\n", kelvin, celsius)
		} else if out == "F" {
			fahrenheit := conv.KelvinToFahrenheit(kelvin)
			fmt.Printf("%.2fK is %.2f°F\n", kelvin, fahrenheit)
		} else {
			fmt.Println("Invalid temperature scale. Please use C, F, or K.")
		}
	}

	if isFlagPassed("ff") && isFlagPassed("t") {
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