package basicoperations

import "fmt"

func ArithmeticOps() {
	// prompt user for a number
	var firstNumber int32
	var secondNumber int32
	fmt.Println("Doing arithmetic operations:")
	fmt.Print("Enter the first Number: ")
	fmt.Scan(&firstNumber)
	fmt.Print("Enter the second Number: ")
	fmt.Scan(&secondNumber)
	fmt.Println("Summing up the 2 numbers")
	addNumbers(int(firstNumber), int(secondNumber))
	fmt.Println("The difference between the 2 numbers")
	subtractNumbers(int(firstNumber), int(secondNumber))
	fmt.Println("The multiple of the numbers are: ")
	multiplyNumbers(int(firstNumber), int(secondNumber))
	fmt.Println("Division of the numbers are: ")
	divideNumbers(int(firstNumber), int(secondNumber))
	fmt.Println("The remainder for the above division is: ")
	modulus(int(firstNumber), int(secondNumber))

	ifState()
	traditionalFor()
	rangeFor()
}

func addNumbers(fNum, SNum int) {
	fmt.Printf("%d + %d = %d \n", fNum, SNum, fNum+SNum)
}

func subtractNumbers(fNum int, sNum int) {
	fmt.Printf("%d - %d = %d \n", sNum, fNum, sNum-fNum)
}

func multiplyNumbers(fNum, sNum int) {
	fmt.Printf("%d x %d = %d \n", sNum, fNum, fNum*sNum)
}

func divideNumbers(fNum, sNum int) {
	var division float32 = float32(sNum) / float32(fNum)
	message := fmt.Sprintf("%d / %d = %.4f \n", sNum, fNum, division)
	fmt.Println(message)
}

func modulus(fNum, sNum int) {
	var mod int = sNum % fNum
	fmt.Printf("%d mod %d = %d \n", sNum, fNum, mod)
}
