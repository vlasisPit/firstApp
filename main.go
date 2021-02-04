package main

import (
	"fmt" //	Package fmt implements formatted I/O with functions analogous to C's printf and scanf.
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

}

func returnTrue() bool {
	fmt.Println("returning true")
	return true
}
