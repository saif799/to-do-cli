package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gocarina/gocsv"
)

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

func Write(Description string) {

	file, err := os.OpenFile("data.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

	lastId := GetLastLineID()

	now := time.Now()
	formated := now.Format(time.RFC3339)

	writer := csv.NewWriter(file)
	defer writer.Flush()

	ID := strconv.FormatInt(int64(lastId+1), 10)

	data := []string{ID, Description, formated, "false"}

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

func Read(all bool) {
	file, err := os.Open("data.csv")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	var records []Record

	if err := gocsv.UnmarshalFile(file, &records); err != nil {
		panic(err)
	}

	for _, record := range records {

		if all {
			fmt.Printf("ID: %s, Description: %s, IsComplete: %s \n", record.ID, record.Description, record.IsComplete)
			continue
		}

		if isdone, err := strconv.ParseBool(record.IsComplete); err != nil {
			panic(err)
		} else if !isdone {
			fmt.Printf("ID: %s, Description: %s, IsComplete: %s \n", record.ID, record.Description, record.IsComplete)
		}

	}

	defer file.Close()
}

func GetLastLineID() int {

	file, err := os.Open("data.csv")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	var records []Record

	if err := gocsv.UnmarshalFile(file, &records); err != nil {
		panic(err)
	}

	if len(records) == 0 {
		return 0
	}

	lastId, err := strconv.Atoi(records[len(records)-1].ID)

	if err != nil {
		panic(err)
	}

	return lastId

}

// func Complete(target string) {

// 	ID := strconv.Atoi(target)

// }
