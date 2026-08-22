package main

import (
	"fmt"
	"strings"
)

var productPrices = map[string]float64{
	"TSHIRT": 20.00,
	"MUG":    12.50,
	"HAT":    18.00,
	"BOOK":   25.99,
}

func calculateItemPrice(itemCode string) (float64, bool) {
	basePrice, found := productPrices[itemCode]

	if !found {
		if strings.HasSuffix(itemCode, "_SALE") {
			originalItemCode := strings.TrimSuffix(itemCode, "_SALE")
			basePrice, found = productPrices[originalItemCode]
			if found {
				salePrice := basePrice * 0.90
				fmt.Printf(" - Item %s (Sale! Original Price: $%.2f, Sale Price: $%.2f\n",
					originalItemCode, basePrice, salePrice)
				return salePrice, true
			}
		}
		fmt.Printf(" - Item %s (Product not found)\n", itemCode)
		return 0.0, false
	}

	return basePrice, true
}

func main() {
	code := "TSHIRT"
	price, found := calculateItemPrice(code)
	if found {
		fmt.Printf("Price of %s is: %.2f", code, price)
	}

	orderItems := []string{
		"TSHIRT", "MUG_SALE", "HAT", "BOOK",
	}

	total := 0.0
	for _, orderItem := range orderItems {
		price, found = calculateItemPrice(orderItem)
		if found {
			total += price
		}
	}
	fmt.Printf("Total price: %.2f\n", total)

}
