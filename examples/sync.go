package main

import (
  "fmt"
  "github.com/abstractpotato/potato-serialization-lib"
  "github.com/abstractpotato/starch-pay-ledger"
)

func main() {
  block := genesisBlock()

  chain := ledger.NewLedger()
  valid_block := chain.ValidBlock(block)
  fmt.Printf("Valid Block: %v\n", valid_block)
}

func genesisBlock() psl.Block {
  return psl.NewBlock()
  // generate genesis block with initial params
}

func generateBlock() psl.Block {
  return psl.NewBlock()
  // generate block with transactions
}
