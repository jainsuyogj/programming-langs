# programming-langs
Learning programming GoLang

https://www.youtube.com/watch?v=CF9S4QZuV30&feature=youtu.be

Syntax: 

Hello.go 

1) Declare a variable using "var <var-name> <var-type> = <value>" or auto deduced "<var> := <value>". 
Ex: a) var myVarName int = 10
    b) myVarName := 10 (now myVarName is an int) (statically typed - now myVarName cannot be changed to string or any other data type)

2) Strings :
    var myVarName string = "Test Me"

    a) Most of the things are functions:
    Like len(myVarName) but not myVarName.length()

3) Prints : 
    fmt package used for formatting.
    Println - Prints with new line
    Printf - similar to 'C' printf()

    Print dataType using "%T".

4) Boolean && , ||, !
    Supports boolean and (&&) or (||) not(!)

5) Constants
    const pi float = 3.14...

6) For loops :
    1) Using just conditions
    2) Using init var, condition, increment

7) If Statements: 

Arrays.go

1) Arrays init.
    var arrName [numOfElems]int

2) Slices.
    slice := []int {1,2,2,2}
    slice := make([]int, <default>, <size>)
    slice := slice[start,end]
    slice := slice[:end] - only want the 0->end values, but the slice still holds all the values and can be expanded again by [:actualend]. actual length does not reduce here.
    <to find len of slice vs actual internal size.>
    slice := slice[start:] - only want to start->actualend values.... but now the slices would never contain the 0->start values and cannot be recovered. actual length reduces here/

    copy(toSlice, fromSlice)

3) Maps:
    Key value pairs.
    remove() func to remove.

    directly maps

4) Functions:
    function for structured code.
    multi value returning.

5) Variadic function:
    any number of values.

    accept in func using (args ...int)

5) Closures (function pointers, but valid inside a function only)
    func some() {
        closureName := func() int {
            someOperation
            return int
        }
        fmt.Println(closureName())
    }

6) Defer:
    It's an operation to be performed at the end of the function call

7) recover()
    It's like catch in Generic exception and then not rethrow

8) PANIC() 
    It's like throwing a exception.

9) Pointers:
    Similar to C

10) Struct:
    a collection of values.

11) Interface:
    any struct which implements all the func of the interface, can share the pointer.
    be it a fruit and shape

12) String utilities functions:
    strings.contains()
    strings.Index()

13) File IO
    os.<functions> to open or read a file.
    ioutils.<functions> to read a File

14) http functions.
    easy handler. http.Handler(path, callback)
    start easy http.Lis...

15) Go routines:
    Threads

16) Channels :
        chanName <- value
        Send the value on the channel

        value := <- chaName
        Get the value from the channel
