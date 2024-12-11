package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reportsList := findNumericValues()
	updatedList := [][]int{}

    for _, row := range reportsList {
		rowElements := []int{}
        for _, val := range row {
            rowElements = append(rowElements, val)
        }

		passes := findDifference(rowElements)
		ascDesPass := findAscOrDes(rowElements)
		
		if(passes && ascDesPass){
			updatedList = append(updatedList, row)
		}
	}
		
	fmt.Println("How many are safe: ", len(updatedList))
}

func findAscOrDes(list[] int ) bool {

	if(list[0] < list[1]){
		//check ascending
		for i:=1; i < len(list) - 1; i++{
			if(list[i] > list[i+1]){
				return false
			}
		}
	} else{
		for i:=1; i < len(list) - 1; i++{
			if(list[i] < list[i+1]){
				return false
			}
		}
	}
	return true
}


func findDifference(list[] int) bool {
	isSafe:= true

	for i:=0; i < len(list) - 1; i++ {
		difference := list[i] - list[i+1]
		positiveNum :=  math.Abs(float64(difference))

		if(positiveNum > 3 || positiveNum  < 1){

			return false
		}
	}
	 return isSafe
 }


func findNumericValues() ([][]int) {
	//initialize both lists to return
	list2D := [][]int{}
	list := []int{}

	file, err := os.Open("input.txt") //open file
	if err != nil {
		fmt.Println("Error opening file:", err)
	
	}
	
	defer file.Close()

	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {

		line := scanner.Text() //read line

		separator := " " // define separator
		
		currentLine := strings.Split(line, separator) //split string by separator 

		for _, value := range currentLine {
		    number, err := strconv.Atoi(value)
			if err != nil {
				fmt.Println("Error converting string to number ", value)   
 			}
			list = append(list, number)
		}		
		list2D = append(list2D, list)
		list = nil
}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return list2D
}