package ledger

import (
  PSL "github.com/abstractpotato/potato-serialization-lib/psl"
  "github.com/fxamacker/cbor/v2"
  "encoding/hex"
  "encoding/json"
)

type Ledger struct {
  Tip          uint              `cbor:"0,keyasint" json:"tip"`
  Genesis      PSL.Genesis       `cbor:"1,keyasint" json:"genesis"`
  Params       PSL.Params        `cbor:"2,keyasint" json:"params"`
  Requests     []PSL.Request     `cbor:"3,keyasint" json:"requests"`
  Certificates []PSL.Certificate `cbor:"4,keyasint" json:"certificates"`
  Mempool      Mempool           `cbor:"5,keyasint" json:"mempool"`
}

func NewLedger() Ledger {
  return Ledger{
    Requests: make([]PSL.Request, 0),
    Certificates: make([]PSL.Certificate, 0),
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
