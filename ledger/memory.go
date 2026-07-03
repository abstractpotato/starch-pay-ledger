package ledger

import (
  "sync"
  PSL "github.com/abstractpotato/potato-serialization-lib/psl"
)

type Memory struct {
  Lock         sync.Mutex
  Self         string
  Peers        []string
  Context      PSL.Context
  Blocks       map[uint]PSL.Block
  Transactions map[string]PSL.Transaction
}

func NewMemory() Memory {
  return Memory{
    Blocks: make(map[uint]PSL.Block, 0),
    Transactions: make(map[string]PSL.Transaction, 0),
  }
}
