package main

import "github.com/abstractpotato/potato-serialization-lib"

type TempState struct {
  Block psl.Block
  Accounts map[string]uint
}

func NewTempState() TempState {
  return TempState{
    Accounts: make(map[string]uint),
  }
}

func ProcessTx(transaction psl.Transaction) bool {
  // simple transactions
  // multi transactions
  // airdrop transactions
}

func (state *TempState) AddAccount(account string, amount uint) {
  _, ok := state.Accounts[account]
  if !ok { state.Accounts[account] = amount }
}

func (state *TempState) Withdrawal(account string, amount uint) bool{
  _, ok := state.Accounts[account]
  if !ok { return false }
  if state.Accounts[account] < amount { return false }
  state.Accounts[account] -= amount
  return true
}

func (state *TempState) Deposit(account string, amount uint) bool {
  _, ok := state.Accounts[account]
  if !ok {
    state.Accounts[account] = amount
    return true
  }

  if state.Accounts[account] > math.MaxUint - amount { return false }

  state.Accounts[account] += amount
  return true
}
