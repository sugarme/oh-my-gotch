package main

import (
	"fmt"

	"github.com/sugarme/gotch"
	"github.com/sugarme/gotch/ts"
)

func main() {
	device := gotch.CudaIfAvailable()
	x := ts.MustRandn([]int64{3, 4, 5}, gotch.Float, device)

	fmt.Printf("%i", x)
}
