package stripe

import (
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
)

func CreateCustomer(key string) *stripe.Customer {
	stripe.Key = key

	params := &stripe.CustomerParams{
		Description: stripe.String("My First Test Customer (created for API docs at https://www.stripe.com/docs/api)"),
	}
	c, _ := customer.New(params)
	return c
}
