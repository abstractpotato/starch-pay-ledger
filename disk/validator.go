package psl

import(
  "github.com/fxamacker/cbor/v2"
  "encoding/hex"
  "encoding/json"
)

type Validator struct {
  Addr           string   `cbor: "addr"`
  CertificateTx  string   `cbor: "certificate"`
  Relays         []string `cbor: "relays"`
  Status         uint     `cbor: "status"`
}

func NewValidator() Validator {
  return Validator{}
}

func ValidatorFromCBOR(cborBytes []byte) (Validator, error) {
  var validator Validator
  err := cbor.Unmarshal(cborBytes, &validator)
  if err != nil { return NewValidator(), err }
  return validator, nil
}

func ValidatorFromHex(hexString string) (Validator, error) {
  cborBytes, err := hex.DecodeString(hexString)
  if err != nil { return NewValidator(), err }
  validator, err :=  ValidatorFromCBOR(cborBytes)
  if err != nil { return NewValidator(), err }
  return validator, nil
}

func (validator *Validator) ToCBOR() ([]byte, error) {
  cborBytes, err := cbor.Marshal(validator)
  if err != nil { return nil, err }
  return cborBytes, nil
}

func (validator *Validator) ToHex() (string, error) {
  cborBytes, err := validator.ToCBOR()
  if err != nil { return "", err }
  return hex.EncodeToString(cborBytes), nil
}

func (validator *Validator) ToJSON() ([]byte, error) {
  jsonBytes, err := json.Marshal(validator)
  if err != nil { return nil, err }
  return jsonBytes, nil
}
