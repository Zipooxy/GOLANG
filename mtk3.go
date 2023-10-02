package main

import "fmt"

func main() {
	var x int

	fmt.Print("masukan bilangan bulat positif kurang dari atau sama dengan (<= 999): ")
	fmt.Scan(&x)

	if x <= 999 {
		d3 := x % 10
		x = x / 10
		d2 := x % 10
		x = x / 10
		d1 := x 

		fmt.Println( d1, d2, d3 )
		 } else {
			fmt.Println(" input tidak valid. bilangan harus kurang dari atau sama dengan 999.")
		 }
	} 
