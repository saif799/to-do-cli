package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/gocarina/gocsv"
)

func Add(first string, second string) (result string) {
	num1, err := strconv.ParseFloat(first, 64)
	if err != nil {
		fmt.Println("Error: First value is invalid")
		return
	}
	num2, err := strconv.ParseFloat(second, 64)
	if err != nil {
		fmt.Println("Error: Second value is invalid")
		return
	}
	return fmt.Sprintf("%f", num1+num2)
}

func Subtract(from string, subtract string) (result string) {
	num1, err := strconv.ParseFloat(from, 64)
	if err != nil {
		fmt.Println("Error: First value is invalid")
		return
	}
	num2, err := strconv.ParseFloat(subtract, 64)
	if err != nil {
		fmt.Println("Error: Second value is invalid")
		return
	}
	return fmt.Sprintf("%f", num1-num2)
}

func Multiply(first string, second string, shouldRoundUp bool) (result string) {
	num1, err := strconv.ParseFloat(first, 64)
	if err != nil {
		fmt.Println("Error: First value is not a decimal")
		return
	}
	num2, err := strconv.ParseFloat(second, 64)
	if err != nil {
		fmt.Println("Error: Second value is not a decimal")
		return
	}
	if shouldRoundUp {
		return fmt.Sprintf("%.2f", num1*num2)
	}
	return fmt.Sprintf("%f", num1*num2)
}

func Divide(divide string, by string, shouldRoundUp bool) (result string, e error) {
	num1, err := strconv.ParseFloat(divide, 64)
	if err != nil {
		return "", fmt.Errorf("first value is not a number")
	}
	num2, err := strconv.ParseFloat(by, 64)
	if err != nil {
		return "", fmt.Errorf("second value is not a number")
	}
	if shouldRoundUp {
		return fmt.Sprintf("%.2f", num1/num2), nil
	}
	return fmt.Sprintf("%f", num1/num2), nil
}

func Write(ID, Description, CreatedAt, IsComplete string) {

	file, err := os.OpenFile("data.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()
	data := []string{ID, Description, CreatedAt, IsComplete}

	errr := writer.Write(data)
	if errr != nil {
		panic("this shit broke")
	}

}

type Record struct {
	ID          string `csv:"ID"`
	Description string `csv:"Description"`
	CreatedAt   string `csv:"CreatedAt"`
	IsComplete  string `csv:"IsComplete"`
}

func Read() {
	file, err := os.Open("data.csv")

	if err != nil {
		panic(err)
	}

	var records []Record

	if err := gocsv.UnmarshalFile(file, &records); err != nil {
		panic(err)
	}

	for _, record := range records {
		fmt.Printf("ID: %s, Description: %s, IsComplete: %s \n", record.ID, record.Description, record.IsComplete)
	}

	defer file.Close()
}
