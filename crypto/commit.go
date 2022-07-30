package crypto

import "github.com/ethereum/go-ethereum/crypto"

// digest hashes all passed byte slices.
// The passed slices won't be mutated.
func digest(ms ...[]byte) []byte {
	h := crypto.Keccak256Hash(ms...)
	return h.Bytes()
}
