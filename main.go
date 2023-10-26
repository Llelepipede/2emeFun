package main

import "fmt"

func main() {

	fmt.Printf("sumIntervals([][]int{{1, 2}, {6, 10}, {11, 15}}) %v\n", sumIntervals([][]int{{1, 2}, {6, 10}, {11, 15}}))
	fmt.Printf("sumIntervals([][]int{{1, 4}, {7, 10}, {3, 5}}) %v\n", sumIntervals([][]int{{1, 4}, {7, 10}, {3, 5}}))
}

func sumIntervals(tab [][]int) int {
	var ret int
	var ref []bool
	var max, min int
	ref, min, max = CreateMaxIntervale(tab)
	var temp int = 0
	for _, ints := range tab {
		ref, temp = CalcTab(ref, min, max, ints)
		ret += temp
		print(ret)
	}
	return ret
}

func CreateMaxIntervale(tab [][]int) ([]bool, int, int) {
	var ret []bool
	var min int = 0
	var max int = 0
	for i, inter := range tab {
		for y, value := range inter {
			if i == 0 && y == 0 {
				min = value
			} else if i == 0 && y == 1 {
				max = value
			} else {
				if min > value {
					min = value
				}
				if max < value {
					max = value
				}
			}
		}
	}
	var reach int = max - min

	for i := 0; i < reach; i++ {
		ret = append(ret, false)
	}
	return ret, min, max
}

func CalcTab(ref []bool, min int, max int, inter []int) ([]bool, int) {
	var ret int = 0
	for i := inter[0]; i < inter[1]; i++ {
		if !ref[i-min] {
			ref[i-min] = true
			ret++
		}
	}

	return ref, ret
}
