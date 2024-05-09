package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	content, err := os.ReadFile("standard.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	file := strings.Split(string(os.Args[1]), "\\n")

	for _, word := range file {
		if word == " " {
			continue
		}
		for char := 0; char < 7; char++ {
			for _, c := range word {
				formule := int(c) - 32
				formule1 := (formule + 1) + (formule * 8)
				if c < 32 || c > 126 {
					fmt.Println("Put a valid Data")
					return
				}
				fmt.Print(lines[formule1+char])
			}
			fmt.Println()
		}
	}

}
