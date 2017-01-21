# Vaccinafy

This is a package that returns a date about when to apply a vaccine and its reference, 
also it returns true or false if you pass a date parameter.

It returns the result in Hash format.

# Installation

    go get github.com/mperezlamadrid/go-baby-vaccines

# Usage

Example:

```go
package main

import (
    "fmt"
    "github.com/mperezlamadrid/go-baby-vaccines"
)

func main(){
    dob := time.Date(2016, 11, 1, 0, 0, 0, 0, time.UTC)
	reqDate := time.Date(2017, 1, 12, 0, 0, 0, 0, time.UTC)
	result := vaccinefy.HasVaccineApplication(dob, reqDate)

    fmt.Println(result)
    // print: true
}
```

