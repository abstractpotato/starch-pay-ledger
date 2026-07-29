package ledger

import (
  "github.com/abstractpotato/potato-serialization-lib"
  "github.com/fxamacker/cbor/v2"
  "encoding/hex"
  "encoding/json"
)

type Ledger struct {
  Tip          uint              `cbor:"0,keyasint" json:"tip"`
  InitTime     uint              `cbor:"1,keyasint" json:"initTime"`
  Genesis      *psl.Genesis      `cbor:"2,keyasint" json:"genesis"`
  Params       *psl.Params       `cbor:"3,keyasint" json:"params"`
  Requests     []psl.Request     `cbor:"4,keyasint" json:"requests"`
  Certificates []psl.Certificate `cbor:"5,keyasint" json:"certificates"`
  Mempool      Mempool           `cbor:"6,keyasint" json:"mempool"`
}

func NewLedger() Ledger {
  return Ledger{
    Requests: make([]psl.Request, 0),
    Certificates: make([]psl.Certificate, 0),
    Mempool: NewMempool(),
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


func (ledger *Ledger) AddRequest(request psl.Request) {
  ledger.Requests = append(ledger.Requests, request)
}

func (ledger *Ledger) AddCertificate(certificate psl.Certificate) {
  ledger.Certificates = append(ledger.Certificates, certificate)
}
