package psl

import(
  "github.com/fxamacker/cbor/v2"
  "encoding/hex"
  "encoding/json"
)

type Account struct {
  Addr string            `cbor: "addr"`
  Assets map[string]uint `cbor: "assets"`
}

func NewAccount() Account {
  return Account{
    Assets: make(map[string]uint),
  }
}

func AccountFromCBOR(cborBytes []byte) (Account, error) {
  var account Account
  err := cbor.Unmarshal(cborBytes, &account)
  if err != nil { return NewAccount(), err }
  return account, nil
}

func AccountFromHex(hexString string) (Account, error) {
  cborBytes, err := hex.DecodeString(hexString)
  if err != nil { return NewAccount(), err }
  account, err :=  AccountFromCBOR(cborBytes)
  if err != nil { return NewAccount(), err }
  return account, nil
}

func (account *Account) ToCBOR() ([]byte, error) {
  cborBytes, err := cbor.Marshal(account)
  if err != nil { return nil, err }
  return cborBytes, nil
}

func (account *Account) ToHex() (string, error) {
  cborBytes, err := account.ToCBOR()
  if err != nil { return "", err }
  return hex.EncodeToString(cborBytes), nil
}

func (account *Account) ToJSON() ([]byte, error) {
  jsonBytes, err := json.Marshal(account)
  if err != nil { return nil, err }
  return jsonBytes, nil
}
