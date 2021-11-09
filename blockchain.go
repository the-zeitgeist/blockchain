package blockchain

import (
	"encoding/json"
	"io/ioutil"
)

type Blockchain struct {
	Chain      []Block
	Difficulty int
}

type BlockChainOption func(*Blockchain)

func NewBlockchain(fn ...BlockChainOption) Blockchain {
	bc := &Blockchain{}

	for _, f := range fn {
		f(bc)
	}

	bc.AddBlock(NewBlock("Initial block"))
	return *bc
}

func (bc *Blockchain) AddBlock(b Block) {
	b.Index = len(bc.Chain)
	if b.Index != 0 {
		b.PrevHash = bc.GetLastBlock().Hash
	}
	b.ProofOfWork(bc.Difficulty)
	bc.Chain = append(bc.Chain, b)
}

func (bc Blockchain) GetLastBlock() Block {
	return bc.Chain[len(bc.Chain)-1]
}

func (bc Blockchain) String() string {
	file, _ := json.MarshalIndent(bc.Chain, "", " ")
	return string(file)
}

func WithDifficulty(d int) BlockChainOption {
	return func(bc *Blockchain) {
		bc.Difficulty = d
	}
}

func (bc Blockchain) ValidateIntegrity() bool {
	i, isValid := 0, true
	for isValid && i < len(bc.Chain) {
		b := bc.Chain[i]

		if b.Hash != b.CalculateHash() {
			isValid = false
		}

		if i != len(bc.Chain)-1 {
			if b.Hash != bc.Chain[i+1].PrevHash {
				isValid = false
			}
		}

		i++
	}

	return isValid
}

func (bc Blockchain) SaveToJSON(filename string) error {
	if filename == "" {
		filename = "blockchain.json"
	}

	file, err := json.MarshalIndent(bc.Chain, "", " ")

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, file, 0666)

	if err != nil {
		return err
	}

	return nil
}
