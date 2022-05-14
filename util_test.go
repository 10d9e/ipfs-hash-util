package util

import (
	"testing"

	u "github.com/ipfs/go-ipfs-util"
)

func TestSmallSum(t *testing.T) {
	digest, err := Sum([]byte("hello world"))
	if err != nil {
		t.Fatal(err)
	}
	expected := "bafkqac3imvwgy3zao5xxe3de"
	digestStr := EncodeToString(digest)

	if digestStr != expected {
		t.Fatalf("Incorrect multihash provided: %s, expected: %s", digestStr, expected)
	}
}

func TestLargeSum(t *testing.T) {
	contents := make([]byte, 10*1024*1024)
	u.NewSeededRand(0xdeadbeef).Read(contents)

	digest, err := Sum(contents)
	if err != nil {
		t.Fatal(err)
	}
	expected := "bafybeiby3w3wmno2roidva2cax4vo4mzbvaq2e332pn5rnugbv5elxnnue"
	digestStr := EncodeToString(digest)

	if digestStr != expected {
		t.Fatalf("Incorrect multihash provided: %s, expected: %s", digestStr, expected)
	}
}
