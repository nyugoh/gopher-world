package main

import "fmt"

func main() {
    var fahr float64
    var cels float64

    // Get fahrenhirts
    fmt.Print("Enter temp in Fahrenheits :")
    fmt.Scanf("%f", &fahr)

    cels = ((fahr - 32)*5/9)
    fmt.Print(fahr)
    fmt.Println(" in celsius = ", cels)
}
