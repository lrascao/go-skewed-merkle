package skewed_merkle

import (
	"bytes"
	"testing"

	"github.com/tj/assert"
)

func TestRootHashOnCreate(t *testing.T) {
	// create a root node with custom hash
	tree := New(Hash([]byte("foo")))
	// no way hash("foo") turns out to be all zeros
	if bytes.Equal(tree.Hash(), []byte{0, 0, 0, 0}) {
		t.Error("invalid root hash")
	}
}

func TestTreeHeight(t *testing.T) {
	for _, ti := range []struct {
		leafs       []string
		want_height int
	}{
		{
			leafs:       []string{"bar"},
			want_height: 1,
		},
		{
			leafs:       []string{"bar", "baz"},
			want_height: 2,
		},
		{
			leafs:       []string{"bar", "baz", "dog"},
			want_height: 3,
		},
		{
			leafs:       []string{"bar", "baz", "dog", "cat", "plant"},
			want_height: 5,
		},
	} {
		t.Run("", func(t *testing.T) {

			tree := New(Hash([]byte("foo")))
			// add all the leafs
			for _, leaf := range ti.leafs {
				tree.Add([]byte(leaf))
			}
			assert.Equal(t, ti.want_height, tree.Height())
		})
	}
}

func TestNonExistingProof(t *testing.T) {
	tree := New(Hash([]byte("foo")))
	tree.Add([]byte("bar"))

	proof, err := tree.Proof(Hash([]byte("baz")))
	assert.Nil(t, proof)
	assert.IsType(t, err, NotFoundError{})
	assert.Equal(t, err.Error(), "not found")
}

func TestExistingProof(t *testing.T) {
	for _, ti := range []struct {
		leafs     []string
		leaf      string
		want_size int
	}{
		{
			leafs:     []string{"bar"},
			leaf:      "bar",
			want_size: 1,
		},
		{
			leafs:     []string{"bar"},
			leaf:      "foo",
			want_size: 1,
		},
		{
			leafs:     []string{"bar", "baz"},
			leaf:      "bar",
			want_size: 2,
		},
		{
			leafs:     []string{"bar", "baz"},
			leaf:      "baz",
			want_size: 1,
		},
		{
			leafs:     []string{"bar", "baz", "dog", "cat", "plant"},
			leaf:      "dog",
			want_size: 3,
		},
	} {
		t.Run("", func(t *testing.T) {

			tree := New(Hash([]byte("foo")))
			// add all the leafs
			for _, leaf := range ti.leafs {
				tree.Add([]byte(leaf))
			}
			var proof []Proof
			proof, err := tree.Proof(Hash([]byte(ti.leaf)))
			assert.Nil(t, err)
			assert.Equal(t, ti.want_size, len(proof))
		})
	}
}

func TestVerifyProof(t *testing.T) {
	for _, ti := range []struct {
		leafs []string
	}{
		{
			leafs: []string{"bar"},
		},
		{
			leafs: []string{"bar"},
		},
		{
			leafs: []string{"bar", "baz"},
		},
		{
			leafs: []string{"bar", "baz"},
		},
		{
			leafs: []string{"bar", "baz", "dog", "cat", "plant"},
		},
	} {
		t.Run("", func(t *testing.T) {

			tree := New(Hash([]byte("foo")))
			// add all the leafs
			for _, leaf := range ti.leafs {
				tree.Add([]byte(leaf))
			}

			// generate and verify a proof for every leaf
			for _, leaf := range ti.leafs {
				var proof []Proof
				proof, err := tree.Proof(Hash([]byte(leaf)))
				assert.Nil(t, err)
				assert.Equal(t, true, tree.Verify(Hash([]byte(leaf)), proof))
			}
		})
	}
}

func TestNonExistingVerify(t *testing.T) {
	tree := New(Hash([]byte("foo")))
	tree.Add([]byte("bar"))

	proof, err := tree.Proof(Hash([]byte("bar")))
	assert.Nil(t, err)
	assert.Equal(t, false, tree.Verify(Hash([]byte("baz")), proof))
}
