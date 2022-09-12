package main

import (
	"fmt"
	"ryuta06012/go_stripe_test/stripe"
)

//"sk_test_4eC39HqLyjWDarjtT1zdp7dc"
func main() {
	customer := stripe.CreateCustomer("sk_test_4eC39HqLyjWDarjtT1zdp7dc")
	fmt.Printf("customer.ID: %v\n", customer.ID)
}
