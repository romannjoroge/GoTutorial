package main

import (
	"errors"
	"fmt"
	"strings"
)

const linelength = 125

func displayGameMenu() error {
    var error error;
    error = printInBox("Russian Rulette")
    if(error != nil) {
        fmt.Println("Error Displaying Game Menu", error.Error())
        error = errors.New("Could Not Display Game Menu")
        return error
    }

    gameIntroduction := []string{
        "You're mission was to infiltrate a secret Russian lab and steal the plans for their newest weapon",
        "But alas things go awry and you find yourself hooked up to their experimental mind reading device.",
        "The russians are confident in their device, so confident that they give you a gamble of a life time.",
        "Try and write down a number without the machine guessing what it is correctly.",
        "Will you spy training best the machine or will this be your final mission?",
        "Good luck Agent G!",
    }
    error = printParagraphInCenter(gameIntroduction)
    if error != nil {
        fmt.Println("Error Printing Paragraph", error.Error())
        return error
    }

    printLine(linelength)
    return error
}


func printInBox(message string) error {
    var error error
    printLine(linelength)
    error = printStringInCenter(message)

    if error != nil {
        fmt.Println("Error Printing string in center", error.Error())
        error = errors.New("Could Not Print Box")
        return error
    }
    printLine(linelength)
    return error
}

func printLine(length int) {
    var line = strings.Repeat("-", length)
    fmt.Println(line);
}

func printStringInCenter(toPrint string) error{
    var error error
    var lenString = len(toPrint)
    if lenString >= linelength {
        error = errors.New("String is longer than line length")
        return error
    }

    var difference = linelength - lenString - 2
    if difference % 2 == 0 {
        var space = strings.Repeat(" ", difference / 2)
        fmt.Printf("|%v%v%v|\n", space, toPrint, space)
    } else {
        var firstSpace = strings.Repeat(" ", difference / 2)
        var lastSpace = strings.Repeat(" ", (difference / 2) + 1)
        fmt.Printf("|%v%v%v|\n", firstSpace, toPrint, lastSpace)
    }

    return error
}

func printParagraphInCenter(toPrint []string) error {
    var error error

    for _, str := range toPrint {
        error = printStringInCenter(str)
        if error != nil {
            fmt.Println("Error Printing Line in Sentence", str, error.Error())
            return error
        }
    }

    return error
}
