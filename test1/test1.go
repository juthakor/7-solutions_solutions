package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://raw.githubusercontent.com/7-solutions/backend-challenge/main/files/hard.json"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	var dataArray [][]int

	err = json.Unmarshal(body, &dataArray)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	fmt.Println(findMaxSumPath(dataArray))
}

func findMaxSumPath(inputArray [][]int) int {
	//Choose maximum child node, then add it to parent from bottom to top of the tree
	for layerIndex := len(inputArray) - 1; layerIndex > 0; layerIndex-- {
		for rowIndex := 0; rowIndex < len(inputArray[layerIndex])-1; rowIndex++ {
			if inputArray[layerIndex][rowIndex] > inputArray[layerIndex][rowIndex+1] {
				inputArray[layerIndex-1][rowIndex] += inputArray[layerIndex][rowIndex]
			} else {
				inputArray[layerIndex-1][rowIndex] += inputArray[layerIndex][rowIndex+1]
			}
		}
	}

	return inputArray[0][0]
}
