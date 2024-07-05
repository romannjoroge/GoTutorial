package main

import (
	"errors"
	"fmt"
)

func getNumberFromUser() (int8, error) {
    var error error
    fmt.Println("\nTo make the challenge harder for you the russians have decided that the number must be between 1 and 100! Choose Wisely")
    
    var number int8
    var tries int8
    for {
        fmt.Printf("Enter your number: ")
        _, error = fmt.Scanln(&number)
        if error != nil {
            fmt.Println("Erorr Getting User Input!", error.Error())
            return 0, error
        }

        if number > 100 || number < 1 {
            if (tries == 3) {
                fmt.Printf("\nYour cheekiness was always your greatest strength and weakness Agent G but this time its gotten you killed! See you in the after life\n")
                return 0, errors.New("Too many tries")
            } else {
            fmt.Printf("\nYou fool! %v is not between 1 and 100. Don't test the russians patience! They've given you one more chance take it seriously!\n", number);
            tries++
            }
        } else {
            return number, error
        }
    }
}
