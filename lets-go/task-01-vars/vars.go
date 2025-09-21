package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Constant declaration at the package level using
// short declaration (type inference)
const version = "1.0.0"
const vat = 0.21

// Variable declaration at the package level using
// explicit declaration
var name string = "Giovanni"

func main() {
	// Variable declaration with type inference
	age := 33

	// Variable declarations of other types using
	// both explicit and short forms
	var price float64 = 99.99
	active := true

	// Displaying the values of variables and constants
	fmt.Printf("Name: %s; Age: %d; Version: %s\n", name, age, version)
	fmt.Printf("Price: %.2f; Active: %t\n", price, active)

	// Using functions with multiple returns
	a, b := swap("Go", "Let's")
	fmt.Printf("Swap: %s, %s!\n", a, b)

	fmt.Printf("Double of your age: %d\n", doubleValue(age))
	fmt.Printf("Calculated price with VAT (Value Added Tax): %.2f\n", calcPrice(price))

	arr := []string{"Joel", "Ellie", "Tess", "Tommy", "Bill"}

	fmt.Printf("Array before swapArr: %s\n", strings.Join(arr, ", "))
	fmt.Printf("Array after swapArr: %s\n", strings.Join(swapArr(arr), ", "))

	data, err := json.MarshalIndent(arr, "", "  ")
	if err != nil {
		fmt.Println("Error generating JSON: ", err)
	}

	fmt.Println(string(data))
}

// Swap: switches the order of the provided values
// Example:
//
//	a, b := swap("Go", "Let's")
//	fmt.Printf("%s, %s\n", a, b)
//
// result:
//
//	Let's, Go
func swap(a, b string) (string, string) {
	return b, a
}

func doubleValue(x int) int {
	return x * 2
}

func calcPrice(p float64) float64 {
	return p + (p * vat)
}

func swapArr(a []string) []string {
	n := len(a)

	// Border Case
	if n == 0 || n == 1 {
		return a
	}

	b := make([]string, n)

	for i := range n {
		b[i] = a[n-1-i]
	}

	return b
}
