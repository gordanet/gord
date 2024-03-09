package difficulty_test

import (
	"fmt"
	"github.com/gordanet/gord/domain/consensus"
	"math"
	"math/big"
	"testing"

	"github.com/gordanet/gord/domain/consensus/utils/testutils"
	"github.com/gordanet/gord/domain/dagconfig"
	"github.com/gordanet/gord/util/difficulty"
)

func TestGetHashrateString(t *testing.T) {
	var results = map[string]string{
		dagconfig.MainnetParams.Name: "1.53 GH/s",
		dagconfig.TestnetParams.Name: "131.07 KH/s",
		dagconfig.DevnetParams.Name:  "830 H/s",
		dagconfig.SimnetParams.Name:  "2.00 KH/s",
	}
	testutils.ForAllNets(t, false, func(t *testing.T, consensusConfig *consensus.Config) {
		targetGenesis := difficulty.CompactToBig(consensusConfig.GenesisBlock.Header.Bits())
		hashrate := difficulty.GetHashrateString(targetGenesis, consensusConfig.TargetTimePerBlock)
		if hashrate != results[consensusConfig.Name] {
			t.Errorf("Expected %s, found %s", results[consensusConfig.Name], hashrate)
		}
	})
}

// TestBigToCompact ensures BigToCompact converts big integers to the expected
// compact representation.
func TestBigToCompact(t *testing.T) {
	tests := []struct {
		in  string
		out uint32
	}{
		{"0000000000000000000000000000000000000000000000000000000000000000", 0},
		{"-1", 25231360},
		{"9223372036854775807", 142606335},
		{"922337203685477580712312312123487", 237861256},
		{"128", 0x02008000},
	}

	for x, test := range tests {
		n := new(big.Int)
		n.SetString(test.in, 10)
		r := difficulty.BigToCompact(n)
		if r != test.out {
			t.Errorf("TestBigToCompact test #%d failed: got %d want %d\n",
				x, r, test.out)
			return
		}
	}
}

// TestCompactToBig ensures CompactToBig converts numbers using the compact
// representation to the expected big integers.
func TestCompactToBig(t *testing.T) {
	tests := []struct {
		before uint32
		intHex string
		after  uint32
	}{
		{math.MaxUint32, "-7fffff000000000000000000000000000000000000000000000000000000000000000000000000" +
			"000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000" +
			"000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000" +
			"000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000" +
			"000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000" +
			"000000000000000000000000", math.MaxUint32},
		{0x00000000, "0000000000000000000000000000000000000000000000000000000000000000", 0x00000000},
		{0x0989680, "0000000000000000000000000000000000000000000000000000000000000000", 0x00000000},
		{0x87fffff, "0000000000000000000000000000000000000000000000007fffff0000000000", 0x87fffff},
		{0x1810000, "-000000000000000000000000000000000000000000000000000000000000001", 0x1810000},
		{0xe2d7988, "0000000000000000000000000000000000002d79880000000000000000000000", 0xe2d7988},
		{0x00123456, "0000000000000000000000000000000000000000000000000000000000000000", 0x00000000},
		{0x01003456, "0000000000000000000000000000000000000000000000000000000000000000", 0x00000000},
		{0x02000056, "0000000000000000000000000000000000000000000000000000000000000000", 0x00000000},
		{0x03000000, "0000000000000000000000000000000000000000000000000000000000000000", 0x00000000},
		{0x04000000, "0000000000000000000000000000000000000000000000000000000000000000", 0x00000000},
		{0x00923456, "0000000000000000000000000000000000000000000000000000000000000000", 0x00000000},
		{0x01803456, "0000000000000000000000000000000000000000000000000000000000000000", 0x00000000},
		{0x02800056, "0000000000000000000000000000000000000000000000000000000000000000", 0x00000000},
		{0x03800000, "0000000000000000000000000000000000000000000000000000000000000000", 0x00000000},
		{0x04800000, "0000000000000000000000000000000000000000000000000000000000000000", 0x00000000},
		{0x01123456, "0000000000000000000000000000000000000000000000000000000000000012", 0x01120000},
		{0x02008000, "0000000000000000000000000000000000000000000000000000000000000080", 0x02008000},
		{0x01fedcba, "-00000000000000000000000000000000000000000000000000000000000007e", 0x01fe0000},
		{0x02123456, "0000000000000000000000000000000000000000000000000000000000001234", 0x02123400},
		{0x03123456, "0000000000000000000000000000000000000000000000000000000000123456", 0x03123456},
		{0x04123456, "0000000000000000000000000000000000000000000000000000000012345600", 0x04123456},
		{0x04923456, "-000000000000000000000000000000000000000000000000000000012345600", 0x04923456},
		{0x05009234, "0000000000000000000000000000000000000000000000000000000092340000", 0x05009234},
		{0x20123456, "1234560000000000000000000000000000000000000000000000000000000000", 0x20123456},
		{0xff123456, "123456000000000000000000000000000000000000000000000000000000000000000000000000000000" +
			"000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000" +
			"000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000" +
			"000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000" +
			"000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000", 0xff123456},
	}

	for i, test := range tests {
		n := difficulty.CompactToBig(test.before)
		convertBack := difficulty.BigToCompact(n)
		got := fmt.Sprintf("%064x", n)
		if got != test.intHex {
			t.Errorf("TestCompactToBig test #%d failed: got %s want %s, input: 0x%08x",
				i, got, test.intHex, test.before)
		}
		if convertBack != test.after {
			t.Errorf("TestCompactToBig test #%d failed: got: 0x%08x want 0x%08x input: 0x%08x", i, convertBack, test.after, test.before)
		}
	}
}