package main

import (
	"fmt"
	"io"
	"log"
	"math/rand/v2"
	"os"
)

func newLine(writer io.Writer) {
	fmt.Fprintln(writer, "")
}

func writeDiagram(writer io.Writer) {
	fmt.Fprintln(writer, `<svg width="734" height="400" xmlns="http://www.w3.org/2000/svg">
    <circle cx="0" cy="400" r="400" fill="rgba(55, 55, 55, .1)" stroke="black" stroke-width="2"/>
    <rect x="0" y="0" width="400" height="400" fill="none" stroke="black" stroke-width="2"/>`)
	newLine(writer)

	numberIn := 0
	for _ = range 400 {
		x := rand.IntN(400)
		y := rand.IntN(400)
		color := "black"

		if x*x+(400-y)*(400-y) < 160_000 {
			numberIn++
			color = "blue"
		}

		fmt.Fprintf(writer, `    <circle cx="%d" cy="%d" r="4" fill="%s"/>`, x, y, color)
		newLine(writer)
	}

	newLine(writer)

	for y := range 20 {
		for x := range 20 {
			color := "black"
			if numberIn > 0 {
				numberIn--
				color = "blue"
			}

			fmt.Fprintf(writer, `    <circle cx="%d" cy="%d" r="4" fill="%s"/>`, 500+12*x, 314-12*y, color)
			newLine(writer)
		}
	}

	fmt.Fprintln(writer, "</svg>")
}

func main() {
	file, err := os.OpenFile("readme_images/monte-carlo.svg", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	writeDiagram(file)
}
