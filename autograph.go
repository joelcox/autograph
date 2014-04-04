package autograph

import (
	"bytes"
	"crypto/sha256"
	"errors"
)

type Manifest struct {
	signature [32]byte
	timestamp int
	namespace string
}

type Payload interface {
	Serialize() []byte
	Manifest() Manifest
}

type StringPayload struct {
	message string
	Manifest
}

func (p *StringPayload) Serialize() []byte {
	return []byte(p.message)
}

func (p *StringPayload) Sign(key []byte) (signature [32]byte) {
	message := p.Serialize()
	signature = sha256.Sum256(append(message[:], key[:]...))
	p.Manifest.signature = signature
	return
}

func (p *StringPayload) Verify(key []byte) error {
	message := p.Serialize()
	signature := sha256.Sum256(append(message[:], key[:]...))

	if !bytes.Equal(p.Manifest.signature[:32], signature[:32]) {
		return errors.New("autograph: invalid signature")
	}

	return nil
}
