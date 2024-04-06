package main

import (
	"fmt"
)

func main() {
	var revenue float64;
	var expenses float64;
	var taxRate float64; 

	fmt.Print("revenue amount : ")
	fmt.Scan(&revenue)

	fmt.Print("expenses amount : ")
	fmt.Scan(&expenses)
	
	fmt.Print("tax rate : ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses // earning before tax
	profit := ebt * (1 - taxRate/100)

	ratio := ebt / profit;

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)
	
}