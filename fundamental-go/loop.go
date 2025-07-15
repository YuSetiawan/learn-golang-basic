package main

import "fmt"

// looping angka
func loop() {
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}
}

// loop angka dengan break & continue
func loopBreak() {
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}

		if i > 8 {
			break
		}

		fmt.Println("Angka", i)
	}
}
