# tcmp

> [WIP] A tool to compare types with generic(T) in Go.


## Getting Started

### Install

`go get https://github.com/tcmp`

### Usage

```Go
// main.go

package main

import (
    "github.com/i0Ek3/tcmp"
)

func main() {
    //...

    tcmp.Equal(tslice1, tslice2)
    tcmp.IsNil(tslice)

    tcmp.Equal(tmap1, tmap2)
    tcmp.IsNil(tmap)

    //...
}
```

## TODO

Hold on...
