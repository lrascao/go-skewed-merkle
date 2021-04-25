package merkle

import (
	"fmt"
	"testing"

	"github.com/tj/assert"
)

// func TestRootHashOnCreate(t *testing.T) {
// 	// create a root node with custom hash
// 	tree := New(Hash([]byte("foo")))
// 	// no way hash("foo") turns out to be all zeros
// 	if bytes.Equal(tree.Hash(), []byte{0, 0, 0, 0}) {
// 		t.Error("invalid root hash")
// 	}
// }

// func TestTreeHeight(t *testing.T) {
// 	tree := New(Hash([]byte("foo")))
// 	tree.Add([]byte("bar"))
// 	assert.Equal(t, 1, tree.Height())
// 	tree.Add([]byte("baz"))
// 	assert.Equal(t, 2, tree.Height())
// 	tree.Add([]byte("dog"))
// 	assert.Equal(t, 3, tree.Height())
// 	tree.Add([]byte("cat"))
// 	assert.Equal(t, 4, tree.Height())
// 	tree.Add([]byte("plant"))
// 	assert.Equal(t, 5, tree.Height())
// }

// func TestExistingProof(t *testing.T) {
// 	tree := New(Hash([]byte("foo")))
// 	tree.Add([]byte("bar"))
// 	tree.Add([]byte("baz"))
// 	tree.Add([]byte("dog"))
// 	tree.Add([]byte("cat"))
// 	tree.Add([]byte("plant"))

// 	// generate proof of existence of a bunch of
// 	// existing hashes
// 	proofFoo := tree.Proof(Hash([]byte("foo")))
// 	assert.Equal(t, 5, len(proofFoo))
// 	for height, hash := range proofFoo {
// 		fmt.Printf("#%d(%x) ", tree.Height()-height, hash)
// 	}
// 	fmt.Printf("\n")

// 	proofBar := tree.Proof(Hash([]byte("bar")))
// 	assert.Equal(t, 5, len(proofBar))
// 	for height, hash := range proofBar {
// 		fmt.Printf("#%d(%x) ", tree.Height()-height, hash)
// 	}
// 	fmt.Printf("\n")

// 	proofDog := tree.Proof(Hash([]byte("dog")))
// 	assert.Equal(t, 3, len(proofDog))
// 	for height, hash := range proofDog {
// 		fmt.Printf("#%d(%x) ", tree.Height()-height, hash)
// 	}
// 	fmt.Printf("\n")

// 	proofPlant := tree.Proof(Hash([]byte("plant")))
// 	assert.Equal(t, 1, len(proofPlant))
// 	for height, hash := range proofPlant {
// 		fmt.Printf("#%d(%x) ", tree.Height()-height, hash)
// 	}
// 	fmt.Printf("\n")
// }

// func TestNonExistingProof(t *testing.T) {
// 	tree := New(Hash([]byte("foo")))
// 	tree.Add([]byte("bar"))

// 	proofNotFound := tree.Proof(Hash([]byte("baz")))
// 	assert.Nil(t, proofNotFound)
// }

func TestVerifyProof(t *testing.T) {
	tree := New(Hash([]byte("foo")))
	tree.Add([]byte("bar"))
	tree.Add([]byte("baz"))

	hash := Hash([]byte("bar"))
	proofFoo := tree.Proof(hash)
	for height, hash := range proofFoo {
		fmt.Printf("#%d(%x) ", tree.Height()-height, hash)
	}
	fmt.Printf("\n")
	assert.Equal(t, true, tree.Verify(hash, proofFoo))
}
