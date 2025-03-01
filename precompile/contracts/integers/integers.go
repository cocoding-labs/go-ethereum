package integers

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/precompile"
	"github.com/ethereum/go-ethereum/precompile/bindings"
)

type Integers struct {
	precompile.StatefulPrecompiledContract
}

func NewIntegers() *Integers {
	return &Integers{
		StatefulPrecompiledContract: precompile.NewStatefulPrecompiledContract(
			bindings.IntegersABI,
		),
	}
}

// Solidity: function toString(uint256 _i) view returns(string)
func (c *Integers) ToString(ctx precompile.StatefulContext, i *big.Int) (string, error) {
	return i.String(), nil
}

// Solidity: function toString(int256 _i) view returns(string)
func (c *Integers) ToString0(ctx precompile.StatefulContext, i *big.Int) (string, error) {
	return i.String(), nil
}

// Solidity: function toHexString(uint256 _i) view external returns (string memory)
func (c *Integers) ToHexString(ctx precompile.StatefulContext, i *big.Int) (string, error) {
	hex := fmt.Sprintf("%x", i)
	if len(hex)%2 != 0 {
		hex = "0" + hex
	}
	hex = "0x" + hex
	return hex, nil
}

// Solidity: function fromHexString(string calldata _str) view external returns (uint256)
func (c *Integers) FromHexString(ctx precompile.StatefulContext, str string) (*big.Int, error) {
	str = strings.TrimPrefix(strings.ToLower(str), "0x")
	i, success := new(big.Int).SetString(str, 16)
	if !success {
		return nil, fmt.Errorf("could not convert hex string to uint256")
	}
	return i, nil
}

func (c *Integers) RequiredGas(input []byte) uint64 {
	return precompile.GasQuickStep
}
