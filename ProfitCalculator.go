package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)
const filename = "result.txt"

func profit_calculate() {

	revenue, err := getInput("Enter revenue = ")
	if (err != nil) {
		fmt.Print("Error = ", err)
		panic("Something went wrong")
	}
	expenses, err := getInput("Enter expenses = ")
	if (err != nil) {
		fmt.Print("Error = ", err)
		panic("Something went wrong")
	}
	taxRate, err := getInput("Enter tax rate = ")
	if (err != nil) {
		fmt.Print("Error = ", err)
		panic("Something went wrong")
	}

	ebt, profit, ratio := calc(revenue, expenses, taxRate)

	writeToFile(ebt, profit, ratio)
	readFromFile()
}

func getInput(text string)(float64, error) {
	var input float64;
	fmt.Print(text)
	fmt.Scan(&input)
	if (input < 0) {
		return 0, errors.New("Negative value is not allowed")
	}
	return input, nil
}

func calc(revenue, expenses, taxRate float64)(float64, float64, float64) {
	ebt := revenue-expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt/profit

	return ebt, profit, ratio
}

func writeToFile(ebt, profit, ratio float64) {
	data := fmt.Sprint(ebt, profit, ratio)
	os.WriteFile(filename, []byte(data), 0644)
}

func readFromFile() {
	byteData, _ := os.ReadFile(filename)
	data := string(byteData)
	arrayOfData := strings.Split(data, " ")
	fmt.Print(arrayOfData)
}
