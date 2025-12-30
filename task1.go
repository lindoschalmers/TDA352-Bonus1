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

// MerkleLeaves holds 8 leaf nodes for the Merkle tree
type MerkleLeaves struct {
	Leaves  [16]string
	Parent1 [8]string
	Parent2 [4]string
	Parent3 [2]string
	Root    string
}

func MerkleTree(FirstName string, SurName string) *MerkleLeaves {
	// Generate hash from input name
	HashInput := fmt.Sprintf("%s%s-MerkleTree", FirstName, SurName)
	hashValue := Hash(HashInput)

	Number, err := firstElement(hashValue)
	if !err {
		fmt.Println("No numbers found in hash for: ", HashInput)
		fmt.Println("Hash: ", hashValue)
		os.Exit(1)
	}
	fmt.Println("Hash: ", hashValue)

	// Generate the hash for "AliceDoe-MerkleTree"
	aliceHash := Hash("AliceDoe-MerkleTree")

	leaves := &MerkleLeaves{}

	for i := 0; i < 16; i++ {
		leaves.Leaves[i] = aliceHash
	}

	for index, _ := range leaves.Leaves {
		if index == Number {
			leaves.Leaves[index] = Hash(FirstName)
		}
		if index == Number+1 {
			leaves.Leaves[index] = Hash(SurName)
		}
	}
	index_count := 0
	for index, _ := range leaves.Leaves {
		if index%2 == 0 {
			hash_string := fmt.Sprintf("%s%s", leaves.Leaves[index], leaves.Leaves[index+1])
			leaves.Parent1[index_count] = Hash(hash_string)
			fmt.Println("Parent 1: ", leaves.Parent1[index_count])
			index_count++
		}
	}
	index_count = 0
	for index, _ := range leaves.Parent1 {
		if index%2 == 0 {
			hash_string := fmt.Sprintf("%s%s", leaves.Parent1[index], leaves.Parent1[index+1])
			leaves.Parent2[index_count] = Hash(hash_string)
			fmt.Println("Parent 2: ", leaves.Parent2[index_count])
			index_count++
		}
	}

	index_count = 0
	for index, _ := range leaves.Parent2 {
		if index%2 == 0 {
			hash_string := fmt.Sprintf("%s%s", leaves.Parent2[index], leaves.Parent2[index+1])
			leaves.Parent3[index_count] = Hash(hash_string)
			fmt.Println("Parent3: ", leaves.Parent3[index_count])
			index_count++
		}
	}

	return leaves
}
