package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// 1. sepal length in cm
	// 2. sepal width in cm
	// 3. petal length in cm
	// 4. petal width in cm
	// 5. class:
	// -- Iris Setosa
	// -- Iris Versicolour
	// -- Iris Virginica
	file, err := os.Open("datasets/iris.csv")
	check(err)

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	check(err)
	fmt.Println(len(records))
}
