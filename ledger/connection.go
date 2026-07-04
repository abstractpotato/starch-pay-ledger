package ledger

import (
  PSL "github.com/abstractpotato/potato-serialization-lib/psl"
  "path/filepath"
)


type Connection struct {
  Disk   Disk
  Memory Memory
}

func NewConnection() Connection {
  return Connection{
    Disk: NewDisk(),
    Memory: NewMemory(),
  }
}

func (connection *Connection) SetUp() {
  connection.Disk.CreatedDirectories()
}

func (connection *Connection) GetTx(hash string) (PSL.Transaction, error) {
  if connection.Memory.HasTx(hash) {
    return connection.Memory.GetTx(hash)
  }
  filePath := filepath.Join("immutable/transactions/", hash)
  cborBytes, err := Disk.Read(filePath)
  if err != nil { PSL.Transaction{}, err }
  return PSL.TransactionFromCBOR(cborBytes)
}

// func (connection *Connection) GetBlock(id uint) (PSL.Block, error) {}
