package main

import (
	"fmt"
	"os"
)

func firstElement(hash string) (int, bool) {
	if len(hash) == 0 {
		return 0, false
	}

	char := rune(hash[0]) // Get first character

	// Check if it's a valid hex character
	if char >= '0' && char <= '9' {
		return int(char - '0'), true // 0-9
	} else if char >= 'a' && char <= 'f' {
		return int(char - 'a' + 10), true // 10-15
	}
	return 0, false
}

func MerkleTree(FirstName string, SurName string) (int, int) {
	HashInput := fmt.Sprintf("%s%s-MerkleTree", FirstName, SurName)
	Hash := Hash(HashInput)
	Number, err := firstElement(Hash)
	if !err {
		fmt.Println("No numbers found in hash for: ", HashInput)
		fmt.Println("Hash: ", Hash)
		os.Exit(1)
	}
	fmt.Println("Hash: ", Hash)
	return Number, Number + 1

}
