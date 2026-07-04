package ledger

import (
  "sync"
  "sort"
  PSL "github.com/abstractpotato/potato-serialization-lib/psl"
)

type Mempool struct {
  mu           sync.Mutex
  Transactions map[string]PSL.Transaction
}

func newMempool() Mempool {
  return Mempool{
    Transactions: make(map[string]PSL.Transaction),
  }
}

func (mempool *Mempool) Lock() {
  mempool.mu.Lock()
}

func (mempool *Mempool) Unlock() {
  mempool.mu.Unlock()
}

func (mempool *Mempool) GetSize() int {
  mempool.Lock()
  defer mempool.Unlock()
  return len(mempool.Transactions)
}

func (mempool *Mempool) HasTx(hash string) bool {
  mempool.Lock()
  defer mempool.Unlock()
  _, ok := mempool.Transactions[hash]
  return ok
}

func (mempool *Mempool) Add(tx *PSL.Transaction) {
  mempool.Lock()
  defer mempool.Unlock()
  _, ok := mempool.Transactions[tx.Header.Hash]
  if !ok {
    mempool.Transactions[tx.Header.Hash] = *tx
  }
}

func (mempool *Mempool) Del(hash string) {
  mempool.Lock()
  defer mempool.Unlock()
  delete(mempool.Transactions, hash)
}

func (mempool *Mempool) Clear() {
  mempool.Lock()
  defer mempool.Unlock()
  mempool.Transactions = make(map[string]PSL.Transaction)
}

func (mempool *Mempool) GetSorted() []PSL.Transaction {
  mempool.Lock()
  defer mempool.Unlock()

  mempoolSlice := make([]PSL.Transaction, 0)
  for _, mtx := range mempool.Transactions {
    mempoolSlice = append(mempoolSlice, mtx)
  }

  sort.Slice(mempoolSlice, func(i, j int) bool {
    return mempoolSlice[i].Body.Timestamp > mempoolSlice[j].Body.Timestamp
  })

  return mempoolSlice
}
