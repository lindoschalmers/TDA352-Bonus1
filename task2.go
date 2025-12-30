package main

import "fmt"

type Proof struct {
	proofNode [6]string
	leafIndex [2]int
}

func printProof(leaves *MerkleLeaves, name string, surname string) {
	fmt.Println(" ")
	fmt.Printf("Proof for %s %s\n\n", name, surname)
	//alice := Hash("AliceDoe-MerkleTree")
	firstName := Hash(name)
	lastName := Hash(surname)
	proof := &Proof{}

	for index, value := range leaves.Leaves {
		if value == firstName {
			fmt.Printf("Hash for %s found at index %d\n", name, index)

			proof.leafIndex[0] = index

			if index%2 == 0 {
				hash_name := fmt.Sprintf("H%d_%d", index, index+1)
				fmt.Printf("Leaf is even, parent leaf (%s) is hash of: %s and leaf: %d\n", hash_name, name, index+1)
				firstNameParent := fmt.Sprintf("%s%s", value, leaves.Leaves[index+1])
				firstNameParent = Hash(firstNameParent)
				fmt.Printf("Hash for %s is: %s\n", hash_name, firstNameParent)
			} else {
				hash_name := fmt.Sprintf("H%d_%d", index-1, index)
				fmt.Printf("Leaf is odd, parent leaf (%s) is hash of: %s and leaf: %d\n", hash_name, name, index-1)
				firstNameParent := fmt.Sprintf("%s%s", value, leaves.Leaves[index-1])
				firstNameParent = Hash(firstNameParent)
				fmt.Printf("Hash for %s is: %s\n", hash_name, firstNameParent)
			}
		}
		if value == lastName {
			fmt.Printf("Hash for %s found at index %d\n", surname, index)
			if index > proof.leafIndex[0] {
				proof.leafIndex[1] = proof.leafIndex[0]
				proof.leafIndex[0] = index
			}
			proof.leafIndex[1] = index

			if index%2 == 0 {
				lastNameParent := fmt.Sprintf("%s%s", value, leaves.Leaves[index+1])
				lastNameParent = Hash(lastNameParent)
				hash_name := fmt.Sprintf("H%d_%d", index, index+1)
				fmt.Printf("Leaf is even, parent leaf (%s) is hash of: %s and leaf: %d\n", hash_name, surname, index+1)
				fmt.Printf("Hash for %s is: %s\n", hash_name, lastNameParent)
			} else {
				hash_name := fmt.Sprintf("H%d_%d", index-1, index)
				fmt.Printf("Leaf is odd, parent leaf (%s) is hash of: %s and leaf: %d\n", hash_name, surname, index-1)
				lastNameParent := fmt.Sprintf("%s%s", value, leaves.Leaves[index-1])
				lastNameParent = Hash(lastNameParent)
				fmt.Printf("Hash %s is: %s\n", hash_name, lastNameParent)
			}
		}
	}

	if proof.leafIndex[0] < 8 && proof.leafIndex[1] < 8 { //Both on left side of tree
		if proof.leafIndex[0] < 4 && proof.leafIndex[1] < 4 { //Leftmost quarter
			if proof.leafIndex[0]+1 == proof.leafIndex[1] { //Same parent
				//Best case, need three nodes to prove.
			}
		}
		if proof.leafIndex[0] > 4 && proof.leafIndex[1] > 4 { //Second quarter
			if proof.leafIndex[0]+1 == proof.leafIndex[1] { //Same parent

			}
		}
		//Same half but different quarters:

	}

	//Worst case we need 6 nodes to prove. That is if it is completely different branches.
	//Best case we need 3 nodes to prove, that is if the leaves share a parent.
	//Same grandparent, but different parent, requires 4 nodes
	//Same great-great-grandparent requires 5 nodes
}
