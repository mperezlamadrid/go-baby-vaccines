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
    "time"
    
    "github.com/mperezlamadrid/go-baby-vaccines"
)

func main(){
    DOB := time.Date(2016, 11, 1, 0, 0, 0, 0, time.UTC)
	reqDate := time.Date(2017, 1, 12, 0, 0, 0, 0, time.UTC)

	result := vaccinefy.HasVaccinesToApply(DOB, reqDate)
    fmt.Println(result)
    // print: true

    vaccines := vaccinefy.GetVaccinesReference(DOB, 7)
    fmt.Printf("%+v\n", vaccines)
    // print: {Date:2017-06-01 00:00:00 +0000 UTC References:[{Name:Influenza Dose:2}]}
}
```

