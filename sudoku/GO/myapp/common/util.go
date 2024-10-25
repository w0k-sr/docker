package common

import (
	// "fmt"
	"strconv"
	"encoding/json"
	"strings"
)

// 文字列→配列(int型)
func ParsePostgresArray(input string) ([][]int, error) {
	input = strings.Trim(input, "{}")
	rows := strings.Split(input, "},{")
	var result [][]int
	for _, row := range rows {
		elements := strings.Split(row, ",")
		var intRow []int
		for _, element := range elements {
			num, err := strconv.Atoi(element)
			if err != nil {
				return nil, err
		}
			intRow = append(intRow, num)
		}
		result = append(result, intRow)
	}
	return result, nil
}

// func strToInt(s string) (int, error) {
    
   
// 	return strconv.Atoi(s)
// 	}
func ArrayToJson(arr [][]int) (string, error) {
	data, err := json.Marshal(arr)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func JsonToArray(j string) ([][]int, error) {
	var arr [][]int
	err := json.Unmarshal([]byte(j), &arr)
	return arr, err
}
