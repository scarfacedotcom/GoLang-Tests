package main

import "fmt"

func totalPrice(nights, rate, cityTax uint) uint {
	return nights*rate + cityTax
}
func main() {
	price := totalPrice(5, 10, 2)
	fmt.Println(price)
}
