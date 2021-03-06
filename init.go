// Package hrtime implements High-Resolution Timing functions for benchmarking.
//
// The basic usage of this package looks like:
//
//     package main
//
//     import (
//         "fmt"
//         "github.com/loov/hrtime"
//     )
//
//     func main() {
//         const numberOfExperiments = 4096
//         bench := hrtime.NewBenchmark(numberOfExperiments)
//         for bench.Next() {
//             time.Sleep(10)
//         }
//         fmt.Println(bench.Histogram(10))
//     }
//
// To see more complex examples refer to the _example folder. (https://github.com/loov/hrtime/tree/master/_example)
package hrtime

const calibrationCalls = 1 << 10

func init() {
	calculateNanosOverhead()

	initCPU()
	{
		_, _, _, edx := cpuid(0x80000007, 0x0)
		rdtscpInvariant = edx&(1<<8) != 0
	}
	calculateTSCOverhead()
}
