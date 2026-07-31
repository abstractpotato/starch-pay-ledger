package ledger

import (
  "math"
  "errors"
  "github.com/abstractpotato/potato-serialization-lib"
)

type Ledger struct {
  Disk         Disk
  Tip          uint
  PreviousHash string
  InitTime     uint
  GenesisSeed  []byte
  Params       psl.Params
  Requests     []psl.Request
  Certificates []psl.Certificate
  Accounts     Accounts
  Mempool      Mempool
}

func NewLedger() Ledger {
  return Ledger{
    Requests: make([]psl.Request, 0),
    Certificates: make([]psl.Certificate, 0),
    Accounts: NewAccounts,
    Mempool: NewMempool(),
    Disk: NewDisk(),
  }
}

func (ledger *Ledger) AddGenesis(block psl.Block) {
  ledger.InitTime = block.Body.Timestamp
  ledger.Genesis = *block.Body.Genesis
  ledger.Params = block.Body.Genesis.Params
  ledger.AddCertificate(block.Body.Genesis.Certificate)
}

func (ledger *Ledger) ValidBlock(block psl.Block) bool {
  if ledger.Tip == 0 && block.Body.ID == 0 { return true }
  if ledger.Tip + 1 != block.Body.ID { return false }
  return ledger.PreviousHash == block.Body.PreviousHash
}

func (ledger *Ledger) Update(TempState) {
  // update ledger with TempState
}

func (ledger *Ledger) AddRequest(request psl.Request) {
  ledger.Requests = append(ledger.Requests, request)
}

func (ledger *Ledger) AddCertificate(certificate psl.Certificate) {
  ledger.Certificates = append(ledger.Certificates, certificate)
}
