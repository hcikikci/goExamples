package main

import "fmt"

/*
	The following verbs can be used with all data types:

	Verb	Description
	%v	Prints the value in the default format
	%#v	Prints the value in Go-syntax format
	%T	Prints the type of the value
	%%	Prints the % sign
*/

func main() {
	var i, j string = "Hello", "World"
	fmt.Print(i)
	fmt.Print(j, "\n")
	fmt.Print(i, "\n")
	fmt.Print(j, "\n")
	var x, y = 10, 20

	fmt.Print(x, y, "\n")

	var text string = "Lorem ipsum"
	var number int8 = 15
	fmt.Printf("text value: %v and type: %T\n", text, text)
	fmt.Printf("number value: %v and type: %T\n", number, number)

	var txt = "Hello World!"

	fmt.Printf("%v\n", number)
	fmt.Printf("%#v\n", number)
	fmt.Printf("%v%%\n", number)
	fmt.Printf("%T\n", number)

	fmt.Printf("%v\n", txt)
	fmt.Printf("%#v\n", txt)
	fmt.Printf("%T\n", txt)

	/*
		The following verbs can be used with the integer data type:
			Verb	Description
		%b	Base 2
		%d	Base 10
		%+d	Base 10 and always show sign
		%o	Base 8
		%O	Base 8, with leading 0o
		%x	Base 16, lowercase
		%X	Base 16, uppercase
		%#x	Base 16, with leading 0x
		%4d	Pad with spaces (width 4, right justified)
		%-4d	Pad with spaces (width 4, left justified)
		%04d	Pad with zeroes (width 4
	*/

	fmt.Printf("%b\n", number)
	fmt.Printf("%d\n", number)
	fmt.Printf("%+d\n", number)
	fmt.Printf("%o\n", number)
	fmt.Printf("%O\n", number)
	fmt.Printf("%x\n", number)
	fmt.Printf("%X\n", number)
	fmt.Printf("%#x\n", number)
	fmt.Printf("%4d\n", number)
	fmt.Printf("%-4d\n", number)
	fmt.Printf("%04d\n", number)

	/*
		The following verbs can be used with the string data type:

		Verb	Description
		%s	Prints the value as plain string
		%q	Prints the value as a double-quoted string
		%8s	Prints the value as plain string (width 8, right justified)
		%-8s	Prints the value as plain string (width 8, left justified)
		%x	Prints the value as hex dump of byte values
		% x	Prints the value as hex dump with spaces
	*/

	fmt.Printf("%s\n", txt)
	fmt.Printf("%q\n", txt)
	fmt.Printf("%24s\n", txt)
	fmt.Printf("%-8s\n", txt)
	fmt.Printf("%x\n", txt)
	fmt.Printf("% x\n", txt)

	/*
		Boolean Formatting Verbs
		The following verb can be used with the boolean data type:

		Verb	Description
		%t	Value of the boolean operator in true or false format (same as using %v)
	*/
	var t = true
	var f = false

	fmt.Printf("%t\n", t)
	fmt.Printf("%v\n", f)

	/*
		Float Formatting Verbs
		The following verbs can be used with the float data type:

		Verb	Description
		%e	Scientific notation with 'e' as exponent
		%f	Decimal point, no exponent
		%.2f	Default width, precision 2
		%6.2f	Width 6, precision 2
		%g	Exponent as needed, only necessary digits
	*/

	var float = 3.141

	fmt.Printf("%e\n", float)
	fmt.Printf("%f\n", float)
	fmt.Printf("%.2f\n", float)
	fmt.Printf("%6.2f\n", float)
	fmt.Printf("%g\n", float)
}
