package main

import (
	"math/rand"

	"github.com/warpstreamlabs/bento/public/bloblang"
)

// In Go, we use func init() to run code when the program initializes
func init() {
	spec := bloblang.NewPluginSpec()

	err := bloblang.RegisterFunctionV2("generate_card_number", spec, func(args *bloblang.ParsedParams) (bloblang.Function, error) {

		return func() (any, error) {
			res := generateCardNumber()

			return res, nil
		}, nil
	})

	if err != nil {
		panic(err)
	}
}

func generateCardNumber() string {
	visa := "4111"
	mc := "5111"

	// Generate the rest of the card number as a string with 12 digits
	rest := generateIntegerWithLength(12)

	// Randomly choose between Visa (1) and Mastercard (0)
	randCC := rand.Intn(2)

	// Concatenate the chosen prefix with the rest of the card number
	if randCC == 1 {
		return visa + rest
	}
	return mc + rest
}
