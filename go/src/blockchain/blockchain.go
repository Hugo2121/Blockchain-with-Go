package blockchain

import (
	"encoding/hex"
)

type Blockchain struct {
	Chain []Block
}

func (chain *Blockchain) Add(blk Block) {
	// You can remove the panic() here if you wish.
	if !blk.ValidHash() {
		panic("adding block with invalid hash")
	}
	// TODO
	chain.Chain = append(chain.Chain, blk)
}

func (chain Blockchain) IsValid() bool {
	// TODO
	if chain.Chain == nil {
		return false
	}

	firstBlock := chain.Chain[0]
	firstHash := [32]byte{}

	diff := firstBlock.Difficulty

	if firstBlock.Generation != 0 ||  hex.EncodeToString(firstBlock.PrevHash) != hex.EncodeToString(firstHash[:]) {
		return false
	}
	for geneNum, blk := range chain.Chain {
		if blk.Difficulty != diff {
			return false
		}
		if uint64(geneNum) != blk.Generation {
			return false
		}
		if geneNum < len(chain.Chain) - 1 {		
			if hex.EncodeToString(blk.Hash) != hex.EncodeToString(chain.Chain[geneNum+1].PrevHash) {
				return false
			}
		}

		if !blk.ValidHash() {
			return false
		}

		if hex.EncodeToString(blk.CalcHash()) != hex.EncodeToString(blk.Hash) {
			return false
		}
	}

	return true


}
