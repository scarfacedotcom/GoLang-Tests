package prices

import "fmt"

func TotalPrice(nights, rate, cityTax uint) uint {
	return nights*rate + cityTax
}

func main() {
	price := TotalPrice(3, 10000, 132)
	fmt.Println(price)
}
