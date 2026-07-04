package ledger

import (
  "sync"
  PSL "github.com/abstractpotato/potato-serialization-lib/psl"
)

type Memory struct {
  mu           sync.Mutex
  Self         string
  Peers        []string
  Context      PSL.Context
  Blocks       map[uint]PSL.Block
  Transactions map[string]PSL.Transaction
}

func NewMemory() Memory {
  return Memory{
    Blocks: make(map[uint]PSL.Block),
    Transactions: make(map[string]PSL.Transaction),
  }
}

func (memory *Memory) Lock() {
  memory.mu.Lock()
}

func (memory *Memory) Unlock() {
  memory.mu.Unlock()
}

func (memory *Memory) HasTx(hash string) bool {
  memory.Lock()
  defer memory.Unlock()
  _, ok := memory.Transactions[hash]
  return ok
}

func (memory *Memory) HasBlock(id uint) bool {
  memory.Lock()
  defer memory.Unlock()
  _, ok := memory.Blocks[id]
  return ok
}

func (memory *Memory) GetTx(hash string) (PSL.Transaction) {
  memory.Lock()
  defer memory.Unlock()
  return memory.Transactions[hash]
}

func (memory *Memory) GetBlock(id uint) (PSL.Block) {
  memory.Lock()
  defer memory.Unlock()
  return memory.Blocks[id]
}
