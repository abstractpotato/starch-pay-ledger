package main

import (
  "fmt"
  "time"
  "github.com/abstractpotato/potato-serialization-lib"
  "github.com/abstractpotato/starch-pay-ledger"
)

func GetPrivateKey() []byte {
  skey, err := psl.GenerateKeys("")
  if err != nil { panic(err) }
  return skey
}

func main() {
  chain := ledger.NewLedger()
  genesisBlock := getGenesisBlock()
  err := chain.AddBlock(genesisBlock)
  if err != nil { panic(err) }
  
  chainJSON, err := chain.ToJSON()
  if err != nil { panic(err) }
  
  fmt.Printf("%s\n", chainJSON)
}

func getGenesisBlock() psl.Block {
  privateKey := GetPrivateKey()
  addr, err := psl.GenerateEnterpriseAddr(privateKey, true)
  if err != nil { panic(err) }
  publicKey, err := psl.GetPublicKey(privateKey[:32])
  if err != nil { panic(err) }
  
  // initital protocol parameters
  params := psl.NewParams()
  params.Network = 0
  params.MaxBlockHeaderSize = 172 // 128 bytes
  params.MaxBlockBodySize = 4000000 // 4 MB or ~15k simple transactions
  params.MaxTxSize = 4000 // 4 KB
  params.TxFeePerByte = 430
  params.MinTxFee = params.TxFeePerByte * 175 // signature size
  params.SlotsPerEpoch = 432000
  params.SlotTimeInMs = 1000
  params.ProtocolVersion = 0

  // initial node certificate
  cert := psl.NewCertificate()
  cert.RewardAddr = addr
  cert.PublicKey = publicKey
  cert.AddRelay("0.0.0.0:5001")
  cert.AddRelay("0.0.0.0:5002")
  cert.Status = 1
  
  genesis := psl.Genesis{}
  genesis.Seed = []byte("bonepool")
  genesis.Certificate = cert
  genesis.Params = params

  block := psl.NewBlock()
  block.Body.Genesis = &genesis
  block.Body.Timestamp = uint(time.Now().UnixMilli())
  block.Hash()
  
  err = block.Sign(privateKey)
  if err != nil { panic(err) }
  
  return block
}