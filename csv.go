package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
)

func readCSVFromUrl(url string) ([][]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	reader := csv.NewReader(resp.Body)
	reader.Comma = ';'
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func main() {
    url := "https://raw.githubusercontent.com/NeoChow/quiz/master/personalIDs.csv"
	data, err := readCSVFromUrl(url)
	if err != nil {
		panic(err)
	}

	for _, row := range data {
		fmt.Println(row[0])
	}
}
