package main

import "fmt"

func main() {
    for e := range 64 {
        var i = 1 << e - 1
        var f = float64(i)
        var is = fmt.Sprintf("%d", i)
        var fs = fmt.Sprintf("%.f", f)
        if is != fs {
            fmt.Printf("%d\t%d\n\t%f\n", e, i, f)
        }
    }
}