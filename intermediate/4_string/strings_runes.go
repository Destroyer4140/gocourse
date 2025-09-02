package main

import (
	"fmt"
	"unicode/utf8"
)

/*
String is sequence of unsinged bites. It's immutable in golang.
*/
func main() {

	message := "Hello, \n\tGo!"
	message1 := "Hello, \r\nGo!"
	rawMessage := `Hello Go`

	fmt.Println("Message :-", message, "\nRawMessage :-", rawMessage)
	fmt.Println("Message1 :-", message1)
	fmt.Println("Length of message is :-", len(message))

	name1 := "Apple"
	name2 := "Appbe"

	fmt.Println(name1 < name2)

	//string iteration
	for i, char := range message {
		fmt.Printf("Character at index %d is %c\n", i, char)
	}

	fmt.Println("Rune in name1 ->", utf8.RuneCountInString(name1))

	var ch rune = 'a'
	fmt.Printf("Rune is %c\n", ch)
	cstr := string(ch)
	fmt.Println(cstr)
	fmt.Printf("Type of cstr is %T", cstr)

}

// --- General Formatting Verbs
/*
%v Prints the vlaue in the default format
%#v Prints the value in Go-Syntax format
%T Prints the type of the value
%% Prints the % sign
*/

// --- Integer Formatting Verbs
/*
%b Base 2
%d Base 10
%+d Base 10 and always show sign
%o Base 8
%0 Base 8, with leading 0o
%x Base 16, lowercase
%X Base 16, uppercase
%#x Base 16, with leading 0x
%4d Pad with spaces (width 4, right justified)
%-4d Pad with spaces (width 4, left justified)
%04d Paad with zeros (width 4)
*/
