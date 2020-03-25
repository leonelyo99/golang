package main
import "fmt"

func main(){

	//=====================================================
	//Numericos
	//=====================================================

	/*
	uint8       the set of all unsigned  8-bit integers (0 to 255)
	uint16      the set of all unsigned 16-bit integers (0 to 65535)
	uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
	uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

	int8        the set of all signed  8-bit integers (-128 to 127)
	int16       the set of all signed 16-bit integers (-32768 to 32767)
	int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
	int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

	float32     the set of all IEEE-754 32-bit floating-point numbers
	float64     the set of all IEEE-754 64-bit floating-point numbers

	complex64   the set of all complex numbers with float32 real and imaginary parts
	complex128  the set of all complex numbers with float64 real and imaginary parts

	byte        alias for uint8
	rune        alias for int32
	*/

	var a int = 12//como mi pc es de 64 bits lo compilara como int64
	var b int8 = 23
	c:=a+int(b)
	d:= "leonel"
	
	fmt.Println(c);
	fmt.Printf("Hola %s %d como estas?",d,b)
	//=====================================================
	//String
	//=====================================================

	//=====================================================
	//Array
	//=====================================================
	var ArrayA [2]string
	ArrayA[0] = "Hello"
	ArrayA[1] = "World"
	fmt.Println(ArrayA[0], ArrayA[1])
	fmt.Println(ArrayA)

	ArrayPrimes := [6]int{2, 3, 5, 7, 11, 13} //creo un array de 6 lugares de tipo entero, y le asigno valores
	fmt.Println(ArrayPrimes)

	//=====================================================
	//Slice
	//=====================================================
	//un slice es un pedaso del array original
	SlicePrimes := [6]int{2, 3, 5, 7, 11, 13}

	var SliceS []int = SlicePrimes[1:4] //datos de la pocision 1 a la 4
	fmt.Println(SliceS)

	//=====================================================
	//Struct
	//=====================================================
	//estructura de elementos clave valor de un tipo determinado
	//un diccionario
	type Vertex struct {
		X int
		Y int
	}

	fmt.Println(Vertex{1, 2})

	//=====================================================
	//Pointer
	//=====================================================
	//guardar el espacio en memoria de otra variable
	i := 42

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	//=====================================================
	//Chanel
	//=====================================================

	//=====================================================
	//Function
	//=====================================================

	//=====================================================
	//Interce
	//=====================================================

	//=====================================================
	//Map
	//=====================================================
}
