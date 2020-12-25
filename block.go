package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// Block es el estructura de cada bloque
type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

// SetHash genera el hash de cada bloque
func (b *Block) SetHash() {
	// convertir el timestamp en slice
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))

	// concatenar todos los slice
	// headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp, []byte{}})
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	// pasando contenido del slice
	b.Hash = hash[:]
}

// NewBlock funcion crear bloque:
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		Timestamp:     time.Now().Unix(),
		Data:          []byte(data),
		PrevBlockHash: prevBlockHash,
		Hash:          []byte{},
	}
	block.SetHash()
	return block
}

func main() {}
