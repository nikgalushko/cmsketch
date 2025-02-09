# CMSketch

Fastest count-min sketch implementation in Go, providing efficient estimation of item frequencies in a data stream.

## Benchmarks

`Apple M2`

```
Benchmark_W_2000_D_10
Benchmark_W_2000_D_10/update
Benchmark_W_2000_D_10/update-8         	45423646	        26.06 ns/op	       0 B/op	       0 allocs/op
Benchmark_W_2000_D_10/estimate
Benchmark_W_2000_D_10/estimate-8       	57425310	        21.84 ns/op	       0 B/op	       0 allocs/op
Benchmark_W_2000000_D_14
Benchmark_W_2000000_D_14/update
Benchmark_W_2000000_D_14/update-8      	 5442068	       211.2 ns/op	      41 B/op	       0 allocs/op
Benchmark_W_2000000_D_14/estimate
Benchmark_W_2000000_D_14/estimate-8    	 7010238	       176.0 ns/op	      31 B/op	       0 allocs/op
```

## Installation

To use this package, simply import it into your Go project:

```shell
go get github.com/nikgalushko/cmsketch
```

## Usage

```go
package main

import (
    "fmt"

    "github.com/nikgalushko/cmsketch"
)

func main() {
    // Create a new sketch with desired error and confidence values
    sketch := cmsketch.NewWithEstimates[string](0.01, 0.99)

    // Update the sketch with an item
    sketch.Update("item", 5)

    // Increment the count of an item
    sketch.Inc("item")

    // Estimate frequency of an item
    estimate := sketch.Estimate("item")
    fmt.Println("Estimated count:", estimate)
}
```