package main

import (
	"fmt"
	"math"
)

// berechneAB berechnet den AB-Abstand für ein Spiral-Oval
// breiteCm: Breite des Ovals in Zentimetern
// hoeheCm: Höhe des Ovals in Zentimetern
// schnurdickeMm: Dicke der verwendeten Schnur in Millimetern
// Rückgabewert: berechneter AB-Abstand in Zentimetern
func berechneAB(breiteCm, hoeheCm, schnurdickeMm float64) float64 {
	d := schnurdickeMm / 10.0 // Umrechnung Schnurdicke von mm in cm

	aMax := breiteCm / 2.0 // Halbe Breite (Halbachse a)
	bMax := hoeheCm / 2.0  // Halbe Höhe (Halbachse b)

	n := math.Floor(aMax / d) // Anzahl der Windungen (grob geschätzt)

	// Dynamischer Korrekturfaktor je nach Ovalform
	alpha := 0.75 * (hoeheCm / breiteCm)

	a := aMax - alpha*n*d  // Angepasste Halbachse a nach Windungen
	b := a * (bMax / aMax) // Angepasste Halbachse b proportional

	// Falls die Achsen negativ oder null werden, ist keine weitere Windung möglich
	if a <= 0 || b <= 0 {
		return 0 // Kein sinnvoller Wert möglich
	}
	c := math.Sqrt(a*a - b*b) // Brennpunktabstand für Ellipse
	return 2 * c              // AB-Abstand ist die doppelte Brennpunktdistanz
}

func main() {
	var breite, hoehe, schnurdicke float64

	// Benutzereingabe: Schnurdicke in mm
	fmt.Print("Schnurdicke in mm: ")
	fmt.Scanln(&schnurdicke)

	// Benutzereingabe: Breite des Ovals in cm
	fmt.Print("Breite des Ovals in cm: ")
	fmt.Scanln(&breite)

	// Benutzereingabe: Höhe des Ovals in cm
	fmt.Print("Höhe des Ovals in cm: ")
	fmt.Scanln(&hoehe)

	if breite <= 0 || hoehe <= 0 || schnurdicke <= 0 {
		fmt.Println("Breite, Höhe und Schnurdicke müssen größer als 0 sein.")
		return
	}

	// Berechnung und Ausgabe des AB-Abstands
	ab := berechneAB(breite, hoehe, schnurdicke)
	fmt.Printf("AB-Abstand bei Spiral-Oval %.1f×%.1f cm: %.2f cm\n", breite, hoehe, ab)
}
