package ledger

import (
  "sync"
  "os"
)

type Disk struct {
  mu sync.Mutex
}

func NewDisk() Disk {
  return Disk{}
}

func (disk *Disk) CreatedDirs() {
    makeDir("storage")
    makeDir("storage/immutable")
    makeDir("storage/mutable")
    makeDir(BLOCKS)
    makeDir(TRANSACTIONS)
    makeDir(ACCOUNTS)
    makeDir(SNAPSHOTS)
}

func makeDir(location string) error {
  return os.Mkdir(location, 0755)
}

func (disk *Disk) Lock() {
  disk.mu.Lock()
}

func (disk *Disk) Unlock() {
  disk.mu.Unlock()
}

func (disk *Disk) Delete(filePath string) error {
  disk.Lock()
  defer disk.Unlock()
  return os.Remove(filePath)
}

func (disk *Disk) Write(filePath string, cborBytes []byte) error {
  disk.Lock()
  defer disk.Unlock()
  file, err := os.Create(filePath)
  defer file.Close()
  if err != nil { return err }
  _, err = file.Write(cborBytes)
  return err
}

func (disk *Disk) Read(filePath string) ([]byte, error) {
  disk.Lock()
  defer disk.Unlock()
  file, err := os.ReadFile(filePath)
  if err != nil { return nil, err }
  return file, nil
}

// BLOCKS ======================================================================

func (disk *Disk) SaveBlockCBOR(id uint, cborBytes []byte) error {
  return disk.Write(FormatBlockPath(id), cborBytes)
}

func (disk *Disk) GetBlockCBOR(id uint) ([]byte, error) {
  return disk.Read(FormatBlockPath(id))
}

func (disk *Disk) DeleteBlock(id uint) error {
  return disk.Delete(FormatBlockPath(id))
}

// TRANSACTIONS ================================================================

func (disk *Disk) SaveTxCBOR(hash string, cborBytes []byte) error {
  return disk.Write(FormatTxPath(hash), cborBytes)
}

func (disk *Disk) GetTxCBOR(hash string) ([]byte, error) {
  return disk.Read(FormatTxPath(hash))
}

func (disk *Disk) DeleteTx(hash string) error {
  return disk.Delete(FormatTxPath(hash))
}

// ACCOUNTS ====================================================================

func (disk *Disk) SaveAccountCBOR(addr string, cborBytes []byte) error {
  return disk.Write(FormatAccountPath(addr), cborBytes)
}

func (disk *Disk) GetAccountCBOR(addr string) ([]byte, error) {
  return disk.Read(FormatAccountPath(addr))
}

func (disk *Disk) DeleteAccount(addr string) error {
  return disk.Delete(FormatAccountPath(addr))
}

// SNAPSHOTS ===================================================================

func (disk *Disk) SaveSnapshotCBOR(id uint, cborBytes []byte) error {
  return disk.Write(FormatSnapshotPath(id), cborBytes)
}

func (disk *Disk) GetSnapshotCBOR(id uint) ([]byte, error) {
  return disk.Read(FormatSnapshotPath(id))
}

func (disk *Disk) DeleteSnapshot(id uint) error {
  return disk.Delete(FormatSnapshotPath(id))
}
