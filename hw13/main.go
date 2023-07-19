package main

import (
	"fmt"

	"github.com/ltdangle/blockchain"
)
func main() {
        // create a new b instance with a mining difficulty of 2
        b := blockchain.CreateBlockchain(2)

        // record transactions on the blockchain for Alice, Bob, and John
        b.AddBlock("Alice", "Bob", 5)
        b.AddBlock("John", "Bob", 2)

        // check if the blockchain is valid; expecting true
        fmt.Println(b.IsValid())
}
