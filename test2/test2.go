package main

import "fmt"

func main() {
	testArray := []string{"LLRR=", "==RLL", "=LLRR", "RRL=R", "RRLLR=RLL", "==LRRLLR=L=R"}
	for i := 0; i < len(testArray); i++ {
		leastSumDecoder(testArray[i])
	}
}

func leastSumDecoder(inputString string) {
	outputStack := make([]int, len(inputString)+1)

	for i := 0; i < len(inputString); {
		switch inputString[i] {
		case '=':
			if outputStack[i] != outputStack[i+1] {
				if outputStack[i+1] == 0 {
					outputStack[i+1] = outputStack[i]
					i++
				} else {
					outputStack[i] = outputStack[i+1]
					if i != 0 {
						i--
					}
				}
			} else {
				i++
			}

		case 'L':
			if outputStack[i] <= outputStack[i+1] {
				outputStack[i] = outputStack[i+1] + 1
				if i != 0 {
					i--
				}
			} else {
				i++
			}
		case 'R':
			if outputStack[i] >= outputStack[i+1] {
				outputStack[i+1] = outputStack[i] + 1
			}
			i++
		}
	}

	fmt.Println(outputStack)
}
