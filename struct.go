package main
import (
	"fmt"
	// "math"
)

func pointer() {
	i, j := 42, 2701
	p := &i // point to i
	fmt.Println("p point to i: ")
	fmt.Println("-- p: ", *p)
	fmt.Println("-- i: ", i)
	*p = 21         // set i through the pointer
	fmt.Println("-- set i through the pointer & new i: ", i)  // see the new value of i
	
	fmt.Println("\n p point to j: ")
	p = &j         // point to j
	*p = *p - 701   // divide j through the pointer
	fmt.Println("-- p: ", *p)
	fmt.Println("-- set j through the pointer & new j: ", j)
}


type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	p = &Vertex{1,2}
)

func hypot1 (x,y float64) float64 {
	return x + y
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3,4)
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}

	c <- sum //send sum to c
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}


func main () {

	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}



	// pointer()

	// hypot := func(x,y float64) float64 {
	// 	return x-y
	// 	// return math.Sqrt(x*x + y*y)
	// }
	// fmt.Println(hypot(5, 12))
	// fmt.Println("\n compute: ", compute(hypot1))
	// fmt.Println(compute(math.Pow))

	// channels
	// s := []int{7,2,8,-9,4,0}
	// c := make(chan int)
	
	// go sum(s[:len(s)/2], c)
	// go sum(s[len(s)/2:], c)
	// x, y := <-c, <-c

	// fmt.Println("--x: ", x)
	// fmt.Println("--y: ", y)
	// fmt.Println("--x: ", x)


}