# go-skewed-merkle-tree

> A Skewed Merkle tree Go implementation

A Skewed Merkle tree is a special case of a classical [Merkle tree](https://en.wikipedia.org/wiki/Merkle_tree) in that it is not balanced and is skewed by a given certain to some side. This is the Go implementation of one
such tree, in this case a left leaning one. [Here](https://medium.com/codechain/skewed-merkle-tree-259b984acc0c) is an excellent writeup of why, in same cases, you're better off with a skewed tree instead of a balanced one.

Here's how you use it, first you create a tree and add some elements to it:

```go
    tree := skewed_merkle.New([]byte("foo"))

    tree.Add([]byte("bar"))
    tree.Add([]byte("baz"))
```

Now generate a proof that some element resides in the tree:

```go
    proof, err := tree.Proof(Hash([]byte("baz")))
```

This proof can now be independently verified by someone in possession of the tree:

```go
	tree.Verify([]byte("baz"), proof))
```
