package main

import "fmt"

func main() {
	fmt.Println("I am learning array in golang")
	var name [25]string
	name[0] = "Pawan"
	name[1] = "Kumar"
	name[2] = "Singh"
	name[3] = "Gupta"
	name[4] = "Radhe Krishna"
	fmt.Println(name)

	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Numbers are ", numbers)
	fmt.Println("Length of array is ", len(numbers))
	fmt.Println("Value at index 3 is ", name[3])

	var price [10]string
	fmt.Println("Price is ", price)

}
