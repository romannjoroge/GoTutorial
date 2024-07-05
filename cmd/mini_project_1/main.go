// A number guessing game
package main

import "fmt"


func main() {
    var error error
    // Tell the user about the game
    error = displayGameMenu()
    if error != nil {
        fmt.Println("Error Displaying Game Menu", error.Error())
    } else {
        // Get the user to input a number
        var number int8
        number, error = getNumberFromUser()
        if error != nil {
            switch error.Error() {
            case "Too many tries":
                fmt.Printf("\nEnding 1: You Cheeky Bastard!\n")
            default:
                fmt.Println("Error Getting Number From User", error.Error());
            }
        } else {
            // Computer guesses number
            error = computerGuess(number)
            if error != nil {
                switch error.Error() {
                case "Lucky":
                    fmt.Printf("\nEnding 3: Lucky Bastard!\n")
                default:
                    fmt.Println("Error Getting Computer's Guess")
                }
            }
        }
    }
}
