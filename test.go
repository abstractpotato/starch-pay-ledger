package main

import "github.com/abstractpotato/starch-pay-ledger/ledger"

func main() {
  connection := ledger.NewConnection()
  connection.SetUp()
}
