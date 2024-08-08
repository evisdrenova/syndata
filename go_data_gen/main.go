package main

import (
	"context"
	"fmt"
	"math"
	"math/rand"
	"strconv"

	_ "github.com/warpstreamlabs/bento/public/components/io"
	_ "github.com/warpstreamlabs/bento/public/components/pure/extended"

	_ "github.com/warpstreamlabs/bento/public/components/sql"
	"github.com/warpstreamlabs/bento/public/service"
)

func main() {
	service.RunCLI(context.Background())
}

// Generates a random int and returns it as a
func generateIntegerWithLength(length int) string {
	if length <= 0 {
		fmt.Println("Length must be a positive integer.")
		return "0" // Return a string "0" because the return type is string
	}

	// The lowest number with 'length' digits (e.g., length 3 -> 100).
	min := int64(math.Pow10(length - 1))

	// The highest number with 'length' digits (e.g., length 3 -> 999).
	max := int64(math.Pow10(length) - 1)

	// Generate a random number in the range [min, max] and convert it to string
	randomNum := rand.Int63n(max-min+1) + min
	return strconv.FormatInt(randomNum, 10)
}
