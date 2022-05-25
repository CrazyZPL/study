package main

import (
	"fmt"
	// "math"
	"strings"
)

// for .. range...
// func main() {
// 	for k, v := range []int{1, 2, 3, 4, 5} {
// 		fmt.Println("now i is", k, v)
// 	}
// }

// use newton to realize sqrt function
// func sqrt(x float64) float64 {
// 	z := 1.0
// 	for i := 0; i < 10; i++ {
// 		z -= (z*z - x) / (2 * z)
// 	}
// 	return z
// }
// func main() {
// 	fmt.Println(sqrt(344), math.Sqrt(344))
// }

// defer
// func main() {
// 	// i := 0
// 	// fmt.Println("first", i)

// 	// defer fmt.Println("defer i", i) // 0

// 	// i = 1

// 	for i := 0; i < 5; i++ {
// 		defer fmt.Println(i) // 4, 3, 2, 1, 0
// 	}
// 	fmt.Println(c()) // 2
// }
// func c() (i int) {
// 	defer func() { i++ }()
// 	return 1
// }

// defer, panic and recover
// func main() {
// 	f()
// 	fmt.Println("Returned normally from f.")
// }
// func f() {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			fmt.Println("Recovered in f", r)
// 		}
// 	}()
// 	fmt.Println("Calling g.")
// 	g(0)
// 	fmt.Println("Returned normally from g.")
// }
// func g(i int) {
// 	if i > 3 {
// 		fmt.Println("Panicking!")
// 		panic(fmt.Sprintf("%v", i))
// 	}
// 	defer fmt.Println("Defer in g", i)
// 	fmt.Println("Printing in g", i)
// 	g(i + 1)
// }

// pointers
// func main() {
// 	i, j := 11, 14
// 	p := &i
// 	fmt.Printf("p is %v, *p is %v, &p is %v\n", p, *p, &p)
// 	p = &j
// 	*p = *p / 7
// 	fmt.Println("j is", j)
// }

// struct
// type Person struct {
// 	Name   string
// 	Age    uint32
// 	Gender string
// }
// func main() {
// 	person1 := Person{
// 		Name:   "zpl",
// 		Age:    23,
// 		Gender: "male",
// 	}
// 	fmt.Println("person1 is", person1)
// 	person2 := Person{}
// 	person2.Name = "lpz"
// 	person2.Age = 23
// 	person2.Gender = "female"
// 	fmt.Println("person2 is", person2)
// }

// array
// func main() {
// 	numbers := []uint32{1, 2, 3, 4, 5}
// 	var otherNumbers [10]uint32
// 	otherNumbers[0] = 1
// 	otherNumbers[1] = 2
// 	fmt.Printf("numbers are %v, and other numbers are %v", numbers, otherNumbers)
// }

//slice
// The length of a slice is the number of elements it contains.
// The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
// The zero value of a slice is nil.
// The make function allocates a zeroed array and returns a slice that refers to that array
// func main() {
// 	var numbers []uint32
// 	nums := []uint32{1, 2, 3, 4, 5}
// 	numbers = append(numbers, nums...)
// 	fmt.Println(numbers, nums)
// 	numbers[0] = 11
// 	fmt.Println(numbers, nums)
// 	numbers = nums[:3]
// 	fmt.Println(numbers)
// 	numbers[0] = 11
// 	fmt.Println(numbers, nums)
// 	numbers = []uint32{}
// 	numbers = append(numbers, 1, 2, 3, 4, 5)
// 	fmt.Println(numbers)
// }

// map
type Vertex struct {
	Lat, Long float64
}

func main() {
	m := make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.1, -74.3,
	}
	fmt.Println(m, m["Bell Labs"])
}
func WordCount(s string) map[string]int {
	strings := strings.Fields(s)
	WordsCount := make(map[string]int)
	for _, v := range strings {
		WordsCount[v] += 1
	}
	return WordsCount
}