package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEthereumAddress(t *testing.T) {
	testInputs := map[string]bool{
		"0x123aDb": false,
		"hello 1 asdas asdas das d asd asd as das d": false,
		"0x1230000000000000000000000000000000000000": true,
		"0x0000000000000000000000000000000000000000": true,
		"0x123ass0000000000000000000000000000000000": false,
		"0xasgdjahsgdhjasgdhjasgdhjasgdhasgdhjasgdj": false,
		"0x9e27c0e4166e7313da6e72f761e56103AebE1DD1": true,
		"; drop tableasdmasndm,nznm,zXmzXasjdmnasmd": false,
		"\t\n\n\n\n\n\n\n\n\n\t\t\t\t\t\t\t\t\t\t\t": false,
		"ngày mai sẽ đi cắm trại chơi chơi ồ yeah\t": false,
		"122378467382647823647823646238468234623843": false,
		"12sdfiuyasuiduasdwe89458h34543j5jk435jj3\t": false,
		"0x123;asdasdaaasdaasdasdasdqwedqwdaszxazzx": false,
		"0xio24h23j4b23b4nmberiu345435b34b4354123--": false,
		"--134nb5n43nbm45nbm453nbm453nm45n3m435nm23": false,
		"; --ajhsdhjsdsahjdhjkajskdhajskhdhjkasdjka": false,
		"true":   false,
		"0xtrue": false,
	}
	for k, v := range testInputs {
		isAddressEthereum := IsEthereumAddressFromString(k)
		assert.Equal(t, v, isAddressEthereum)
	}
}
