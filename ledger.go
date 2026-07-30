package ledger

import (
  "math"
  "errors"
  "encoding/hex"
  "encoding/json"
  "github.com/fxamacker/cbor/v2"
  "github.com/abstractpotato/potato-serialization-lib"
)

type Ledger struct {
  Disk         Disk              `cbor:"-" json:"-"`
  Tip          uint              `cbor:"0,keyasint" json:"tip"`
  PreviousHash string            `cbor:"1,keyasint", json:"previousHash"`
  InitTime     uint              `cbor:"2,keyasint" json:"initTime"`
  Genesis      psl.Genesis       `cbor:"3,keyasint" json:"genesis"`
  Params       psl.Params        `cbor:"4,keyasint" json:"params"`
  Requests     []psl.Request     `cbor:"5,keyasint" json:"requests"`
  Certificates []psl.Certificate `cbor:"6,keyasint" json:"certificates"`
  Accounts     map[string]uint   `cbor:"7,keyasint" json:"accounts"`
  Mempool      Mempool           `cbor:"8,keyasint" json:"mempool"`
  Volitile     Volitile          `cbor:"9,keyasint" json:"volitile"`
}

func NewLedger() Ledger {
  return Ledger{
    Requests: make([]psl.Request, 0),
    Certificates: make([]psl.Certificate, 0),
    Accounts: make(map[string]uint),
    Mempool: NewMempool(),
    Volitile: NewVolitile(),
    Disk: NewDisk(),
  }
}

func LedgerFromCBOR(cborBytes []byte) (Ledger, error) {
  var ledger Ledger
  err := cbor.Unmarshal(cborBytes, &ledger)
  if err != nil { return NewLedger(), err }
  return ledger, nil
}

func LedgerFromHex(hexString string) (Ledger, error) {
  cborBytes, err := hex.DecodeString(hexString)
  if err != nil { return NewLedger(), err }
  ledger, err :=  LedgerFromCBOR(cborBytes)
  if err != nil { return NewLedger(), err }
  return ledger, nil
}

func (ledger *Ledger) ToCBOR() ([]byte, error) {
  cborBytes, err := cbor.Marshal(ledger)
  if err != nil { return nil, err }
  return cborBytes, nil
}

func (ledger *Ledger) ToHex() (string, error) {
  cborBytes, err := ledger.ToCBOR()
  if err != nil { return "", err }
  return hex.EncodeToString(cborBytes), nil
}

func (ledger *Ledger) ToJSON() ([]byte, error) {
  jsonBytes, err := json.Marshal(ledger)
  if err != nil { return nil, err }
  return jsonBytes, nil
}

func (ledger *Ledger) AddGenesis(block psl.Block) {
  ledger.InitTime = block.Body.Timestamp
  ledger.Genesis = *block.Body.Genesis
  ledger.Params = block.Body.Genesis.Params
  ledger.AddCertificate(block.Body.Genesis.Certificate)
}

func (ledger *Ledger) AddBlock(block psl.Block) error {
  if !block.Verify() { return errors.New("invalid block validation") }
  
  if ledger.Tip == 0 && block.Body.ID == 0 { 
    ledger.AddGenesis(block) 
  } else {
    if block.Body.PreviousHash != ledger.PreviousHash {
      return errors.New("invalid block PreviousHash")
    }
    if block.Body.ID != ledger.Tip + 1 {
      return errors.New("invalid block ID")
    }
  } 
  
  blockCBOR, err := block.ToCBOR()
  if err != nil { return err }
  
  err = ledger.Disk.SaveBlockCBOR(block.Body.ID, blockCBOR)
  if err != nil { return err }
  
  ledger.PreviousHash = block.Header.Hash
  ledger.Tip += 1
  return nil
}

func (ledger *Ledger) AddRequest(request psl.Request) {
  ledger.Requests = append(ledger.Requests, request)
}

func (ledger *Ledger) AddCertificate(certificate psl.Certificate) {
  ledger.Certificates = append(ledger.Certificates, certificate)
}

func (ledger *Ledger) UpdateParams(params psl.Params) {
  ledger.Params = params
}

func (ledger *Ledger) AddAccount(account string, amount uint) {
  _, ok := ledger.Accounts[account]
  if !ok { ledger.Accounts[account] = amount }
}

func (ledger *Ledger) WithdrawalFromAccount(account string, amount uint) bool {
  _, ok := ledger.Accounts[account]
  if !ok { return false }
  if ledger.Accounts[account] < amount { return false }
  ledger.Accounts[account] -= amount
  return true
}

func (ledger *Ledger) DepositToAccount(account string, amount uint) bool {
  _, ok := ledger.Accounts[account]
  if !ok { 
    ledger.Accounts[account] = amount
    return true
  }
  
  if ledger.Accounts[account] > math.MaxUint-amount { return false }
  
  ledger.Accounts[account] += amount
  return true
}