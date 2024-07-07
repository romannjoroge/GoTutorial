### Strings in Go

So strings just behave really wierldy in Go, maybe wierd is not the best word but there you have to put alot more thought than I thought you'd need to to work with strings.

First indexing a string gives you the int8 representation of the string which is just odd and looping through a string does the same thing!. 

Casting the string to a list of runes doesn't solve the appove issues but atleast it deals with the odd issues of different characters in strings having different number of bytes because of how they are encoded.

Strings can be concatenated using the + operator but they are immutable so you can change the characters in a string i.e `mystring[0] = 'a'` will give an error. To do something like this you'd have to create another string.

You can use the string builder to build a string i.e maybe you want to create a student administration string that has the following format ~Reg No Name Course~. One way to do this would be to concatenate the regno, name and course or you can as well use string builder as shown below:

```go
var strBuilder strings.Builder
strBuilder.WriteString(regNo)
strBuilder.WriteString(name)
strBuilder.WriteString(course)

var studentString = strBuilder.String()
```
`Strings.Builder` can be imported from strings package
