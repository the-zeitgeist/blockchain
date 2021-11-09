package blockchain

import (
	"testing"
)

func TestBlockChain(t *testing.T) {
	bc := NewBlockchain(WithDifficulty(3))

	if len(bc.Chain) != 1 {
		t.Errorf("Expected initial length to be %v but instead got %v", 1, len(bc.Chain))
	}

	if bc.Difficulty != 3 {
		t.Errorf("Expected difficulty to be %v but instead got %v", 3, bc.Difficulty)
	}

	nb := Block{
		Index:     0,
		Data:      "test string",
		Nonce:     0,
		PrevHash:  "",
		Timestamp: 1636387586233423900,
	}

	nb.AssingHash()

	bc.AddBlock(nb)

	if bc.Chain[1].Index != 1 {
		t.Errorf("Expected automated index to be %v but instead got %v", 1, bc.Chain[1].Index)
	}

	if len(bc.Chain) != 2 {
		t.Errorf("Expected new length to be %v but instead got %v", 2, len(bc.Chain))
	}

	if bc.GetLastBlock().Timestamp != nb.Timestamp {
		t.Errorf("Expected last block to be %+v but instead got %+v", nb, bc.GetLastBlock())
	}

	if bc.Chain[0].Hash != bc.Chain[1].PrevHash {
		t.Errorf("Expected last block hash to be %+v but instead got %+v", bc.Chain[0].Hash, bc.Chain[1].PrevHash)
	}

	isValid := bc.ValidateIntegrity()

	if !isValid {
		t.Error("Expected chain to be valid")
	}

	bc.Chain[len(bc.Chain)-1].Data = "test string1"

	isValid = bc.ValidateIntegrity()

	if isValid {
		t.Error("Expected chain to be invalid")
	}
}
