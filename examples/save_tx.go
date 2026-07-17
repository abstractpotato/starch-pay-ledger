package main

import (
  "fmt"
  PSL "github.com/abstractpotato/potato-serialization-lib/psl"
  Builders "github.com/abstractpotato/potato-serialization-lib/builders"
  Ledger   "github.com/abstractpotato/starch-pay-ledger/ledger"
)

func main() {
  // sample param data
  params := PSL.NewParams()
  params.Network = 0
  params.MaxTxSize = 4000
  params.TxFeePerByte = 430
  params.MinTxFee = params.TxFeePerByte * 175 // signature size

  txBuilder := Builders.NewTxBuilder()
  txBuilder.Params = params

  output := PSL.SimpleOutput{}
  output.To = "target_cardano_addr"
  output.Asset = "policy_id+asset_name"
  output.Amount = 10000

  txBuilder.AddSimpleOutput(output)
  txBuilder.Build()
  txHash := txBuilder.Tx.Header.Hash

  fmt.Printf("save: %+v\n\n", txBuilder.Tx)

  txCBOR, err := txBuilder.Tx.ToCBOR()
  if err != nil { panic(err) }

  disk := Ledger.NewDisk()
  disk.CreatedDirs()
  disk.SaveTxCBOR(txHash, txCBOR)

  txCBOR, err = disk.GetTxCBOR(txHash)
  if err != nil { panic(err) }

  tx, err := PSL.TransactionFromCBOR(txCBOR)
  if err != nil { panic(err) }

  // passed if loads the same hash
  tx.Hash()

  fmt.Printf("load: %+v\n", tx)

  // deletes transaction
  disk.DeleteTx(txHash)
}
