package ledger

import (
  PSL "github.com/abstractpotato/potato-serialization-lib/psl"
  "github.com/fxamacker/cbor/v2"
  "encoding/hex"
  "encoding/json"
)

type Snapshot struct {
  Context    Context                    `cbor: "context"`
  Requests   map[string]PSL.Request     `cbor: "requests"`
  Certs      map[string]PSL.Certificate `cbor: "certificates"`
  Validators []PSL.Validator            `cbor: "validators"`
  Accounts   map[string]Account         `cbor: "accounts"`
}

func NewSnapshot() Snapshot {
  return Snapshot{
    Context: PSL.NewContext(),
    Requests: make(map[string]PSL.Request),
    Certs: make(map[string]PSL.Certificate),
    Validators: make([]Validator, 0),
    Accounts: make(map[string]Account),
  }
}

func SnapshotFromCBOR(cborBytes []byte) (Snapshot, error) {
  var snapshot Snapshot
  err := cbor.Unmarshal(cborBytes, &snapshot)
  if err != nil { return NewSnapshot(), err }
  return snapshot, nil
}

func SnapshotFromHex(hexString string) (Snapshot, error) {
  cborBytes, err := hex.DecodeString(hexString)
  if err != nil { return NewSnapshot(), err }
  snapshot, err := SnapshotFromCBOR(cborBytes)
  if err != nil { return NewSnapshot(), err }
  return snapshot, nil
}

func (snapshot *Snapshot) ToCBOR() ([]byte, error) {
  cborBytes, err := cbor.Marshal(snapshot)
  if err != nil { return nil, err}
  return cborBytes, nil
}

func (snapshot *Snapshot) ToHex() (string, error) {
  cborBytes, err := snapshot.ToCBOR()
  if err != nil { return "", err }
  return hex.EncodeToString(cborBytes), nil
}

func (snapshot *Snapshot) ToJSON() ([]byte, error) {
  jsonBytes, err := json.Marshal(snapshot)
  if err != nil { return nil, err }
  return jsonBytes, nil
}
