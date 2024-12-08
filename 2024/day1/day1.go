package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	
	differenceValuesList := [] int{}

	firstList, secondList := findNumericValues()

	//sort lists in ascending order
	sort.Slice(firstList, func(i, j int) bool {
		return firstList[i] < firstList[j]
	})
	sort.Slice(secondList, func(i, j int) bool {
		return secondList[i] < secondList[j]
	})

	for i := 0; i < len(firstList); i++ {

		differenceValue := calculateDifference(firstList[i], secondList[i])

		differenceValuesList = append(differenceValuesList, int(differenceValue))
	}

	totalDifference := sumValues(differenceValuesList)
	fmt.Println("TOTAL DIFFERENCE: ", totalDifference)
}

func sumValues(list []int) int {
	result := 0
	for _, value := range list {
		result = result + value
	}
	return result
}

func calculateDifference (a int, b int ) float64 {
	
	result := a - b
	positiveNum :=  math.Abs(float64(result))

	return positiveNum
}

func findNumericValues() ([]int, []int) {
	//initialize both lists to return
	list1 := []int{}
	list2 := [] int{}

	file, err := os.Open("input.txt") //open file
	if err != nil {
		fmt.Println("Error opening file:", err)
	
	}
	
	defer file.Close()

	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {

		line := scanner.Text() //read line

		separator := "   " // define separator
		
		currentLine := strings.Split(line, separator) //split string by seperator 


		for index, element := range currentLine {
		    number, err := strconv.Atoi(element)
			if err != nil {
				fmt.Println("Error converting string to number ", element)   
 			}

			if index == 0 {
        		list1 = append(list1, number)
		
    		} else {
        		list2 = append(list2, number)
    		}
		}		
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return list1, list2
}