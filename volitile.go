package ledger

import (
  "math"
  "github.com/abstractpotato/potato-serialization-lib"
)

type Volitile struct {
  Block psl.Block
  Accounts map[string]uint
}

func NewVolitile() Volitile {
  return Volitile{
    Accounts: make(map[string]uint),
  }
}

func (volitile *Volitile) ProcessTx(tx psl.Transaction) bool {

  if !volitile.ProcessSimpleOutputs(tx) { return false }
  if !volitile.ProcessMultiOutputs(tx) {return false }
  return volitile.ProcessAirDropOutput(tx)
}

func (volitile *Volitile) ProcessSimpleOutputs(tx psl.Transaction) bool {

  for _, simpleOutput := range tx.Body.SimpleOutputs {
    from := simpleOutput.From+simpleOutput.Asset
    if !volitile.Withdrawal(from, simpleOutput.Amount) { return false }

    to := simpleOutput.To+simpleOutput.Asset
    if !volitile.Deposit(to, simpleOutput.Amount) { return false }
  }
  return true
}

func (volitile *Volitile) ProcessMultiOutputs(tx psl.Transaction) bool {
  for _, multiOutput := range tx.Body.MultiOutputs {
    for _, asset := range multiOutput.Assets {
      from := multiOutput.From+asset.Asset
      if !volitile.Withdrawal(from, asset.Amount) { return false }

      to := multiOutput.To+asset.Asset
      if !volitile.Deposit(to, asset.Amount) { return false }
    }
  }
  return true
}

func (volitile *Volitile) ProcessAirDropOutput(tx psl.Transaction) bool {
  if tx.Body.AirDropOutput == nil { return true }

  //checks if Amount is divisible by To
  if len(tx.Body.AirDropOutput.To) == 0 { return false }
  if uint(len(tx.Body.AirDropOutput.To))%tx.Body.AirDropOutput.Amount != 0 {return false}

  from := tx.Body.AirDropOutput.From+tx.Body.AirDropOutput.Asset
  if !volitile.Withdrawal(from, tx.Body.AirDropOutput.Amount) {return false }

  amount := tx.Body.AirDropOutput.Amount / uint(len(tx.Body.AirDropOutput.To))

  for _, addr := range tx.Body.AirDropOutput.To {
    to := addr+tx.Body.AirDropOutput.Asset
    if !volitile.Deposit(to, amount) { return false }
  }

  return true
}

func (volitile *Volitile) AccountExists(account string) bool {
  _, ok := volitile.Accounts[account]
  return ok
}

func (volitile *Volitile) AddAccount(account string, amount uint) {
  volitile.Accounts[account] = amount
}

func (volitile *Volitile) Withdrawal(account string, amount uint) bool{
  if !volitile.AccountExists(account) { return false }
  if volitile.Accounts[account] < amount { return false }
  volitile.Accounts[account] -= amount
  return true
}

func (volitile *Volitile) Deposit(account string, amount uint) bool {
  if !volitile.AccountExists(account) {
    volitile.AddAccount(account, amount)
    return true
  }

  if volitile.Accounts[account] > math.MaxUint - amount { return false }

  volitile.Accounts[account] += amount
  return true
}
