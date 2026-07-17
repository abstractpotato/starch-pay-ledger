package ledger

import (
  "strconv"
  "path/filepath"
)

const BLOCKS       = "storage/immutable/blocks"
const TRANSACTIONS = "storage/immutable/transactions"

const ACCOUNTS     = "storage/mutable/accounts"
const SNAPSHOTS    = "storage/mutable/snapshots"

func FormatBlockPath(id uint) string {
  blockId := strconv.FormatUint(uint64(id), 10)
  return filepath.Join(BLOCKS, blockId)
}

func FormatTxPath(hash string) string {
  return filepath.Join(TRANSACTIONS, hash)
}

func FormatAccountPath(addr string) string {
  return filepath.Join(ACCOUNTS, addr)
}

func FormatSnapshotPath(id uint) string {
  snapshotId := strconv.FormatUint(uint64(id), 10)
  return filepath.Join(SNAPSHOTS, snapshotId)
}
