/* datatypes & variables */

/*
	string
	bool

	int
	int8
	int16
	int32
	int64

	uint
	uint8
	uint16
	uint32
	uint64

	float32
	float64

	complex64 (real [float32] + imaginary [float32])
	complex128 (real [float64] + imaginary [float64])

	byte
	rune (unicode character code point)

	NOTE : No implicit type conversions
*/

package main

import "fmt"

/* package level declarations */

//no := 100

var no int

//var no int = 100
//var no = 100

func main() {
	/*
		var x int
		x = 10
	*/
	/*
		var x int = 10
	*/

	//type inference
	/*
		var x = 10
	*/

	x := 10 /* := syntax is application only for function scoped variables (NOT for package scope) */
	fmt.Println(x)

	//var n1, n2 int
	//var n1, n2 int = 100, 200
	n1, n2 := 100, 200
	fmt.Println(n1, n2)

	var (
		y1, y2 int
		z      string
	)
	y1 = 100
	y2 = 200
	z = "a dummy string"
	fmt.Println(y1, y2, z)

	//type conversion
	var x1 int32 = 100
	//var f1 float32 = x1 //implicit type conversions are not supported
	var f1 float32 = float32(x1)
	fmt.Println(f1)

	//constant
	const pi float32 = 3.14

	//iota
	const (
		RED = 1 << iota
		GREEN
		_
		BLUE
	)

	fmt.Printf("RED = %b, GREEN = %b, BLUE = %b\n", RED, GREEN, BLUE)

}

func f1() {

}
