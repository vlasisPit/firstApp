package main

import (
	"fmt" //	Package fmt implements formatted I/O with functions analogous to C's printf and scanf.
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
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

/*You need to use capital letters to all variables of the struct to be visible outside of the package !!!!!
No underscores on field names or struct names*/
type Doctor struct {
	Number     int
	ActorName  string
	Companions []string
}

/*
Tags to make some validations on the data `required max:"100"`
*/
type Animal struct {
	Name   string `required max : "100"`
	Origin string
}

type Bird struct {
	Animal   //composition or embedding
	SpeedKPH float32
	CanFly   bool
}

func main() {
	/*
		n Go, := is for declaration + assignment, whereas = is for assignment only.
		For example, var foo int = 10 is the same as foo := 10
	*/

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
	fmt.Printf("Is Admin? %v \n", isAdmin&roles == isAdmin)
	fmt.Printf("Is canSeeEurope? %v \n", canSeeEurope&roles == canSeeEurope)

	//   ARRAYS AND SLICES (two collection types)
	grades := [3]int{97, 85, 93} //or
	grades2 := [...]int{97, 85, 93}
	fmt.Printf("Grades: %v \n", grades)
	fmt.Printf("Grades: %v \n", grades2)

	var students [3]string
	fmt.Printf("Students: %v \n", students)
	students[0] = "Vlasis"
	students[1] = "Gianna"
	students[2] = "Tsal"
	fmt.Printf("Students: %v \n", students)
	fmt.Printf("Size of students: %v \n", len(students))

	//copies the whole array into a NEW array
	gradesCopy := grades
	gradesCopy[1] = 5
	fmt.Printf("Grades: %v \n", grades)
	fmt.Printf("Grades: %v \n", gradesCopy)

	//pointers. Pass a reference
	gradesPointer := &grades
	gradesPointer[1] = 5
	fmt.Printf("Grades: %v \n", grades)
	fmt.Printf("Grades: %v \n", gradesPointer)

	//initialize a slice. No dots into []
	gradesSlice := []int{97, 85, 93}
	fmt.Printf("Size of grades slice: %v \n", len(gradesSlice))     //size of slice
	fmt.Printf("Capacity of grades slice: %v \n", cap(gradesSlice)) //size of underline array

	//slices are reference types (add and remove elements from slice during their lifetime)
	gradesSliceCopy := gradesSlice
	gradesSliceCopy[1] = 5
	fmt.Printf("Grades: %v \n", grades)
	fmt.Printf("Grades: %v \n", gradesSliceCopy)

	sliceWithMake := make([]int, 3, 3)
	fmt.Printf("sliceWithMake: %v \n", sliceWithMake)
	fmt.Printf("Size of sliceWithMake slice: %v \n", len(sliceWithMake))     //size of slice
	fmt.Printf("Capacity of sliceWithMake slice: %v \n", cap(sliceWithMake)) //size of underline array
	sliceWithMake = append(sliceWithMake, 3)                                 //capacity increased
	fmt.Printf("sliceWithMake: %v \n", sliceWithMake)
	fmt.Printf("Size of sliceWithMake slice: %v \n", len(sliceWithMake))     //size of slice
	fmt.Printf("Capacity of sliceWithMake slice: %v \n", cap(sliceWithMake)) //size of underline array
	sliceWithMake = append(sliceWithMake, 5, 6, 7, 8)                        //capacity increased
	fmt.Printf("sliceWithMake: %v \n", sliceWithMake)
	fmt.Printf("Size of sliceWithMake slice: %v \n", len(sliceWithMake))     //size of slice
	fmt.Printf("Capacity of sliceWithMake slice: %v \n", cap(sliceWithMake)) //size of underline array

	//remove the element from index 2. Careful yoy work with references here
	testSlice := []int{1, 2, 3, 4, 5}
	testSliceResult := append(testSlice[:2], testSlice[3:]...)
	fmt.Printf("testSliceResult: %v \n", testSliceResult)

	//		MAPS
	statePopulations := map[string]int{
		"California": 39250018,
		"Texas":      27232432,
		"Florida":    20232432,
		"Ohio":       11632432,
	}
	fmt.Printf("statePopulations: %v \n", statePopulations)
	fmt.Printf("Ohio population: %v \n", statePopulations["Ohio"])
	statePopulations["Georgia"] = 10310371
	fmt.Printf("Georgia population: %v \n", statePopulations["Georgia"]) //add entry to a map
	delete(statePopulations, "Texas")                                    //Delete Texas
	fmt.Printf("Texas population: %v \n", statePopulations["Texas"])     //returns zero

	georgiaPop, ok := statePopulations["Georgia"]
	fmt.Println(georgiaPop, ok) //ok == true

	georgiaPop2, ok2 := statePopulations["Georgiasss"] //misspelling
	fmt.Println(georgiaPop2, ok2)                      //ok == false

	fmt.Printf("size %v\n", len(statePopulations))

	sp := statePopulations //pass by reference
	delete(sp, "Ohio")
	fmt.Printf("size %v\n", len(sp))

	//		STRUCT
	aDoctor := Doctor{
		Number:    3,
		ActorName: "John",
		Companions: []string{
			"Mike",
			"Jim",
		},
	}
	fmt.Printf("aDoctor %v\n", aDoctor)
	fmt.Printf("aDoctor name %v\n", aDoctor.ActorName)
	fmt.Printf("aDoctor companion 1 %v\n", aDoctor.Companions[1])

	//pass copy of the same data
	anotherDoctor := aDoctor
	anotherDoctor.ActorName = "Tim"
	fmt.Printf("aDoctor name %v\n", aDoctor.ActorName)
	fmt.Printf("anotherDoctor name %v\n", anotherDoctor.ActorName)
	fmt.Println(aDoctor)
	fmt.Println(anotherDoctor)

	//pass reference of the same data
	anotherDoctorRef := &aDoctor
	anotherDoctorRef.ActorName = "Tim"
	fmt.Printf("aDoctor name %v\n", aDoctor.ActorName)
	fmt.Printf("anotherDoctor name %v\n", anotherDoctorRef.ActorName)
	fmt.Println(aDoctor)
	fmt.Println(anotherDoctorRef)

	//GO does not support inheritance. GO does not support traditional object oriented principles. Uses composition instead
	birdInstance := Bird{}
	birdInstance.Name = "Emu"
	birdInstance.Origin = "Australia"
	birdInstance.SpeedKPH = 48
	birdInstance.CanFly = false
	fmt.Println(birdInstance)
	fmt.Println(birdInstance.Name)

	birdInstance2 := Bird{
		Animal:   Animal{Name: "Em2", Origin: "Europe"},
		CanFly:   true,
		SpeedKPH: 67,
	}
	fmt.Println(birdInstance2.Name)

	//validation library should read tag via reflection
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)

	//   IF AND SWITCH STATEMENTS
	if pop, ok := statePopulations["Florida"]; ok {
		fmt.Printf("Florida population: %v \n", pop)
	}

	number := 50
	guess := 60
	if number > guess || returnTrue() {
		fmt.Println("True")
	}

	//switch statement for data type
	var checkType interface{} = "dfas"
	switch checkType.(type) {
	case int:
		fmt.Println("checkType is an int")
	case float64:
		fmt.Println("checkType is an float64")
	case string:
		fmt.Println("checkType is an string")
	default:
		fmt.Println("checkType is another type")
	}

	// LOOPING
	// cnt is valid only inside for loop
	for cnt := 0; cnt < 5; cnt++ {
		fmt.Println(cnt)
	}

	//GO does not support while and do. Instead of while you can use
	cnt2 := 0
	for cnt2 < 7 {
		fmt.Println(cnt2)
		cnt2++
	}

	//Do-while loop
	cnt3 := 0
	for {
		fmt.Println(cnt3)
		cnt3++
		if cnt3 == 15 {
			break
		}
	}

	// works for slices and arrays
	collectionExample := []int{1, 2, 3, 7}
	for k, v := range collectionExample { //k,v key-value , index-value
		fmt.Println(k, v)
	}

	// works for maps
	for k, v := range statePopulations {
		fmt.Println(k, v)
	}

	//works for strings (unicode representation for a char in string)
	text := "check range keyword"
	for k, v := range text {
		fmt.Println(k, v)
		fmt.Println(k, string(v))
	}

	//works also with channels !!!!
	//you should use both k and v because this is mandatory from the language
	//if you need only the keys
	for k := range statePopulations {
		fmt.Println(k)
	}

	for _, v := range statePopulations {
		fmt.Println(v)
	}

	//  Control flow constructs: DEFER, PANIC and RECOVER
	/*
		DEFER: Delay execution to some future point in time
		PANIC: Fail fast on errors that shouldn’t occur during normal operation, or that we aren’t prepared to handle gracefully.
		RECOVER: Save the program when it starts to panic
	*/

	//Defer keyword to close the resources in an opposite order we opened them !!!!!!!!
	/*	fmt.Println("start")
		defer deferExample("middle")
		fmt.Println("end")*/

	/**
	They run on a LIFO order result: end middle start
	defer fmt.Println("start")
	defer deferExample("middle")
	defer fmt.Println("end")
	*/

	/*
		open and close a resource with defer
	*/
	runResourceRequest()

	//GO does not support exceptions. Use panic when the application can continue to function
	/*num1, num2 := 1, 0
	ans := num1 / num2	//program stops working. You can do the same with panic
	fmt.Println(ans)*/

	/*fmt.Println("start 1")
	defer fmt.Println("this was deferred 1")	//this will be executed before panicking
	panic("something bad happened")
	fmt.Println("end")*/

	fmt.Println("start")
	panicker()
	fmt.Println("end")

	// POINTERS
	passByValueExample()
	passByReferenceExample()
	pointersOnArrays()
	pointersOnStructs()
	pointersOnArraysAndSlices()

	//	FUNCTIONS
	greeting := "Hello"
	name := "Vlasis"
	sayGreetingByValue(greeting, name) //pass by value
	fmt.Println("Caller name " + name)

	//passing a reference is more performant because you are passing a reference and not a data structure
	sayGreetingByReference(&greeting, &name) //pass by value
	fmt.Println("Caller name " + name)

	sum(1, 2, 3, 4, 5)
	sumAlt("The sum is", 1, 2, 3, 4, 5)
	fmt.Println("The sum is", getSum(1, 2, 3, 4, 5))

	sumResult := getSumPointer(1, 2, 3, 4, 5)
	fmt.Println("The sumPointer is (1) ", sumResult) //you need to dereference
	fmt.Println("The sumPointer is (2) ", *sumResult)

	divResult := divide(5, 0) //return +inf. The program does not stop
	//divWithPanicResult := divideWithPanic(5, 0) //rThe program will stop
	fmt.Println(divResult)
	divResult2Types, err := divideWithTwoReturnTypes(5, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(divResult2Types)

	//in GO functions can be passed as parameters in functions
	//anonymous function.
	func() {
		msg := "Hello from anonymous function"
		fmt.Println(msg)
	}() //if you use these parentheses here, the code inside the function will be executed

	//Type signature is func()
	var anonymousFunc func(test string) = func(message string) {
		msg := "Hello from anonymous function 2"
		fmt.Println(msg + message)
	}
	anonymousFunc("Test this")

	//Methods
	g := greeter{
		greeting: "Hello",
		name:     "Go from method",
	}
	g.greet()
}

type greeter struct {
	greeting string
	name     string
}

/*
Method which executed in a known context. We get a copy (the value) of a greeter
You can pass a pointer also. eg -> func (g1 *greeter)
 */
func (g1 greeter) greet() {
	fmt.Println(g1.greeting, g1.name)
}

func divideWithTwoReturnTypes(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

func divide(a, b float64) float64 {
	return a / b
}

/*
The application will stop
*/
func divideWithPanic(a, b float64) float64 {
	if b == 0.0 {
		panic("Cannot provide zero as second value")
	}
	return a / b
}

func getSumPointer(values ...int) *int { //slice
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println("The sum is", result)
	return &result //in GO, this variable is promoted to be on the share memory (heap memory). The memory is not cleared
}

func getSum(values ...int) int { //slice
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println("The sum is", result)
	return result
}

func sum(values ...int) { //slice
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println("The sum is", result)
}

func sumAlt(message string, values ...int) { //slice
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println(message, result)
}

func sayGreetingByValue(greeting, name string) {
	fmt.Println(greeting, name)
	name = "Ted"
	fmt.Println(greeting, name)
}

func sayGreetingByReference(greeting, name *string) {
	fmt.Println(*greeting, *name)
	*name = "Ted"
	fmt.Println(*greeting, *name)
}

/**
Be careful when you pass around maps and slices because they passes pointers.
This is not happened when you use primitives and arrays
*/
func pointersOnArraysAndSlices() {
	a1 := [3]int{1, 2, 3} //this is an ARRAY
	b1 := a1              //pass by value because this is an ARRAY. Copy to a new instance
	fmt.Println(a1, b1)
	a1[1] = 42
	fmt.Println(a1, b1)

	//slice contains a pointer to the underline array.They are copying pointers
	a2 := []int{1, 2, 3} //this is an SLICE
	b2 := a2             //pass by reference because this is an SLICE. It copies the reference
	fmt.Println(a2, b2)
	a2[1] = 42
	fmt.Println(a2, b2)

	//same as slices happens with maps because a map contains pointers to underline data
	a3 := map[string]string{"foo": "bar", "baz": "buz"}
	b3 := a3
	fmt.Println(a3, b3)
	a3["foo"] = "qux"
	fmt.Println(a3, b3)
}

func pointersOnStructs() {
	var ms testStruct
	ms = testStruct{foo: 42}
	fmt.Println(ms)

	var msRef *testStruct
	msRef = &testStruct{foo: 43}
	fmt.Println(msRef)
	msRef.foo = 8978
	fmt.Println(msRef)

	var msNew *testStruct
	fmt.Println(msNew) //prints nil
	//fmt.Println(msNew.foo) //this will give a runTimeException
	msNew = new(testStruct)
	fmt.Println(msNew)
	(*msNew).foo = 45 //you can use directly msNew.foo without use the dereference first
	fmt.Println((*msNew).foo)
}

type testStruct struct {
	foo int
}

func pointersOnArrays() {
	a := [3]int{1, 2, 3}
	b := &a[0]
	c := &a[1]
	fmt.Printf("%v %p %p \n", a, b, c)

	//To take the value of the reference use *c,
	*c = 89
	fmt.Println(a)

}

func passByReferenceExample() {
	var test1 int = 42
	var test2 *int = &test1    //test2 is pointer to an integer and is pointing on test1
	fmt.Println(test1, test2)  //test2 print the numerical representation of the memory address which holds the test1
	fmt.Println(&test1, test2) //same result. Both print the address location

	//dereference operator. Check tha value of an address
	fmt.Println(&test1, *test2)
	test1 = 27
	fmt.Println(test1, *test2)

	*test2 = 11
	fmt.Println(test1, *test2)
}

func passByValueExample() {
	test1 := 42
	test2 := test1 //does not point to the same memory location
	test1 = 26
	fmt.Println(test1, test2)
	fmt.Println(test1, test2)
}

/*
Use an anonymous function to recover from an error. Call the recover function and check the error and the
execution will continue. If you use panic inside the anonymous function then the program execution will stop
and you will se the full stacktrace
*/
func panicker() {
	fmt.Println("about to panic")
	defer func() {                        //anonymous function
		if err := recover(); err != nil { //call recover function and check the error
			fmt.Println("Error:", err)
			//panic(err)	//if you want to panic then you need here to use the panic statement
		}
	}()
	panic("something bad happened")
	fmt.Println("done panicking")
}

/*
Resource request from http package. With defer you can associate the opening and closing of a resource the one
next to the other
*/
func runResourceRequest() {
	res, err := http.Get("https://restcountries.eu/rest/v2/name/greece")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}

func returnTrue() bool {
	fmt.Println("returning true")
	return true
}

/**
Executes the function pass to defer after the functions finishes the final statement but before returns
*/
func deferExample(message string) {
	fmt.Println(message)
}
