package main
import "fmt"
import "unicode/utf8"

func main(){
	var intNum int = 32767 //var "variablename" variable_type, go also has int8, int16, int32, int64 to sepcify how much memory or bits is needed to store the number, if not defined which int to use, go will use one depending on the architecture x32/x64
	// there are also "uint", it can be understood as int8 stores (-128, 127) and uint8 stores (0, 255)
	intNum = intNum + 1
	fmt.Println(intNum)


	var floatNum float32 = 125.651 //float32 and float64, used to store decimal numbers
	fmt.Println(floatNum)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32) //typecasting
	fmt.Println(result)

	var intNum1 int8 = 3
	var intNum2 int8 = 2
	fmt.Println(intNum1/intNum2) //for division
	fmt.Println(intNum1%intNum2) //for remainder

	var myString string = "Hello World"
	fmt.Println(myString)

	var myString2 string = "Hello \nWorld"
	fmt.Println(myString2)

	var myString3 string = `Hello
world`
	fmt.Println(myString3)

	var myString4 string = "Hello" + " " + "World" 
	fmt.Println(myString4)

	fmt.Println(len("test")) //output will be 4, which is the number of bytes
	//to get the accurate length can be done with importing the builtin package "unicode/utf8"

	//for examle:
	fmt.Println(len("γΔ")) //output is 4

	fmt.Println(utf8.RuneCountInString("γΔ")) //output is 2

	var myBool bool = false
	var myBool2 bool = true

	fmt.Println(myBool)
	fmt.Println(myBool2)

	// the default value for empty
	// integers = 0
	// strings = ""
	// boolean = false

	//ANOTHER WAY OF DECLARING VARIABLE ARE:
	// var myVar = "text" //the type of the variable is inferred if we set the value of the variable right away
	// myVar :=  "text" 
	//but adding the type is a good practise'

	const myConst string = "constant variables"
	fmt.Println(myConst)

}