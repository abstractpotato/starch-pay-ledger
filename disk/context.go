package psl

import(
  "github.com/fxamacker/cbor/v2"
  "encoding/hex"
  "encoding/json"
)

type Context struct {
  Params Params `cbor:"0,keyasint" json:"params"`
  Epoch  uint   `cbor:"1,keyasint" json:"epoch"`
  Slot   uint   `cbor:"2,keyasint" json:"slot"`
  Tip    uint   `cbor:"3,keyasint" json:"tip"`
}

func NewContext() Context {
  return Context{}
}

func ContextFromCBOR(cborBytes []byte) (Context, error) {
  var context Context
  err := cbor.Unmarshal(cborBytes, &context)
  if err != nil { return NewContext(), err }
  return context, nil
}

func ContextFromHex(hexString string) (Context, error) {
  cborBytes, err := hex.DecodeString(hexString)
  if err != nil { return NewContext(), err }
  context, err :=  ContextFromCBOR(cborBytes)
  if err != nil { return NewContext(), err }
  return context, nil
}

func (context *Context) ToCBOR() ([]byte, error) {
  cborBytes, err := cbor.Marshal(context)
  if err != nil { return nil, err }
  return cborBytes, nil
}

func (context *Context) ToHex() (string, error) {
  cborBytes, err := context.ToCBOR()
  if err != nil { return "", err }
  return hex.EncodeToString(cborBytes), nil
}

func (context *Context) ToJSON() ([]byte, error) {
  jsonBytes, err := json.Marshal(context)
  if err != nil { return nil, err }
  return jsonBytes, nil
}
