package blockchain

import (
	"crypto/sha256"
	"fmt"
	"strings"
	"time"
)

type Block struct {
	Index     int
	PrevHash  string
	Hash      string
	Data      string
	Timestamp int64
	Nonce     int
}

func NewBlock(data string) Block {
	b := Block{
		Data:      data,
		Timestamp: time.Now().UnixNano(),
	}

	b.AssingHash()
	return b
}

func (b Block) CalculateHash() string {
	h := sha256.New()
	d := fmt.Sprint(b.Index, b.PrevHash, b.Timestamp, b.Data, b.Nonce)
	h.Write([]byte(d))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func (b *Block) AssingHash() {
	b.Hash = b.CalculateHash()
}

func (b *Block) ProofOfWork(difficulty int) string {
	ds := strings.Repeat("0", difficulty)
	for b.Hash[:difficulty] != ds {
		b.Nonce += 1
		b.AssingHash()
	}
	return b.Hash
}
