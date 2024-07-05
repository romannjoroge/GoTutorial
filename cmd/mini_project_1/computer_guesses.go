package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func computerGuess(userGuess int8) error {
    var error error
    
    fmt.Printf("\n")
    error = printInBox("The Moment Of Truth")
    if error != nil {
        fmt.Println("Error Printing Box!", error.Error())
        return error
    }

    error = printParagraphInCenter([]string{
        "You can here beeping sounds coming from the machine as it carefully reads your thoughts",
        "Is this it! Is this the end for the great agent G!",
        "The machine beeps read and a sheet of paper with a number on it gets printed out.",
        "Its do or die!",
        "...",
    })
    if error != nil {
        fmt.Println("Error Printing Paragraph", error.Error())
        return error
    }

    var guess int8 = int8(rand.Intn(100))
    if guess <= 0 {
        error = printParagraphInCenter([] string {
            "The machine suddenly explodes giving you a perfect distraction to make an escape!",
            "You never got to see the number the machine produced but you're happy to be alive",
        })
        if error != nil {
            fmt.Println("Error Printing Paragraph", error.Error())
            return error
        }
        printLine(linelength)
        error = errors.New("Lucky")
        return error
    }

    if guess == userGuess {
        error = printParagraphInCenter([] string {
            "You refuse to believe your eyes but numbers don't lie!",
            "The machine succesfuly read your thoughts! Is this the end of freedom as you know it!",
            "The russians get the last laugh today!",
        })
        if error != nil {
            fmt.Println("Error Printing Paragraph", error.Error())
            return error
        }
        printLine(linelength)
        fmt.Printf("\nEnding 13: Unlucky day!\n")
        return error
    } 

    error = printParagraphInCenter([] string {
        "The russians look at you in shock when the machine prints a different number",
        "They wonder how this Kenyan spy beat the latest in Russian technology!",
        "You laugh as troops bust into the lab to come to your rescue!",
        "Agent G has done it again!",
    })
    if error != nil {
        fmt.Println("Error Printing Paragraph", error.Error())
        return error
    }
    printLine(linelength)
    fmt.Printf("\nEnding 2: Hero\n")
    return error
}
