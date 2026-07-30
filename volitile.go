package ledger

// this is where blocks go before they get accepted into immutable

type Volitile struct {
  Blocks       map[uint]TempBlock `cbor:"0,keyasint"`
  Transactions map[string][]byte  `cbor:"2,keyasint"`
}

type TempBlock struct {
  CBOR []byte
  Transactions map[string][]byte
}

func NewVolitile() Volitile {
  return Volitile {
    Blocks: make(map[uint]TempBlock),
    Transactions: make(map[string][]byte),
  }
}

// func (volitile *Volitile) ToImmutable() {
//   disk := NewDisk()
// 
//   for id, blockCBOR := range volitile.Blocks {
//     disk.SaveBlockCBOR(id, blockCBOR)
//   }
// }