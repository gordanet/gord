package bip32

import "github.com/pkg/errors"

// BitcoinMainnetPrivate is the version that is used for
// bitcoin mainnet bip32 private extended keys.
// Ecnodes to xprv in base58.
var BitcoinMainnetPrivate = [4]byte{
	0x04,
	0x88,
	0xad,
	0xe4,
}

// BitcoinMainnetPublic is the version that is used for
// bitcoin mainnet bip32 public extended keys.
// Ecnodes to xpub in base58.
var BitcoinMainnetPublic = [4]byte{
	0x04,
	0x88,
	0xb2,
	0x1e,
}

// GorMainnetPrivate is the version that is used for
// gor mainnet bip32 private extended keys.
// Ecnodes to xprv in base58.
var GorMainnetPrivate = [4]byte{
	0x03,
	0x8f,
	0x2e,
	0xf4,
}

// GorMainnetPublic is the version that is used for
// gor mainnet bip32 public extended keys.
// Ecnodes to kpub in base58.
var GorMainnetPublic = [4]byte{
	0x03,
	0x8f,
	0x33,
	0x2e,
}

// GorTestnetPrivate is the version that is used for
// gor testnet bip32 public extended keys.
// Ecnodes to ktrv in base58.
var GorTestnetPrivate = [4]byte{
	0x03,
	0x90,
	0x9e,
	0x07,
}

// GorTestnetPublic is the version that is used for
// gor testnet bip32 public extended keys.
// Ecnodes to ktub in base58.
var GorTestnetPublic = [4]byte{
	0x03,
	0x90,
	0xa2,
	0x41,
}

// GorDevnetPrivate is the version that is used for
// gor devnet bip32 public extended keys.
// Ecnodes to kdrv in base58.
var GorDevnetPrivate = [4]byte{
	0x03,
	0x8b,
	0x3d,
	0x80,
}

// GorDevnetPublic is the version that is used for
// gor devnet bip32 public extended keys.
// Ecnodes to xdub in base58.
var GorDevnetPublic = [4]byte{
	0x03,
	0x8b,
	0x41,
	0xba,
}

// GorSimnetPrivate is the version that is used for
// gor simnet bip32 public extended keys.
// Ecnodes to ksrv in base58.
var GorSimnetPrivate = [4]byte{
	0x03,
	0x90,
	0x42,
	0x42,
}

// GorSimnetPublic is the version that is used for
// gor simnet bip32 public extended keys.
// Ecnodes to xsub in base58.
var GorSimnetPublic = [4]byte{
	0x03,
	0x90,
	0x46,
	0x7d,
}

func toPublicVersion(version [4]byte) ([4]byte, error) {
	switch version {
	case BitcoinMainnetPrivate:
		return BitcoinMainnetPublic, nil
	case GorMainnetPrivate:
		return GorMainnetPublic, nil
	case GorTestnetPrivate:
		return GorTestnetPublic, nil
	case GorDevnetPrivate:
		return GorDevnetPublic, nil
	case GorSimnetPrivate:
		return GorSimnetPublic, nil
	}

	return [4]byte{}, errors.Errorf("unknown version %x", version)
}

func isPrivateVersion(version [4]byte) bool {
	switch version {
	case BitcoinMainnetPrivate:
		return true
	case GorMainnetPrivate:
		return true
	case GorTestnetPrivate:
		return true
	case GorDevnetPrivate:
		return true
	case GorSimnetPrivate:
		return true
	}

	return false
}
