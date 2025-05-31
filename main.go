package main

import (
	"fmt"
	"math"
)

func berechneAB(breiteCm, hoeheCm, schnurdickeMm float64) float64 {
	d := schnurdickeMm / 10.0 // mm → cm

	aMax := breiteCm / 2.0
	bMax := hoeheCm / 2.0

	n := math.Floor(aMax / d)

	// Dynamischer Korrekturfaktor je nach Form
	alpha := 0.75 * (hoeheCm / breiteCm)

	a := aMax - alpha*n*d
	b := a * (bMax / aMax)

	if a <= 0 || b <= 0 {
		return 0
	}

	c := math.Sqrt(a*a - b*b)
	return 2 * c
}

func main() {
	var breite, hoehe, schnurdicke float64

	fmt.Print("Schnurdicke in mm: ")
	fmt.Scanln(&schnurdicke)

	fmt.Print("Breite des Ovals in cm: ")
	fmt.Scanln(&breite)

	fmt.Print("Höhe des Ovals in cm: ")
	fmt.Scanln(&hoehe)

	ab := berechneAB(breite, hoehe, schnurdicke)
	fmt.Printf("AB-Abstand bei Spiral-Oval %.1f×%.1f cm: %.2f cm\n", breite, hoehe, ab)
}
