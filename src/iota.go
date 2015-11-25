package main

import "fmt"

type ByteSize float64

// Go's iota identifier is used in const declarations to simplify definitions of incrementing numbers.
// Because it can be used in expressions, it provides a generality beyond that of simple enumerations.
// Wiki : https://github.com/golang/go/wiki/Iota
const (
    _           = iota // ignore first value by assigning to blank identifier
    KB ByteSize = 1 << (10 * iota)
    MB
    GB
    TB
    PB
    EB
    ZB
    YB
)

func main() {
    b := 1024;
    fmt.Println( ByteSize(b) );
}

func (b ByteSize) String() string {
    switch {
    case b >= YB:
        return fmt.Sprintf("%.2f YB", b/YB)
    case b >= ZB:
        return fmt.Sprintf("%.2f ZB", b/ZB)
    case b >= EB:
        return fmt.Sprintf("%.2f EB", b/EB)
    case b >= PB:
        return fmt.Sprintf("%.2f PB", b/PB)
    case b >= TB:
        return fmt.Sprintf("%.2f TB", b/TB)
    case b >= GB:
        return fmt.Sprintf("%.2f GB", b/GB)
    case b >= MB:
        return fmt.Sprintf("%.2f MB", b/MB)
    case b >= KB:
        return fmt.Sprintf("%.2f KB", b/KB)
    }
    return fmt.Sprintf("%.2fB", b)
}


