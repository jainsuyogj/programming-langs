package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Youtube video

	//basics()

	//arrays()

	//factorial
	/*fmt.Println(" Factorial of five using recursion ", factorial(5))

	if val, err := factorialNonRecursive(5); err == nil {
		fmt.Println(" Factorial of five using loop ", val)
	} else {
		fmt.Println(" Error occured in factorial ", err)
	}
	*/

	//defer
	//deferMe()

	//recover
	//recoverMe(2, 0)

	//panic
	//panicMe()

	//pointer
	//numVal := 1 // OR numVal2 := new(int)
	//pointerChange(&numVal) // OR pointerChange(numVal2)
	//fmt.Println(numVal) // OR fmt.Println(*numVal2)

	//structures
	//rect := Rectangle{10, 20}
	//rect.printMe()

	// interface
	//rect1 := Rectangle{10, 20}
	//frut := Fruit{}
	//printMeInterfaceTest(rect1)
	//printMeInterfaceTest(frut)

	//strings
	//stringsOps()

	//File IO
	//fileIOTest()

	//parser functions
	//castFunctions()

	//http functions
	//httpFuncs()

	//go routines
	//goRoutTest()

	//channels
	//channelsTest()

	//gobyexample.com

	linkedList()
}

// Node representation
type nodeVal struct {
	val int
}

//Linked list node
type linkedListNode struct {
	data nodeVal
	next *linkedListNode
}

func (node *linkedListNode) insert(inp nodeVal) {
	//recursive
	if node.next == nil {
		node.next = &linkedListNode{inp, nil}
		return
	} else {
		node.next.insert(inp)
	}

}

func linkedList() {
	fmt.Println("Linked list demo: ")
	head := &linkedListNode{nodeVal{1}, nil}
	head.insert(nodeVal{2})
	tempPtr := head
	for tempPtr != nil {
		fmt.Print(tempPtr.data)
		tempPtr = tempPtr.next
	}
}

func factorialNonRecursive(inputNumber int) (uint64, error) {

	if inputNumber < 0 {
		return 0, errors.New("Input number is less than 0")
	}

	var ans uint64 = 1
	for i := 1; i <= inputNumber; i++ {
		if ans = ans * uint64(i); ans <= 0 {
			return 0, errors.New("Out of range of 64 bytes number at " + strconv.Itoa(inputNumber))
		}
	}

	return ans, nil

}

func printAr(arr []float64) {

	for _, value := range arr {
		fmt.Print(value)
	}
	fmt.Println()

}

func arrays() {

	//Arrays
	var favNums [2]int

	favNums[0] = 1
	favNums[1] = 2

	fmt.Println("Printing index 2 of array 1,2 ", favNums[1])

	newFav := [5]float64{1, 2, 3, 4, 5}
	fmt.Println("Using range to print the index and value in the array:")
	for i, value := range newFav {
		fmt.Print(value, i, " .. ")
	}

	//I don't want i
	fmt.Println("Ranging without index")
	for _, value := range newFav {
		fmt.Println(value)
	}

	fmt.Println("Starting slices, What are they? Let's print")

	numSlice := []float64{5, 4, 3, 2, 1}

	for _, value := range numSlice {
		fmt.Print(value, " ")
	}

	numSlice = numSlice[3:5]

	numSlice = numSlice[:1]

	numSlice = numSlice[1:]

	fmt.Print("Performed numSlice[3:5], numSlice[:1], numSlice[1:] on the array above gives me: ")
	printAr(numSlice)

	numSliceMake := make([]int, 5, 10)

	fmt.Print("Printing the slict using fmt will give address: ")
	fmt.Println(numSliceMake)

	fmt.Println("Using print inbuilt function I guess? ")
	print(numSlice)
	fmt.Println()

	//Maps
	// key values pairs
	mapMe := make(map[string]int)
	mapMe["test"] = 42
	fmt.Println("Value of key test in the map ", mapMe["test"])
	fmt.Println("Length of map: ", len(mapMe))

	delete(mapMe, "test")
	fmt.Println("Length of map after deleting: ", len(mapMe))

	//Functions:
	testFunc()

	fmt.Print("Testing multi return value: ")
	// multi ret func
	my1, my2 := multiRetFunc()
	fmt.Println(my1, my2)

	//Variadic funcs
	fmt.Println("Var args ", varArgs(1, 2, 3))
}

func testFunc() {
	fmt.Println("Learning functions")
}

func multiRetFunc() (int, int) {
	myInt := 1
	myInt2 := 2
	return myInt, myInt2
}

func varArgs(args ...int) int {
	returnVal := 0
	for _, value := range args {
		returnVal += value
	}
	return returnVal
}

