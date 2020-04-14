package main

import (
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

/*
	This is how a comments is written
*/
func basics() {
	fmt.Println("Hello World")

	var age int = 40

	var number float32 = 1.2

	//Auto deduced types
	inti := 1

	// Strings
	var stringVar string = "string"

	fmt.Println(age, number,
		inti, stringVar)

	// boolean &&,  || , !
	fmt.Println("true && false = ", true && false)

	// loops
	i := 1

	for i <= 10 {
		fmt.Println(i)
		i++
	}

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
	defer fmt.Println("second")
	fmt.Println("first")
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

func main() {

	//basics()

	//arrays()

	//factorial
	//fmt.Println(factorial(5))

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
}
