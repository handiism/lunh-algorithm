package main

import (
	"fmt"
	"strconv"
)

func isValid(number uint64) bool {
	return prefixMatched(getPrefix(number, 2), 4) && (sumOfDoubleEvenPlace(number)+sumOfOddPlace(number))%10 == 0 && getSize(number) >= 13 && getSize(number) <= 16
}

func sumOfDoubleEvenPlace(number uint64) int {
	var rtn int
	numberAsString := reversedString(number)
	for i := 0; i < len(numberAsString); i++ {
		if i%2 == 1 {
			dump, _ := strconv.Atoi(string(numberAsString[i]))
			rtn += getDigit(dump * 2)
		}
	}
	return rtn
}

func getDigit(number int) int {
	if len(strconv.Itoa(number)) == 1 {
		return number
	} else {
		numberAsString := strconv.Itoa(number)
		number1, _ := strconv.Atoi(string(numberAsString[0]))
		number2, _ := strconv.Atoi(string(numberAsString[1]))
		return number1 + number2
	}
}

func reversedString(number uint64) string {
	var rtn string
	numberAsString := strconv.FormatUint(number, 10)
	for i := len(numberAsString) - 1; i >= 0; i-- {
		rtn += string(numberAsString[i])
	}
	return rtn
}

func sumOfOddPlace(number uint64) int {
	var rtn int
	numberAsString := reversedString(number)
	for i := 0; i < len(numberAsString); i++ {
		if i%2 == 0 {
			dump, _ := strconv.Atoi(string(numberAsString[i]))
			rtn += dump
		}
	}
	return rtn
}

func prefixMatched(number uint64, d int) bool {
	return int(number) == d
}

func getSize(d uint64) int {
	return len(strconv.FormatUint(d, 10))
}

func getPrefix(number uint64, k int) uint64 {
	numberAsString := strconv.FormatUint(number, 10)
	kAsString := strconv.Itoa(k)
	if len(kAsString) > len(numberAsString) {
		return number
	} else {
		dump := []rune(numberAsString)
		rtn, _ := strconv.ParseUint(string(dump[0:len(kAsString)]), 10, 64)
		return rtn
	}
}

func main() {
	var number uint64
	fmt.Println("Masukkan nomor pada kartu kredit:")
	_, _ = fmt.Scanln(&number)
	if isValid(number) {
		fmt.Println("Nomor kartu", number, "valid.")
	} else {
		fmt.Println("Nomor kartu", number, "tidak valid.")
	}
}
