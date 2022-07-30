package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEtherToWei(t *testing.T) {
	wei := EtherToWei("1")
	assert.Equal(t, "1000000000000000000", wei.String())

	wei2 := EtherToWei("10000000000")
	assert.Equal(t, "10000000000000000000000000000", wei2.String())
}

func TestWeiToEther(t *testing.T) {
	ether := WeiToEther("1000000000000000000")
	assert.Equal(t, "1", ether.String())

	ether2 := WeiToEther("10000000000000000000000000000")
	assert.Equal(t, "10000000000", ether2.String())
}
