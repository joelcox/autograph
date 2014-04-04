// A package for signing and verifing Go structs.
package autograph

import (
	"bytes"
	"crypto/sha256"
	"errors"
)

// Contains a serialized message and corresponding signature.
type Payload struct {
	message   []byte
	signature [32]byte
}

// Allows for serialization to a byte array of more complex structures.
type Serializer interface {
	Serialize() []byte
}

// Creates a new signer from a struct implementing the Serializer interface.
func NewSigner(s Serializer) *Payload {
	return &Payload{message: s.Serialize()}
}

// Signs a message by saving the corresponding signature to the payload.
func (p *Payload) Sign(key []byte) (signature [32]byte) {
	signature = sha256.Sum256(append(p.message[:], key[:]...))
	p.signature = signature
	return
}

// Verifies the message' signature by recomputing the signature and comparing
// the byte array
func (p *Payload) Verify(key []byte) error {
	signature := sha256.Sum256(append(p.message[:], key[:]...))

	if !bytes.Equal(p.signature[:32], signature[:32]) {
		return errors.New("autograph: invalid signature")
	}

	return nil
}
