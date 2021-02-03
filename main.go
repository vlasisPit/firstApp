package main

import (
	"fmt"     //	Package fmt implements formatted I/O with functions analogous to C's printf and scanf.
	"strconv" //Package strconv implements conversions to and from string representations of basic data types.
)

//Declare variable on package level. Have to use full declaration syntax
var k float32 = 42

//You can declare multiple variables in a var block and you can delete the var on every variable
var (
	actorName string = "Vlasis Pitsios"
	companion string = "Gianna Tsaloufi"
	season    int    = 11
)

//1st scope: this is a lower case variable, so it is visible only inside the package name
var test string = "test variable"

//2nd scope: it is globally visible because the first the letter is uppercase
var Test string = "test variable 2"

const (
	catSpecialist = iota
	dogSpecialist
	snakeSpecialist
)

const (
	isAdmin = 1 << iota
	isHeadquarters
	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)

func main() {
	//3rd scope: block scoped. It is visible only inside the block
	i := 43            //this will infer an int by default
	var j float32 = 27 //you have more control over the variable
	fmt.Printf("%v, %T \n", j, j)
	fmt.Printf("%v, %T \n", i, i)
	fmt.Printf("%v, %T \n", k, k)
	fmt.Printf("%s %s %v \n", actorName, companion, season)

	//change variable type
	var a int = 42
	var b float32
	b = float32(a)
	fmt.Printf("%v, %T \n", b, b)

	//prints an the unicode representation of 42 which is an asterisk
	var changeVariableType string
	changeVariableType = string(a)
	fmt.Printf("%v, %T \n", changeVariableType, changeVariableType)
	changeVariableType = strconv.Itoa(a)
	fmt.Printf("%v, %T \n", changeVariableType, changeVariableType)

	//     PRIMITIVES
	var n bool = true
	var defaultBool bool //false
	fmt.Printf("%v, %T \n", n, n)
	fmt.Printf("%v, %T \n", defaultBool, defaultBool)

	var signedInteger int32 = 422
	fmt.Printf("%v, %T \n", signedInteger, signedInteger)

	var unsignedInteger uint16 = 422
	fmt.Printf("%v, %T \n", unsignedInteger, unsignedInteger)

	//		CONSTANTS
	//naming convention is valid here. That's why we dont use upper case letters
	//Immutable and can be swadowed
	const myConst int = 42
	//myConst = 43 compiler throws an error
	fmt.Printf("%v, %T \n", myConst, myConst)

	var specialistType int = 2
	fmt.Printf("%v, %T \n", specialistType == snakeSpecialist, specialistType == snakeSpecialist)

	//iota as a switch statement
	var roles byte = isAdmin | canSeeNorthAmerica | canSeeSouthAmerica
	fmt.Printf("%b, %T \n", roles, roles)
	fmt.Printf("Is Admin? %v \n", isAdmin & roles == isAdmin)
	fmt.Printf("Is canSeeEurope? %v \n", canSeeEurope & roles == canSeeEurope)

}
