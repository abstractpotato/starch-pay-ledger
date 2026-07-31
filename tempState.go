package ledger

import (
  "math"
  "github.com/abstractpotato/potato-serialization-lib"
)

type TempState struct {
  Block psl.Block
  Accounts map[string]uint
}

func NewTempState() TempState {
  return TempState{
    Accounts: make(map[string]uint),
  }
}

func (state *TempState) ProcessTx(tx psl.Transaction) bool {

  if !state.ProcessSimpleOutputs(tx) { return false }
  if !state.ProcessMultiOutputs(tx) {return false }
  return state.ProcessAirDropOutput(tx)
}

func (state *TempState) ProcessSimpleOutputs(tx psl.Transaction) bool {

  for _, simpleOutput := range tx.Body.SimpleOutputs {
    from := simpleOutput.From+simpleOutput.Asset
    if !state.Withdrawal(from, simpleOutput.Amount) { return false }

    to := simpleOutput.To+simpleOutput.Asset
    if !state.Deposit(to, simpleOutput.Amount) { return false }
  }
  return true
}

func (state *TempState) ProcessMultiOutputs(tx psl.Transaction) bool {
  for _, multiOutput := range tx.Body.MultiOutputs {
    for _, asset := range multiOutput.Assets {
      from := multiOutput.From+asset.Asset
      if !state.Withdrawal(from, asset.Amount) { return false }

      to := multiOutput.To+asset.Asset
      if !state.Deposit(to, asset.Amount) { return false }
    }
  }
  return true
}

func (state *TempState) ProcessAirDropOutput(tx psl.Transaction) bool {
  if tx.Body.AirDropOutput == nil { return true }

  //checks if Amount is divisible by To
  if len(tx.Body.AirDropOutput.To) == 0 { return false }
  if uint(len(tx.Body.AirDropOutput.To))%tx.Body.AirDropOutput.Amount != 0 {return false}

  from := tx.Body.AirDropOutput.From+tx.Body.AirDropOutput.Asset
  if !state.Withdrawal(from, tx.Body.AirDropOutput.Amount) {return false }

  amount := tx.Body.AirDropOutput.Amount / uint(len(tx.Body.AirDropOutput.To))

  for _, addr := range tx.Body.AirDropOutput.To {
    to := addr+tx.Body.AirDropOutput.Asset
    if !state.Deposit(to, amount) { return false }
  }

  return true
}

func (state *TempState) AccountExists(account string) bool {
  _, ok := state.Accounts[account]
  return ok
}

func (state *TempState) AddAccount(account string, amount uint) {
  state.Accounts[account] = amount
}

func (state *TempState) Withdrawal(account string, amount uint) bool{
  if !state.AccountExists(account) { return false }
  if state.Accounts[account] < amount { return false }
  state.Accounts[account] -= amount
  return true
}

func (state *TempState) Deposit(account string, amount uint) bool {
  if !state.AccountExists(account) {
    state.AddAccount(account, amount)
    return true
  }

  if state.Accounts[account] > math.MaxUint - amount { return false }

  state.Accounts[account] += amount
  return true
}
