package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"math"
	"sort"
)

func read_file(file string) interface{} {
	data, err := ioutil.ReadFile(file)
	var listA []float64
	var listB []float64
	if err != nil {
		fmt.Println("File reading error", err)
		return nil
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		columns := strings.Fields(line)
		if len(columns) != 2 {
			fmt.Println("Invalid line:", line)
			continue
		}
		a, errA := strconv.Atoi(columns[0])
		b, errB := strconv.Atoi(columns[1])
		if errA != nil || errB != nil {
			fmt.Println("Invalid number in line:", line)
			continue
		}
		listA = append(listA, float64(a))
		listB = append(listB, float64(b))
	}
	return []interface{}{listA, listB}
}

func compute_difference(listA []float64, listB []float64) float64 {
	var difference float64
	for j, item := range listA {
		diff := listB[j] - item
		abs_diff := math.Abs(diff)
		fmt.Println("Difference between", item, "and", listB[j], "is", abs_diff)
		difference += abs_diff
	}
	return difference
}

func compute_similarity(listA []float64, listB []float64) float64 {
	var similarity float64
	for _, item1 := range listA {
		this_similarity := 0.0
		for _, item2 := range listB {
			if item1 == item2 {
				this_similarity += 1
			}
		}
		similarity = similarity + (this_similarity * item1)
	}
	return similarity
}

func main() {
	fmt.Println("Computing difference between two lists")
	lists := read_file("input.txt")
	if lists == nil || len(lists.([]interface{})) < 2 {
		return
	}
	listA := lists.([]interface{})[0].([]float64)
	listB := lists.([]interface{})[1].([]float64)

	sort.Float64s(listA)
	sort.Float64s(listB)	

	fmt.Printf("%f\n", compute_similarity(listA, listB))
	
}
