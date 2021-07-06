package monacoin

import (
	"github.com/btcsuite/btcd/chaincfg"
)

func init() {
	if err := chaincfg.Register(&MainNetParams); err != nil {
		panic(err)
	}
	if err := chaincfg.Register(&TestNetParams); err != nil {
		panic(err)
	}
	if err := chaincfg.Register(&RegressionNetParams); err != nil {
		panic(err)
	}
}

// MainNetParams returns the chain configuration for mainnet.
var MainNetParams = chaincfg.Params{
	Name: "mainnet",
	Net:  0xdbb6c0fb,

	// Address encoding magics
	PubKeyHashAddrID: 50,
	ScriptHashAddrID: 55,
	PrivateKeyID:     176,

	// BIP32 hierarchical deterministic extended key magics
	HDPrivateKeyID: [4]byte{0x04, 0x88, 0xad, 0xe4}, // starts with xprv
	HDPublicKeyID:  [4]byte{0x04, 0x88, 0xb2, 0x1e}, // starts with xpub

	// Human-readable part for Bech32 encoded segwit addresses, as defined in
	// BIP 173. 
	Bech32HRPSegwit: "mona",
}

// TestNetParams returns the chain configuration for testnet.
var TestNetParams = chaincfg.Params{
	Name: "testnet",
	Net:  0xf1c8d2fd,

	// Address encoding magics
	PubKeyHashAddrID: 111,
	ScriptHashAddrID: 117,
	PrivateKeyID:     239,

	// BIP32 hierarchical deterministic extended key magics
	HDPrivateKeyID: [4]byte{0x04, 0x35, 0x83, 0x94}, // starts with tprv
	HDPublicKeyID:  [4]byte{0x04, 0x35, 0x87, 0xcf}, // starts with tpub

	// Human-readable part for Bech32 encoded segwit addresses, as defined in
	// BIP 173. 
	Bech32HRPSegwit: "tmona",
}

// RegressionNetParams returns the chain configuration for regression net.
var RegressionNetParams = chaincfg.Params{
	Name: "regtest",

	// Monacoin has 0xdab5bffa as RegTest (same as Bitcoin's RegTest).
	// Setting it to an arbitrary value (leet_hex(monacoin)), so that we can
	// register the regtest network.
	Net: 0xdab5bffa,

	// Address encoding magics
	PubKeyHashAddrID: 111,
	ScriptHashAddrID: 196,
	PrivateKeyID:     239,

	// BIP32 hierarchical deterministic extended key magics
	HDPrivateKeyID: [4]byte{0x04, 0x35, 0x83, 0x94}, // starts with tprv
	HDPublicKeyID:  [4]byte{0x04, 0x35, 0x87, 0xcf}, // starts with tpub

	// Human-readable part for Bech32 encoded segwit addresses, as defined in
	// BIP 173. 
	Bech32HRPSegwit: "rmona",
}
