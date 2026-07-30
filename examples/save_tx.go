package main

import (
  "fmt"
  "github.com/abstractpotato/potato-serialization-lib"
  "github.com/abstractpotato/starch-pay-ledger"
)


func GetPrivateKey() []byte {
  skey, err := psl.GenerateKeys("")
  if err != nil { panic(err) }
  return skey
}

func main() {
  privateKey := GetPrivateKey()
  addr, err := psl.GenerateEnterpriseAddr(privateKey, true)
  if err != nil { panic(err) }
  
  // sample param data
  params := psl.NewParams()
  params.Network = 0
  params.MaxTxSize = 4000
  params.TxFeePerByte = 430
  params.MinTxFee = params.TxFeePerByte * 176 // typical header size
  
  txBuilder := psl.NewTxBuilder()
  txBuilder.Params = params
  
  output := psl.SimpleOutput{}
  output.From = addr
  output.To = addr
  output.Asset = "3d77d63dfa6033be98021417e08e3368cc80e67f8d7afa196aaa0b3953746172636820546f6b656e"
  output.Amount = 10000
  
  txBuilder.AddSimpleOutput(output)
  txBuilder.Build()
  txHash := txBuilder.Tx.Header.Hash
  
  txBuilder.Sign(privateKey)
  
  txJSON, err := txBuilder.Tx.ToJSON()
  if err != nil { panic(err) }
  
  fmt.Printf("save:\n%s\n\n", txJSON)
  
  // LOAD TX ===================================================================
  txCBOR, err := txBuilder.Tx.ToCBOR()
  if err != nil { panic(err) }
  
  disk := ledger.NewDisk()
  disk.CreatedDirs()
  disk.SaveTxCBOR(txHash, txCBOR)

  txCBOR, err = disk.GetTxCBOR(txHash)
  if err != nil { panic(err) }

  tx, err := psl.TransactionFromCBOR(txCBOR)
  if err != nil { panic(err) }

  // passed if loads the same hash
  tx.Hash()
  
  txJSON, err = txBuilder.Tx.ToJSON()
  if err != nil { panic(err) }
  
  txBodyCBOR, err := tx.Body.ToCBOR()
  if err != nil { panic(err) }
  
  fmt.Printf("load:\n%+s\n", txJSON)
  fmt.Printf("size: %v bytes\n", len(txCBOR))
  fmt.Printf("header size: %v bytes\n", len(txCBOR) - len(txBodyCBOR))
  fmt.Printf("body size: %v bytes\n", len(txBodyCBOR))
  
  fmt.Printf("verified: %v\n", tx.Verify())

  // deletes transact
  disk.DeleteTx(txHash)
}
