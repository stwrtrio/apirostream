package utils

import (
	"fmt"
)

// Menampilkan pilihan dan meminta input dari user
func SelectOption(prompt string, options []string) string {
	fmt.Println(prompt)
	for i, option := range options {
		fmt.Printf("%d. %s\n", i+1, option)
	}

	var choice int
	fmt.Print("Pilih nomor: ")
	fmt.Scanln(&choice)

	if choice < 1 || choice > len(options) {
		fmt.Println("Pilihan tidak valid. Menggunakan default.")
		return options[0] // Default: pilihan pertama
	}

	return options[choice-1]
}
