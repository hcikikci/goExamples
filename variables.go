package main

import (
	"fmt"
	"reflect"
)

/*
uint8		Unsigned 8-bit integers (0 to 255)
uint16		Unsigned 16-bit integers (0 to 65535)
uint32		Unsigned 32-bit integers (0 to 4294967295)
uint64		Unsigned 64-bit integers (0 to 18446744073709551615)

int8		Signed 8-bit integers (-128 to 127)
int16		Signed 16-bit integers (-32768 to 32767)
int32		Signed 32-bit integers (-2147483648 to 2147483647)
int64		Signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

int			Both int and uint contain same size, either 32 or 64 bit.
uint		Both int and uint contain same size, either 32 or 64 bit.
rune		It is a synonym of int32 and also represent Unicode code points.
byte		It is a synonym of uint8.
uintptr		It is an unsigned integer type. Its width is not defined, but its can hold all the bits of a pointer value.

float32		32-bit floating-point numbers
float64		64-bit floating-point numbers

complex64	Complex numbers with float32 real and imaginary parts
complex128	Complex numbers with float64 real and imaginary parts

string		A string represents a sequence of characters. Strings are immutable; once they are created they can’t be modified.

bool		A boolean can hold one of two possible values, either true or false.

array		Arrays are list with static capacity. They can’t change their capacity after the declaration.
slice		Slices can change their capacity dynamically even after their declaration. Under the hood a slice references an array.
			If the array changes, so does the slice.

structs		In Go, a struct is a more complex type that can contain custom fields.
			It’s similar to an object in JavaScript or a dictionary in Python.
			The fields of a struct can be accessed with a dot .

pointers	Contain the memory address of the variable they are based on. Pointers used *.
*/

func main() {
	// Variable names are case-sensitive (age, Age and AGE are three different variables)
	// Declaring (Creating) Variables
	// Declaration with var keyword -> var variableName type = value (Can be used inside and outside of functions)
	var myName string = "Halit"
	fmt.Println("My name is " + myName)
	// Short Declaration with := sign -> variableName := value (Can only be used inside functions)
	isMarried := true
	fmt.Println("Am i married ?: ", isMarried)

	// Declaration without initial value
	var myName2 string
	fmt.Println(myName2) // Zero value -> ""
	myName2 = "Halitcan"
	fmt.Println(myName2)

	var isPhpDead bool
	fmt.Println("is php dead ?", isPhpDead) // Zero value -> false
	isPhpDead = true
	fmt.Println("is php dead ??", isPhpDead)

	var age int
	fmt.Println(age) // Zero value -> 0
	age = 23
	fmt.Println(age)

	/*
		The zero value is: 0 for numeric types,
		false for the boolean type,
		and. "" (the empty string) for strings.
	*/

	var array [5]int
	fmt.Println("empty array", array)
	array[1] = 1
	array[2] = 2
	fmt.Println("all array after set", array)
	fmt.Println("only value of index 1 in array -> ", array[1])
	fmt.Println("length of array", len(array))

	array2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array2)

	// Multiple Variable Declaration
	var a, b, c, d int = 1, 3, 5, 7
	fmt.Println(a, b, c, d)
	var (
		firstName string = "Halitcan"
		lastName  string = "Çıkıkçı"
		height    int
		age2      int
	)
	age2 = 23
	fmt.Println("First Name: " + firstName)
	fmt.Println("Last Name: " + lastName)
	fmt.Println("Age: ", age2)
	fmt.Println("Height Type: ", reflect.TypeOf(height))

}
