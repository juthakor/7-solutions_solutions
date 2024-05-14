package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"unicode"
)

func fetchBeefData() (string, error) {
	url := "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text"
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func countBeefOccurrences(input string) map[string]int {
	words := make(map[string]int)
	currentWord := ""

	for _, char := range input {
		if unicode.IsLetter(char) || char == '-' {
			currentWord += string(unicode.ToLower(char))
		} else if currentWord != "" {
			words[currentWord] += 1
			currentWord = ""
		}
	}

	if currentWord != "" {
		words[currentWord] += 1
	}

	return words
}

func beefSummaryHandler(w http.ResponseWriter, r *http.Request) {
	data, err := fetchBeefData()
	if err != nil {
		http.Error(w, "Failed to fetch beef data", http.StatusInternalServerError)
		return
	}
	fmt.Println(data)

	response := countBeefOccurrences(data)
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to marshal JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func main() {
	http.HandleFunc("/beef/summary", beefSummaryHandler)

	err := http.ListenAndServe(":8085", nil)
	if err != nil {
		panic(err)
	}
}
