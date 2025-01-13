package main

import "fmt"

func main() {
	// Preset exchange rates (for example purposes)
	rates := map[string]float64{
		"USD_TO_EUR": 0.85,
		"USD_TO_GBP": 0.75,
		"EUR_TO_USD": 1.18,
		"EUR_TO_GBP": 0.88,
		"GBP_TO_USD": 1.33,
		"GBP_TO_EUR": 1.14,
	}

	fmt.Println("Currency Converter")
	fmt.Println("===================")

	// Input the source currency
	var fromCurrency, toCurrency string
	fmt.Print("Enter the source currency (USD, EUR, GBP): ")
	fmt.Scan(&fromCurrency)

	// Input the target currency
	fmt.Print("Enter the target currency (USD, EUR, GBP): ")
	fmt.Scan(&toCurrency)

	// Input the amount to convert
	var amount float64
	fmt.Print("Enter the amount to convert: ")
	fmt.Scan(&amount)

	// Build the conversion key
	conversionKey := fromCurrency + "_TO_" + toCurrency

	// Check if the conversion is valid
	rate, exists := rates[conversionKey]
	if !exists {
		fmt.Printf("Conversion from %s to %s is not supported.\n", fromCurrency, toCurrency)
		return
	}

	// Perform the conversion
	convertedAmount := amount * rate
	fmt.Printf("%.2f %s is equal to %.2f %s (Rate: %.2f).\n", amount, fromCurrency, convertedAmount, toCurrency, rate)
}
