package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

// pass by value
func main() {
	dog := Dog{"hey d", 3, 4}
	dog.speak()
	fmt.Println("inside main: ", dog.Breed)
}

func (d Dog) speak() {
	d.Breed = "yoyo"
	fmt.Println("inside speak: ", d.Breed)
}

const TIME int = 5

var directions = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func solve(board [][]byte) {
	row := len(board)
	col := len(board[0])

	// first col and last column
	for i := 0; i < row; i++ {
		dfs(board, row, col, i, 0)
		dfs(board, row, col, i, col-1)
	}

	// first row
	for i := 0; i < col; i++ {
		dfs(board, row, col, 0, i)
		dfs(board, row, col, row-1, i)
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == 'P' {
				board[i][j] = 'O'
			}
		}
	}
}

func dfs(board [][]byte, row, col, i, j int) {
	if i < 0 || i == row || j < 0 || j == col || board[i][j] != 'O' {
		return
	}
	board[i][j] = 'P'
	for _, d := range directions {
		newx := i + d[0]
		newy := j + d[1]
		dfs(board, row, col, newx, newy)
	}
}

// structs
func main7() {

	dog := Dog{"Golden retriever", 10, 20}
	fmt.Println(dog)
	fmt.Printf("%+v\n", dog)
	fmt.Println(dog.age) // private field don't know how it's accessible

	for e, d := range directions {
		fmt.Println(e)
		fmt.Println(d)
	}

	dir := []int{11, 21, 31}
	for _, d := range dir {
		fmt.Println()
		fmt.Println(d)
	}
	board := [][]string{{"adarsh", "jaiswal", " adarsh"}, {"adarsh2", "jaiswal2", " adarsh2"}}
	for _, boa := range board {
		for _, b := range boa {
			fmt.Println(b)
		}
	}

}

type Dog struct {
	Breed  string
	Weight int
	age    int
}

// maps
func main6() {
	states := make(map[string]string)
	states["CA"] = "California"
	states["WA"] = "Washington"
	states["OR"] = "Oregon"
	states["NY"] = "New York"
	fmt.Println(states)

	cali := states["CA"]
	fmt.Println(cali)

	delete(states, "OR")
	fmt.Println(states)

	for k, v := range states {
		fmt.Printf("key: %v, val: %v \n", k, v)
	}

	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	fmt.Println("The keys: ", keys)
	sort.Strings(keys)
	fmt.Println("The keys: ", keys)
}

// arrays and slices
func main5() {
	var colors [3]string
	colors[0] = "red"
	fmt.Println(colors)

	numbers := [2]int{1, 2}
	fmt.Println(numbers)
	fmt.Println(len(numbers))

	// slices
	var color2 = []string{"red", "yellow"}
	fmt.Println(color2)
	color2 = append(color2, "orange")
	fmt.Println(color2)

	color2 = color2[1:]
	fmt.Println(color2)

	color3 := color2[1:]
	fmt.Println(color3)

	numbers3 := make([]int, 5)
	numbers3[0] = 1
	fmt.Println(numbers3)

	numbers3 = make([]int, 2, 12)
	numbers3[0] = 1

	numbers3 = append(numbers3, 40)
	numbers4 := append(numbers3, 50)
	fmt.Println(numbers3)
	fmt.Println(numbers4)
	sort.Ints(numbers4)
	fmt.Println(numbers4)

}

// pointers
func main4() {

	anInt := 42
	var p = &anInt
	fmt.Println("value of anInt: ", *p)

	*p = 44
	fmt.Println("value of anInt: ", anInt)
	fmt.Println("value of anInt via pointer: ", *p)

	*p = *p / 2
	fmt.Println("value of anInt: ", anInt)
	fmt.Println("value of anInt via pointer: ", *p)

}

func main3() {

	t := time.Now()
	fmt.Println(t.String())

	fmt.Println(t.Format("01/02/2006"))

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("ENter text: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("You entered: ", input)

	fmt.Print("Enter a number: ")
	numInput, _ := reader.ReadString('\n')
	aFloat, err := strconv.ParseFloat(strings.TrimSpace(numInput), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("you entered: ", aFloat)
	}

	fmt.Println("%T", numInput)
}

func main2() {
	fmt.Println("Hello adarsh")

	var s string = "adarsh"
	fmt.Println("this is new string", s)

	var in = 5
	fmt.Printf("this %d integer\n", in)

	var in2 int = 10
	fmt.Printf("this %d integer\n", in2)

	anotherIn := 20 // works only inside a function
	fmt.Printf("this %d is new integer\n", anotherIn)

	fmt.Printf("this %d is time", TIME)
}
