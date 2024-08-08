package main

import (
	"github.com/warpstreamlabs/bento/public/bloblang"
)

// In Go, we use func init() to run code when the program initializes
func init() {
	spec := bloblang.NewPluginSpec()

	err := bloblang.RegisterFunctionV2("generate_routing_number", spec, func(args *bloblang.ParsedParams) (bloblang.Function, error) {

		return func() (any, error) {
			res := generateIntegerWithLength(9)

			return res, nil
		}, nil
	})

	if err != nil {
		panic(err)
	}
}