/*
	This is how a comments is written
*/
func basics() {

	fmt.Println(" I am starting to learn Hello World")

	var age int = 40

	var number float32 = 1.2

	//Auto deduced types
	inti := 1

	fmt.Printf("Auto deduced type is %T \n", inti)

	// Strings
	var stringVar string = "string"

	fmt.Println("Printing the datatypes known ", age, number,
		inti, stringVar)

	// boolean &&,  || , !
	fmt.Println("true && false = ", true && false)

	// loops
	i := 1

	fmt.Println("While loop alternative in go ")
	for i <= 10 {
		fmt.Print(i)
		i++
	}

	fmt.Println()

	fmt.Println("Learning if/else, switch conditions")
	yourAge := 10
	if yourAge == 8 {
		fmt.Println("age is 10")
	} else if yourAge > 8 {
		fmt.Println("age is greater than 8")
	} else {
		fmt.Println("Age is lesser that 8")
	}

	//Switch
	myAge := 11

	switch myAge {
	case 10:
		fmt.Println("Age is 10")
	case 11:
		fmt.Println("age is 11")
	default:
		fmt.Println("age is unknown")
	}
}

func factorial(input int) int {
	if input < 2 {
		return 1
	}
	return factorial(input-1) * input
}

func deferMe() {
	// It will print first and then second..
	//Defer makes sure that it will always execute at the end of function return
	defer fmt.Println("Ok, as you say")
	fmt.Println("I have written the code for defer before this, but it prints after me")
}

func recoverMe(num1, num2 int) {
	defer func() {
		fmt.Println(recover())
	}()

	fmt.Println(num1 / num2)
}

func panicMe() {
	defer func() {
		fmt.Println(recover())
	}()

	panic("panicMe")
}

func pointerChange(num *int) {
	*num++
}

// Inter me
type Inter interface {
	printMe()
}

// Rectangle asda
type Rectangle struct {
	leftX int64
	leftY int64
}

func (rect Rectangle) printMe() {
	fmt.Println(rect.leftX, rect.leftY)
}

// Fruit asd
type Fruit struct {
}

func (frt Fruit) printMe() {
	fmt.Println("I am a fruit")
}

func printMeInterfaceTest(inter Inter) {
	inter.printMe()
}

func stringsOps() {
	test := "stringMe"

	// Contains strings
	fmt.Println(" contains ", strings.Contains(test, "Me"))

	testSplit := "1,2,3,4"

	//SPLIT strings
	fmt.Println(" split ", strings.Split(testSplit, ","))

	testArr := []string{"arr2", "arr1"}

	//SORT strings
	sort.Strings(testArr)

	fmt.Println(testArr)

	// joined strings
	joinedString := strings.Join([]string{"1", "2", "3"}, ", ")

	fmt.Println(joinedString)

}

func fileIOTest() {
	file, err := os.Create("test.txt")

	if err != nil {
		log.Fatal(err, " failed to create ")
	}

	file.WriteString("GO test")

	file.Close()

	stream, err := ioutil.ReadFile("test.txt")

	if err != nil {
		log.Fatal(err, " failed to open")
	}

	readVal := string(stream)

	fmt.Println(readVal)

}

func castFunctions() {
	numString := "1"
	num, _ := strconv.ParseInt(numString, 0, 64)
	fmt.Println(" cast succes ", num)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello worlds")
}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello worlds")
}

func httpFuncs() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/test", handler2)
	http.ListenAndServe(":8080", nil)
}

func count(id int) {
	for i := 0; i < 10; i++ {
		fmt.Println(id, i)
		time.Sleep(time.Second * 1)
	}
}

func goRoutTest() {
	for i := 0; i < 10; i++ {
		go count(i)
	}

	time.Sleep(time.Second * 15)
	fmt.Println("finished")
}

//PizzaNum as
var PizzaNum = 0

//PizzaName asd
var PizzaName = ""

func makeDough(chanName chan string) {
	PizzaNum++

	PizzaName = "Pizza #" + strconv.Itoa(PizzaNum)

	fmt.Println("Make dough ", PizzaName, "and send for sauce")
	chanName <- PizzaName

	time.Sleep(time.Millisecond * 100)
}

func makeSauce(chanName chan string) {
	pizzaRcv := <-chanName

	fmt.Println("Make sauce", pizzaRcv, " and send for toppings")
	chanName <- pizzaRcv

	time.Sleep(time.Millisecond * 100)
}

func maketops(chanName chan string) {
	pizza := <-chanName
	fmt.Println(" Add toppings to ", pizza, " and ship")
	time.Sleep(time.Millisecond * 100)
}

func channelsTest() {
	stringChan := make(chan string)

	for i := 0; i < 1; i++ {
		go makeDough(stringChan)
		go makeSauce(stringChan)

		go maketops(stringChan)

		time.Sleep(time.Second * 5)
	}
}
