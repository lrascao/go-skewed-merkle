package merkle

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

func New(hash []byte) Tree {
	return Tree{
		root: &Node{height: 0,
			hash:  hash[:],
			left:  nil,
			right: nil},
	}
}

func Hash(value []byte) []byte {
	h := sha256.Sum256(value)
	return h[:]
}

func IsLeaf(n *Node) bool {
	return n.left == nil && n.right == nil
}

func (t Tree) Hash() []byte {
	return t.root.hash
}

func (t Tree) Height() int {
	return t.root.height
}

func (n Node) Height() int {
	return n.height
}

func (n Node) Hash() []byte {
	return n.hash
}

func (t *Tree) Add(value []byte) {
	valueHash := sha256.Sum256(value)
	// find hash(rootHash + valueHash)
	nodeHash := sha256.New()
	nodeHash.Write(t.Hash())
	nodeHash.Write(valueHash[:])

	// notice how the left arm of this node will now point to the
	// previous tree root, this is where the skewing comes into effect
	n := &Node{height: t.root.Height() + 1,
		hash: nodeHash.Sum(nil),
		left: t.root,
		right: &Node{hash: valueHash[:],
			left:  nil,
			right: nil}}
	// this new node is now the tree root, it's left arm will point to
	// the old root
	t.root = n
}

func (t *Tree) Proof(hash []byte) [][]byte {
	// create a slice of byte slices with capacity at the height of the tree
	var acc0 [][]byte
	return proof(hash, t.root, acc0)
}

func (t *Tree) Verify(hash []byte, proof [][]byte) bool {
	if bytes.Equal(t.root.Hash(), verify(hash, t.root, proof)) {
		return true
	}
	return false
}

//
// Internal functions
//

func verify(hash []byte, n *Node, proof [][]byte) []byte {
	if IsLeaf(n) {
		return hash
	}

	// pop the top of the stack
	proof0, proof := proof[0], proof[1:]
	fmt.Printf("proof0: %x\n", proof0)

	return verify(hash, n.left, proof)
}

func proof(hash []byte, n *Node, acc0 [][]byte) [][]byte {
	// we got to the left leaf node, if this is not the
	// hash we're looking for then this means that it doesn't
	// exist in the tree. If it is then just return whatever
	// proof we've accumulated so far
	if IsLeaf(n) {
		if bytes.Equal(hash, n.Hash()) {
			return acc0
		}
		return nil
	}
	// not a leaf node

	// since this is a Skewed Merkle tree we can get
	// away at just peeking at the hash value of the right
	// hand side of every node
	if bytes.Equal(n.right.Hash(), hash) {
		// we found the provided hash, we can now
		// return the accumulated proof
		return append(acc0, n.left.Hash())
	}

	return proof(hash, n.left, append(acc0, n.right.Hash()))
}
