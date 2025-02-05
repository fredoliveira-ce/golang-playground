package typeof

import "fmt"
import "reflect"

func showVariableType()  {
	var name string = "Fred"
	company := "Fred"
	var version float32 = 1.1
	var aFloatValue = 1.1
	var age int = 24

	fmt.Println("Hello, Sr.!", name, ", we're at the vertion", version)
	fmt.Println("Your age is", age)

	fmt.Println("The type of the variable name is: ", reflect.TypeOf(name))
	fmt.Println("The type of the variable aFloatValue is: ", reflect.TypeOf(aFloatValue))
	fmt.Println("The type of the variable company is: ", reflect.TypeOf(company))
	fmt.Println("The type of the variable version is: ", reflect.TypeOf(version))

	var b byte = 255
	fmt.Println("The type of the variable b is:", b) //have to print uint8

	var unicode rune = 'a'
	fmt.Println("The type of the variable rune is: ", reflect.TypeOf(unicode))
}