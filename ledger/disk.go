package ledger

import (
  "os"
  // "errors"
  // "strconv"
  "path/filepath"
)

type Disk struct {
  RootDir      string
  ImmutableDir string
  MutableDir   string
}

func NewDisk() Disk {
  return Disk{
    RootDir: "storage",
    ImmutableDir: "immutable",
    MutableDir: "mutable",
  }
}

func (disk *Disk) CreatedDirectories() {
    makeDir(disk.RootDir)

    immutableDir := filepath.Join(disk.RootDir, disk.ImmutableDir)
    mutableDir := filepath.Join(disk.RootDir, disk.MutableDir)

    // Immutable Directory
    makeDir(immutableDir)
    makeDir(filepath.Join(immutableDir, "params"))
    makeDir(filepath.Join(immutableDir, "epoch"))
    makeDir(filepath.Join(immutableDir, "blocks"))
    makeDir(filepath.Join(immutableDir, "transactions"))
    makeDir(filepath.Join(immutableDir, "validators"))
    makeDir(filepath.Join(immutableDir, "certificates"))
    makeDir(filepath.Join(immutableDir, "requests"))

    // Mutable Directory
    makeDir(mutableDir)
    makeDir(filepath.Join(mutableDir, "accounts"))
    makeDir(filepath.Join(mutableDir, "snapshots"))
}

func makeDir(location string) error {
  return os.Mkdir(location, 0755)
}

func (disk *Disk) Delete(filePath string) error {
  location := filepath.Join(disk.RootDir, filePath)
  return os.Remove(location)
}

func (disk *Disk) Write(filePath string, content []byte) error {
  location := filepath.Join(disk.RootDir, filePath)
  file, err := os.Create(location)
  defer file.Close()
  if err != nil { return err }
  _, err = file.Write(content)
  return err
}

func (disk *Disk) Read(filePath string) ([]byte, error) {
  location := filepath.Join(disk.RootDir, filePath)
  file, err := os.ReadFile(location)
  if err != nil { return nil, err }
  return file, nil
}
