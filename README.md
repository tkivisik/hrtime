# hrtime

[![GoDoc](https://godoc.org/github.com/loov/hrtime?status.svg)](http://godoc.org/github.com/loov/hrtime)

**BETA QUALITY**

Package hrtime implements high-resolution timing functions and benchmarking utilities.

For example measuring `time.Sleep` precision on Mac and Windows.

## Example
```
package main

import (
    "fmt"
    "time"

    "github.com/loov/hrtime"
)

func main() {
    start := hrtime.Now()
    time.Sleep(1000 * time.Nanosecond)
    fmt.Println(hrtime.Since(start))

    const numberOfExperiments = 4096

    bench := hrtime.NewBenchmark(numberOfExperiments)
    for bench.Next() {
        time.Sleep(1000 * time.Nanosecond)
    }
    fmt.Println(bench.Histogram(10))
}
```

Output on Mac:

```
12µs
  avg 14.5µs;  min 2µs;  p50 12µs;  max 74µs;
  p90 22µs;  p99 44µs;  p999 69µs;  p9999 74µs;
        2µs [ 229] ██▌
       10µs [3239] ████████████████████████████████████████
       20µs [ 483] ██████
       30µs [  80] █
       40µs [  39] ▌
       50µs [  17] ▌
       60µs [   6]
       70µs [   3]
       80µs [   0]
       90µs [   0]
```

Output on Windows:

```
1.5155ms
  avg 1.49ms;  min 576µs;  p50 1.17ms;  max 2.47ms;
  p90 2.02ms;  p99 2.3ms;  p999 2.37ms;  p9999 2.47ms;
      577µs [   1]
      600µs [  57] █▌
      800µs [ 599] █████████████████
        1ms [1399] ████████████████████████████████████████
      1.2ms [  35] █
      1.4ms [   7]
      1.6ms [  91] ██▌
      1.8ms [ 995] ████████████████████████████
        2ms [ 778] ██████████████████████
      2.2ms [ 134] ███▌
```

_A full explanation why it outputs this is out of the scope of this document. However, all sleep instructions have a specified granularity and `time.Sleep` actual sleeping time is `requested time ± sleep granularity`. There are also other explanations to that behavior._