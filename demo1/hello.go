package main

import (
	"fmt"
	"io/ioutil"
)


func enums() {
	const (
		c1 = iota
		c2
		c3
		c4
	)
	fmt.Println(c1,c2,c3,c4)
}

func if_else(){
	var age int = 27
	if age > 25 {
		fmt.Println("我已经超过25岁了！")
	} else if age < 25 {
		fmt.Println("我还没有25岁！")
	} else {
		fmt.Println("我就是25岁！")
	}
}

func is_exists_file() {
	const filename = "abc.txt"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}

func is_exists_file1() {
	const filename = "abc.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}

// panic: Wrong score: 105
func grade(score int) string {
	var g string = ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

func sum1(number int) int {
	var sum int = 0
	for i := 0; i <= number; i++ {
		sum += i
	}
	return sum
}

func pointer() {
	var a int = 2
	fmt.Println("start:", &a)
	var pa *int = &a
	*pa = 3
	fmt.Println(pa, *pa , &*pa, a, &a)
}

func arr() {
	var arr1 [5]int
	arr2 := [3]int{1,3,5}
	arr3 := [...]int{2,4,6,8,10}
	slice1 := arr3[2:4]
	var grid [4][5]int

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	for i := 0; i< len(arr3); i++ {
		fmt.Println(arr3[i])
	}
	fmt.Println("---------------------------")
	for i := range arr2 {
		fmt.Println(i, arr2[i])
	} 
	fmt.Println("---------------------------")
	for i, v := range grid {
		fmt.Println(i, v)
	}
	fmt.Println(slice1)
}

func main() {
	var s string  = "hello world" 
	fmt.Println(s)
	enums()
	if_else()
	is_exists_file()
	fmt.Println(
		grade(0),
		grade(59),
		grade(60),
		grade(70),
		grade(85),
		grade(95),
		grade(100),
		// grade(105),
		// grade(-1),
		sum1(10),
	)
	fmt.Println(sum1(5))
	pointer()
	arr()
}


