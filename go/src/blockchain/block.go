package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	// "math"
)

type Block struct {
	PrevHash   []byte
	Generation uint64
	Difficulty uint8
	Data       string
	Proof      uint64
	Hash       []byte
}

// Create new initial (generation 0) block.
func Initial(difficulty uint8) Block {
	// TODO
	prehash := [32]byte{}

	b := Block {PrevHash: prehash[:], 
				Generation: 0, 
				Difficulty: difficulty, 
				Data: ""}
	return b
}

// Create new block to follow this block, with provided data.
func (prev_block Block) Next(data string) Block {
	// TODO
	b := Block {Generation: prev_block.Generation + 1, 
				Difficulty: prev_block.Difficulty,
				Data: data,
				PrevHash: prev_block.Hash}

	return b
}

// Calculate the block's hash.
func (blk Block) CalcHash() []byte {
	// TODO
	hashString := hex.EncodeToString(blk.PrevHash) + ":" +
				  fmt.Sprint(blk.Generation) + ":" +
				  fmt.Sprint(blk.Difficulty) + ":" +
				  blk.Data + ":" +
				  fmt.Sprint(blk.Proof)

	// fmt.Println(hashString)

	ret := sha256.Sum256([]byte(hashString))

	return ret[:]


}

// Is this block's hash valid?
func (blk Block) ValidHash() bool {
	// TODO
	nBytes := blk.Difficulty / 8
	nBits := blk.Difficulty % 8


	hashString := blk.CalcHash()

	// fmt.Println(hex.EncodeToString(hashString))

	for i := len(hashString) - 1; i >= len(hashString) - int(nBytes); i-- {
		if hashString[i] != 0 {
			return false
		} 
	}

	if hashString[len(hashString) - int(nBytes) - 1] % (1<<nBits) != 0 {
		return false
	}
	return true
}

// Set the proof-of-work and calculate the block's "true" hash.
func (blk *Block) SetProof(proof uint64) {
	blk.Proof = proof
	blk.Hash = blk.CalcHash()
}
