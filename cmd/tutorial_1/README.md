## LESSON 1: Packages and Imports

So as I said before all of your Go code will be located in some package. To create a package all you need to do is to create a directory and in that directory define package name in the GO files. This declaration is done at the TOP of the file. MAKE SURE ITS AT THE TOP! EVEN IMPORTS SHOULD BE BELOW IT

Example of declaring package
```go
package <package-name>
```

Every module needs to have a main package so this folder acts as the main package!

To import something we use the import keyword followed by the name of the package in quotes, an example is:
```go
import "fmt" // this package contains functionality for things such as STD out
```

To compile and run your Go code you run the following command
```bash
go run <location-of-go-file>
```

To only compile run the command
```bash
go build <location-of-go-file>
```

## Some more details about packages, modules and imports
[Source Article](https://www.alexedwards.net/blog/an-introduction-to-packages-imports-and-modules)

So packages are just collections of related code that help one in the reuse of code. I'm guessing one imports code from packages throughout the project. Packages contain go files and each of these files indicate at the top what package it belongs to. This is giving the impression that a single folder could have files from multiple different packages? This is abit odd but will confirm later!.

Code that is in a package can access all the functionality that is defined in the package even if it is in a different Go file. So it seems that if you want your files to share functionality one way is to have all the files located in one package.

A misconception that I had was that a package needed a main.go file. This is not the case! The main function can be located in a file of any name but it must be in the package!

Interestingly to run the code in a package you don't just run the file that contains the main function and instead you have to build and run all of the files in the package. So to run the roman_test_1 code you would have to do:
```bash
go run cmd/roman_test_1/*.go
```

So the main package actually identifies the package that can be executed. So any code that is executable must be contained in the main package! If you try to execute code that is not contained in the main package you get this somewhat misleading error `package command-line-arguments is not a main package`.

When importing packages you have to give the full path to the package i.e when I used the Intn() function to generate a random number I had to specify the import as `import "math/rand"`. I couldn't just do `import "rand"`. This makes alot of sense. After importing a package we use the name of the package to access the functionality located in it i.e in the previous example I imported a package with the name rand and to access the Intn function I have to do **rand.Intn()**. This makes sense. 

Conveniently for the standard library all of the package names are the same as the last name in the import path. This is not a must but it helps alot.

Packages in a way act similarly to classes. They contain related functionality and can be used by other classes or functions etc. Packages in a similar way to classes also have the concept of public and private members. Private members are those that have a small lettter. They can be accessed by anything in the same package but not by things external to the package. Public members are defined using a capital letter and can be accessed by both items internal and external to the package. Examples are included below:

```go
func Public() {} // public member
func private() {} // private member
```

`This is as far as I'll go for now! There is alot more to read on modules, arranging your code to have multiple packages etc`

## Assesment of Go so far

I guess Go isn't too odd. I have barely any experience with it though and I know that there are bound to be some oddities ahead. I guess the oddest thing is the need to compile all files in a package and how visibility of items is handled. The compile part might be because I haven't used a compiled language in a while though
