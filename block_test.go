package blockchain

import (
	"testing"
)

func TestBlock(t *testing.T) {
	eh := "000d9676dc58cf09b7ecf21018d75ea3eccbb89206f977e288ccec3d8a62c0bd"
	en := 2093
	b := Block{
		Index:     0,
		Data:      "test string",
		Nonce:     0,
		PrevHash:  "",
		Timestamp: 1636387586233423900,
	}

	b.AssingHash()

	b.ProofOfWork(3)

	if b.Index != 0 {
		t.Errorf("Expected Index to be %v but instead got %v", 0, b.Index)
	}

	if b.Data != "test string" {
		t.Errorf("Expected data to be %v but instead got %v", "test string", b.Data)
	}

	if b.Hash != eh {
		t.Errorf("Expected hash to be %v but instead got %v", eh, b.Hash)
	}

	if b.Nonce != en {
		t.Errorf("Expected nonce to be %v but instead got %v", en, b.Nonce)

	}
}
