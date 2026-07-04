package ledger

import (
  PSL "github.com/abstractpotato/potato-serialization-lib/psl"
  "path/filepath"
  "strconv"
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
    return connection.Memory.GetTx(hash), nil
  }

  filePath := filepath.Join("immutable/transactions/", hash)
  cborBytes, err := connection.Disk.Read(filePath)
  if err != nil { return PSL.NewTransaction(), err }
  return PSL.TransactionFromCBOR(cborBytes)
}

func (connection *Connection) GetBlock(id uint) (PSL.Block, error) {
  if connection.Memory.HasBlock(id) {
    return connection.Memory.GetBlock(id), nil
  }

  blockId := strconv.FormatUint(uint64(id), 10)
  filePath := filepath.Join("immutable/blocks/", blockId)
  cborBytes, err := connection.Disk.Read(filePath)
  if err != nil { return PSL.NewBlock(), err }
  return PSL.BlockFromCBOR(cborBytes)
}

func (connection *Connection) LoadBlocksFromDrive(){}
