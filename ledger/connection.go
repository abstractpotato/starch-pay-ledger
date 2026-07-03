package ledger

type Connection struct {
  Disk Disk
}

func NewConnection() Connection {
  return Connection{
    Disk: NewDisk(),
  }
}

func (connection *Connection) SetUp() {
  connection.Disk.CreatedDirectories()
}
