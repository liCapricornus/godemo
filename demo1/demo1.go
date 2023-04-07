package main

import "fmt"

func demo1() {
    var decimal int
    fmt.Print("Enter a decimal number: ")
    fmt.Scanln(&decimal)

    // Convert decimal to binary
    binary := make([]int, 0)
    for decimal > 0 {
        remainder := decimal % 2
        binary = append([]int{remainder}, binary...)
        decimal = decimal / 2
    }

    // Print the binary number
    fmt.Println("Binary number:", binary)
}