package main

import (
	"fmt"
	"strconv"
)

func main() {
	//soal 1
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	fmt.Println(ArrayMerge([]string{"sergi", "jin"}, []string{"jin", "steve", "bryan"}))
	fmt.Println(ArrayMerge([]string{"alisa", "yoshimatu"}, []string{"devil jin", "yoshimatu", "alisa", "law"}))
	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	fmt.Println(ArrayMerge([]string{}, []string{}))

	// soal 2
	fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"}))
	fmt.Println(Mapping([]string{}))

	// // soal 3
	fmt.Println(munculSekali("1234123"))
	fmt.Println(munculSekali("76523752"))
	fmt.Println(munculSekali("12345"))
	fmt.Println(munculSekali("1122334455"))
	fmt.Println(munculSekali("0872504"))
}

func ArrayMerge(a, b []string) []string {
	uniqueNames := make(map[string]bool)
	for _, name := range a {
		uniqueNames[name] = true
	}
	for _, name := range b {
		if _, exists := uniqueNames[name]; !exists {
			uniqueNames[name] = true
		}
	}
	mergedArray := make([]string, 0)
	for name := range uniqueNames {
		mergedArray = append(mergedArray, name)
	}
	return mergedArray
}

func Mapping(slices []string) map[string]int {
	resaults := make(map[string]int)
	for _, itemSlice := range slices {
		if _, ok := resaults[itemSlice]; !ok {
			resaults[itemSlice] = 1
		} else {
			resaults[itemSlice]++
		}
	}
	return resaults
}

func munculSekali(angka string) []int {
	result := []int{}
	seen := map[int]bool{}
	for _, digit := range angka {
		digitInt, _ := strconv.Atoi(string(digit))
		if _, found := seen[digitInt]; !found {
			seen[digitInt] = true
			result = append(result, digitInt)
		}
	}
	return result
}
