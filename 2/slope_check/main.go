package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func read_file(file string) [][]float64 {
	data, err := ioutil.ReadFile(file)
	var matrix [][]float64
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
		if len(columns) == 0 {
			//fmt.Println("Invalid line:", line)
			continue
		}
		var list []float64
		for _, column := range columns {
			number, err := strconv.Atoi(column)
			if err != nil {
				//fmt.Println("Invalid number in line:", line)
				continue
			}
			list = append(list, float64(number))
		}
		matrix = append(matrix, list)
	}
	return matrix
}

func check_slope(matrix [][]float64) int {
	// Iterate over every row
	var safetyStatus []bool
	var safeCount int
	for _, line := range matrix {
		//fmt.Println(i, line)
		lastNum := 0
		lastDirection := 0
		safe := true
		for _, column := range line {
			//fmt.Println("Checking", column)
			if lastNum == 0 {
				//fmt.Println("First number")
				lastNum = int(column)
				continue
			}
			if lastNum < int(column) {
				if int(column) - lastNum > 3 {
					fmt.Println("Unsafe, too big of a jump from", lastNum, "to", column)
					safe = false
					continue
				}
				//fmt.Println("Increasing")
				if lastDirection == 1 || lastDirection == 0 {
					//fmt.Println("Direction is", lastDirection)
					lastDirection = 1
					lastNum = int(column)
					continue
				}
				fmt.Println("Unsafe, increasing and stopped on ", column)
				safe = false
				
			}
			if lastNum > int(column) {
				if lastNum - int(column) > 3 {
					fmt.Println("Unsafe, too big of a drop from", lastNum, "to", column)
					safe = false
					continue
				}
				//fmt.Println("Decreasing")
				if lastDirection == -1 || lastDirection == 0 {
					//fmt.Println("Direction is", lastDirection)
					lastDirection = -1
					lastNum = int(column)
					continue
				}
				fmt.Println("Unsafe, decreasing and stopped on ", column)
				safe = false
			}
			if lastNum == int(column) {
				fmt.Println("Unsafe, same number", lastNum, "and", column)
				safe = false
			}
		}
		safetyStatus = append(safetyStatus, safe)
		if safe {
			safeCount++
		}
	}
	return safeCount
}

func main() {
	fmt.Println("Computing safety of slope of list of reports")
	matrix := read_file("input.txt")
	if matrix == nil || len(matrix) == 0 {
		return
	}

	fmt.Printf("Read %d lines\n", len(matrix))
	fmt.Printf("Read %d columns\n", len(matrix[0]))

	safety := check_slope(matrix)
	fmt.Printf("Counted %d safe slopes\n", safety)
	
	//fmt.Println("Safety of slope is:", safety)
	
}
