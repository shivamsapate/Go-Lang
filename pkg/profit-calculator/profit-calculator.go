package main

import (
	"fmt"
)

func main() {
	var revenue float64;
	var expenses float64;
	var taxRate float64; 

	printOutText("revenue amount : ")
	fmt.Scan(&revenue)

	printOutText("expenses amount : ")
	fmt.Scan(&expenses)
	
	printOutText("tax rate : ")
	fmt.Scan(&taxRate)

	 // earning before tax
	ebt, profit := calculateValues(revenue, expenses, taxRate)
	
	ratio := ebt / profit;

	printCalculate(ebt, profit, ratio)
}

func calculateValues(revenue float64, expenses float64, taxRate float64) (ebt float64,profit float64){
	ebt = revenue - expenses;
	profit = ebt * (1 - taxRate/100)
	return ebt, profit
}

func printOutText(text string){
	fmt.Print(text)
}

func printCalculate(ebt float64, profit float64, ratio float64) {
	fmt.Printf("Earning before tax: %f \nProfit : %f\nEBT to Profit ratio : %f", ebt, profit, ratio)
}