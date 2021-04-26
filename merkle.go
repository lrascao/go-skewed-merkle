package skewed_merkle

import (
	"bytes"
	"crypto/sha256"
)

// New creates a new Skewed Merkle tree
func New(hash []byte) Tree {
	return Tree{
		root: &Node{height: 0,
			hash:  hash[:],
			left:  nil,
			right: nil},
	}
}

// Hash creates a SHA256 hash of the provided value
func Hash(value []byte) []byte {
	h := sha256.Sum256(value)
	return h[:]
}

// Hash returns the tree root hash
func (t Tree) Hash() []byte {
	return t.root.hash
}

// Height returns the height of the tree
func (t Tree) Height() int {
	return t.root.height
}

// Height returns the height that the node is at
func (n Node) Height() int {
	return n.height
}

// Hash returns the node's hash
func (n Node) Hash() []byte {
	return n.hash
}

// Add adds a new value to the tree
func (t *Tree) Add(value []byte) {
	valueHash := sha256.Sum256(value)

	// notice how the left arm of this node will now point to the
	// previous tree root, this is where the skewing happens
	n := &Node{height: t.root.Height() + 1,
		// find hash(rootHash + valueHash)
		hash: hash(t.Hash(), valueHash[:]),
		left: t.root,
		right: &Node{hash: valueHash[:],
			left:  nil,
			right: nil}}
	// this new node is now the tree root, it's left arm will point to
	// the old root
	t.root = n
}

func (t *Tree) Proof(hash []byte) ([]Proof, error) {
	// create a slice of proofs
	var acc0 []Proof
	return proof(hash, t.root, acc0)
}

func (t *Tree) Verify(hash []byte, proof []Proof) bool {
	if bytes.Equal(t.root.Hash(), verify(hash, proof)) {
		return true
	}
	return false
}

//
// Internal functions
//

// verify returns the hashed proof with the provided hash
func verify(valueHash []byte, proof []Proof) []byte {
	// we've gone through the whole proof slice, ie.
	// we're at the tree leaf now, return the valuehash
	// so the recursion can unwind
	if len(proof) == 0 {
		return valueHash
	}
	// pop the top of the proof stack
	proof0, proof := proof[0], proof[1:]

	// if this proof entry came from a left hand leaf
	// then we need to hash(L, R) to keep the correct order
	if proof0.side == Left {
		return hash(proof0.hash, verify(valueHash, proof))
	}
	// normal case, proof came from a right hand side leaf
	return hash(verify(valueHash, proof), proof0.hash)
}

// proof returns a slice containing a proof of existence
// of the provided hash in the tree
func proof(hash []byte, n *Node, acc0 []Proof) ([]Proof, error) {
	// we got to the left leaf node, if this is not the
	// hash we're looking for then this means that it doesn't
	// exist in the tree. If it is then just return whatever
	// proof we've accumulated so far
	if isLeaf(n) {
		if bytes.Equal(hash, n.Hash()) {
			return acc0, nil
		}
		return nil, NotFoundError{}
	}
	// not a leaf node

	// since this is a Skewed Merkle tree we can get
	// away at just peeking at the hash value of the right
	// hand side of every node
	if bytes.Equal(n.right.Hash(), hash) {
		// found the provided hash, append the hash
		// of the co-node and return
		return append(acc0, Proof{side: Left,
			hash: n.left.Hash()}), nil
	}

	// continue searching on the left hand side of the tree
	return proof(hash, n.left, append(acc0, Proof{side: Right,
		hash: n.right.Hash()}))
}

// hash hashes two concatenated byte slices
func hash(h1 []byte, h2 []byte) []byte {
	h := sha256.New()
	h.Write(h1)
	h.Write(h2)
	r := h.Sum(nil)
	return r
}

// isLeaf returns a boolean wether if it it's a leaf node or not
func isLeaf(n *Node) bool {
	return n.left == nil && n.right == nil
}
