/*
	Buatlah sebuah program dengan output sebagai berikut. Input bisa dinamis yang
	menghasilkan output yang berbeda-beda sesuai input yang dimasukan. Gunakan rumus
	A000124 of Sloane’s OEIS.
	Contoh:
	○ Input 7
	○ Output : 1-2-4-7-11-16-22
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Print("Input: ")
	_, err := fmt.Scanln(&n)
	if err != nil || n < 1 {
		fmt.Println("Invalid input, please enter an integer")
		return
	}

	// Generate n terms of the sequence
	result := SequenceA000124(n)
	fmt.Printf(strings.Join(result, " "))
	fmt.Println()
}

func SequenceA000124(n int) (result []string) {
	for i := 1; i <= n; i++ {
		result = append(result, fmt.Sprintf("%d", i*(i+1)/2+1))
	}
	return
}
