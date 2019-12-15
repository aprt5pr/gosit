# Go Scale Interval Trainer (gosit)

Learn the major and minor scale intervals in all 12 keys.

[![Build Status](https://travis-ci.org/aprt5pr/gosit.svg?branch=master)](https://travis-ci.org/aprt5pr/gosit)

### Usage

```
package main

import (
    "fmt"
    "github.com/aprt5pr/gosit"
)

func main() {
    mynote := "C"
    if gosit.IsValidNoteName(mynote) {
      note, err := gosit.ResolveNoteName(mynote)
      fmt.Println(gosit.NewScale(note, "major"))
    } else {
      fmt.Println(mynote, "is not a valid note name")
    }
}

// &{C major [C D E F G A B C]} 
```

### License

This code is released under the MIT License. See [LICENSE](LICENSE).
