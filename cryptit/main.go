package main

import (
    "fmt"
    "github.com/med3633/cryptit/encrypt"
    "github.com/med3633/cryptit/decrypt"
)

func main() {
    // Use the function from the encrypt package
    encryptedMessage := encrypt.Enc("hello world ")
    fmt.Printf(encryptedMessage)
    fmt.Printf("\n")
        // Use the function from the decrypt package
        decryptedMessage := decrypt.Dec("khoor#zruog")
        fmt.Printf(decryptedMessage)
}     
