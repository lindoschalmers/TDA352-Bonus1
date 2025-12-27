package main

import (
	"fmt"
	"os"
	"strconv"
)

func firstNumber(hash string) (int, bool) {
	for _, char := range hash {
		if char >= '0' && char <= '9' {
			num, _ := strconv.Atoi(string(char))
			return num, true
		}
	}
	return 0, false
}

func MerkleTree(FirstName string, SurName string) (int, int) {
	HashInput := fmt.Sprintf("%s%s-MerkleTree", FirstName, SurName)
	Hash := Hash(HashInput)
	Number, err := firstNumber(Hash)
	if !err {
		fmt.Println("No numbers found in hash for: ", HashInput)
		fmt.Println("Hash: ", Hash)
		os.Exit(1)
	}
	return Number, Number + 1

}
