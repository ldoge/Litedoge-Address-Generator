package main

type Altcoin struct {
	Name             string
	PubKeyHashAddrID byte
	PrivateKeyID     byte
	CoinType         int
}

var altcoins = []Altcoin{
	{"LiteDoge", 0x5a, 0xab, -1}
}
