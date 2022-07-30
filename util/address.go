package util

import (
	"regexp"

	"github.com/ethereum/go-ethereum/common"
)

const (
	regexEthereumAddress = "^0x[0-9a-fA-F]{40}$"
)

func IsEthereumAddressFromString(address string) bool {
	return regexp.MustCompile(regexEthereumAddress).MatchString(address)
}

func IsEthereumAddress(address common.Address) bool {
	return IsEthereumAddressFromString(address.Hex())
}
