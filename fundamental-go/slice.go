package main

import "fmt"

// choosing slice
func pickSlice() {
	var fruits = []string{"manggo", "grape", "banana", "melon"}
	var newFruits = fruits[0:2]

	fmt.Println(newFruits) // output: ["manggo", "grape"]
}

// fungsi len() untuk menghitung jumlah data pada variabel data
func countSlice() {
	var num = []string{"0", "1", "2", "3", "4"}
	fmt.Println(len(num)) // output: 5
}

// fungsi append() digunakan untuk menambahkan elemen pada slice
func addSlice() {
	var fruits = []string{"apple", "grape", "banana"}
	var cFruits = append(fruits, "papaya")

	fmt.Println(fruits)  // ["apple", "grape", "banana"]
	fmt.Println(cFruits) // ["apple", "grape", "banana", "papaya"]
}
