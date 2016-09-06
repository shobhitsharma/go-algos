package main

import "fmt"
import "os"
import "strconv"
import "strings"

func sort(unsorted []int) []int {
	changed := true
	for changed {
		changed = false
		for i := 0; i < len(unsorted)-1; i++ {
			if unsorted[i+1] < unsorted[i] {
				temp := unsorted[i+1]
				unsorted[i+1] = unsorted[i]
				unsorted[i] = temp
				changed = true
			}
		}
	}
	return unsorted
}
func convert(args []string) []int {
	foo := make([]int, len(args))
	for _, element := range args {
		values := strings.Split(element, ",")
		foo = make([]int, len(values))
		for j := 0; j < len(values); j++ {
			fmt.Print("\n", values[j])
			foo[j], _ = strconv.Atoi(values[j])
		}
	}
	return foo
}

func main() {
	//get the argument array
	arg := os.Args[1:]
	fmt.Print(arg)
	unsorted := convert(arg)
	fmt.Print("Unsorted list ", unsorted)
	foo := sort(unsorted)
	fmt.Println(foo)
}
