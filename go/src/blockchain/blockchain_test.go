package blockchain

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"encoding/hex"
)

// TODO: some useful tests of Blocks
func TestProof(t *testing.T) {
	b0 := Initial(16)
	b0.Mine(1)
	b1 := b0.Next("message")
	b1.Mine(1)

	assert.Equal(t, b0.Proof, uint64(56231))
	assert.Equal(t,b1.Proof, uint64(2159))

}

func TestCalcHash(t *testing.T) {
	b0 := Initial(16)
	b0.Mine(1)
	b1 := b0.Next("message")
	b1.Mine(1)	
	assert.Equal(t, hex.EncodeToString(b0.CalcHash()), "6c71ff02a08a22309b7dbbcee45d291d4ce955caa32031c50d941e3e9dbd0000")
	assert.Equal(t, hex.EncodeToString(b1.CalcHash()), "9b4417b36afa6d31c728eed7abc14dd84468fdb055d8f3cbe308b0179df40000")

}

func TestValidHash(t *testing.T) {
	b0 := Initial(16)
	assert.False(t, b0.ValidHash(), "validhash for initial(16) is wrong")
	b0.Proof = uint64(56231)
	assert.True(t, b0.ValidHash(), "validhash should return true after set proof")

	b1 := Initial(19)
	b1.SetProof(87745)
	b2 := b1.Next("hash example 1234")
	b2.SetProof(1407891)

	assert.True(t, b2.ValidHash())

	b2.SetProof(346082)
	assert.False(t, b2.ValidHash())
}

func TestMining(t *testing.T) {
	b0 := Initial(7)
	b0.Mine(1)
	assert.Equal(t, b0.Proof, uint64(385))
	assert.Equal(t, hex.EncodeToString(b0.Hash), "379bf2fb1a558872f09442a45e300e72f00f03f2c6f4dd29971f67ea4f3d5300")

	b1 := b0.Next("this is an interesting message")
	b1.Mine(1)
	assert.Equal(t, b1.Proof, uint64(20))
	assert.Equal(t, hex.EncodeToString(b1.Hash), "4a1c722d8021346fa2f440d7f0bbaa585e632f68fd20fed812fc944613b92500")

	b2 := b1.Next("this is not interesting")
	b2.Mine(1)
	assert.Equal(t, b2.Proof, uint64(40))
	assert.Equal(t, hex.EncodeToString(b2.Hash), "ba2f9bf0f9ec629db726f1a5fe7312eb76270459e3f5bfdc4e213df9e47cd380")

	b3 := Initial(20)
	b3.Mine(1)
	b4 := b3.Next("this is an interesting message")
	b4.Mine(1)
	b5 := b4.Next("this is not interesting")
	b5.Mine(1)
	assert.Equal(t, b3.Proof, uint64(1209938))
	assert.Equal(t, hex.EncodeToString(b3.Hash), "19e2d3b3f0e2ebda3891979d76f957a5d51e1ba0b43f4296d8fb37c470600000")
	assert.Equal(t, b4.Proof, uint64(989099))
	assert.Equal(t, hex.EncodeToString(b4.Hash), "a42b7e319ee2dee845f1eb842c31dac60a94c04432319638ec1b9f989d000000")
	assert.Equal(t, b5.Proof, uint64(1017262))
	assert.Equal(t, hex.EncodeToString(b5.Hash), "6c589f7a3d2df217fdb39cd969006bc8651a0a3251ffb50470cbc9a0e4d00000")

}

func testBlockchain(t *testing.T) {
	b0 := Initial(7)
	b0.Mine(1)
	b1 := b0.Next("this is an interesting message")
	b1.Mine(1)
	b2 := b1.Next("this is not interesting")
	b2.Mine(1)

	chain := Blockchain{}
	assert.False(t, chain.IsValid())
	chain.Add(b0)
	assert.True(t, chain.IsValid())
	chain.Add(b1)
	assert.True(t, chain.IsValid())
	chain.Add(b2)
	assert.True(t, chain.IsValid())
}