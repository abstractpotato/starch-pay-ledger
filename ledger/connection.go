package ledger

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

func GetTransaction() {
  // look in memory
  // if not found look in disk
  // load transaction into memory
}
