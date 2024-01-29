// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package tests

import (
	"encoding/json"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum-op/common"
	"github.com/ethereum/go-ethereum-op/common/math"
)

var _ = (*stEnvMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (s stEnv) MarshalJSON() ([]byte, error) {
	type stEnv struct {
		Coinbase   common.UnprefixedAddress `json:"currentCoinbase"   gencodec:"required"`
		Difficulty *math.HexOrDecimal256    `json:"currentDifficulty" gencodec:"optional"`
		Random     *math.HexOrDecimal256    `json:"currentRandom"     gencodec:"optional"`
		GasLimit   math.HexOrDecimal64      `json:"currentGasLimit"   gencodec:"required"`
		Number     math.HexOrDecimal64      `json:"currentNumber"     gencodec:"required"`
		Timestamp  math.HexOrDecimal64      `json:"currentTimestamp"  gencodec:"required"`
		BaseFee    *math.HexOrDecimal256    `json:"currentBaseFee"    gencodec:"optional"`
	}
	var enc stEnv
	enc.Coinbase = common.UnprefixedAddress(s.Coinbase)
	enc.Difficulty = (*math.HexOrDecimal256)(s.Difficulty)
	enc.Random = (*math.HexOrDecimal256)(s.Random)
	enc.GasLimit = math.HexOrDecimal64(s.GasLimit)
	enc.Number = math.HexOrDecimal64(s.Number)
	enc.Timestamp = math.HexOrDecimal64(s.Timestamp)
	enc.BaseFee = (*math.HexOrDecimal256)(s.BaseFee)
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (s *stEnv) UnmarshalJSON(input []byte) error {
	type stEnv struct {
		Coinbase   *common.UnprefixedAddress `json:"currentCoinbase"   gencodec:"required"`
		Difficulty *math.HexOrDecimal256     `json:"currentDifficulty" gencodec:"optional"`
		Random     *math.HexOrDecimal256     `json:"currentRandom"     gencodec:"optional"`
		GasLimit   *math.HexOrDecimal64      `json:"currentGasLimit"   gencodec:"required"`
		Number     *math.HexOrDecimal64      `json:"currentNumber"     gencodec:"required"`
		Timestamp  *math.HexOrDecimal64      `json:"currentTimestamp"  gencodec:"required"`
		BaseFee    *math.HexOrDecimal256     `json:"currentBaseFee"    gencodec:"optional"`
	}
	var dec stEnv
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Coinbase == nil {
		return errors.New("missing required field 'currentCoinbase' for stEnv")
	}
	s.Coinbase = common.Address(*dec.Coinbase)
	if dec.Difficulty != nil {
		s.Difficulty = (*big.Int)(dec.Difficulty)
	}
	if dec.Random != nil {
		s.Random = (*big.Int)(dec.Random)
	}
	if dec.GasLimit == nil {
		return errors.New("missing required field 'currentGasLimit' for stEnv")
	}
	s.GasLimit = uint64(*dec.GasLimit)
	if dec.Number == nil {
		return errors.New("missing required field 'currentNumber' for stEnv")
	}
	s.Number = uint64(*dec.Number)
	if dec.Timestamp == nil {
		return errors.New("missing required field 'currentTimestamp' for stEnv")
	}
	s.Timestamp = uint64(*dec.Timestamp)
	if dec.BaseFee != nil {
		s.BaseFee = (*big.Int)(dec.BaseFee)
	}
	return nil
}
