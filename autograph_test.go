package autograph

import (
	"testing"
)

type StringPayload struct {
	message string
}

func (s StringPayload) Serialize() []byte {
	return []byte(s.message)
}

func TestStringPayloadSerializer(t *testing.T) {

	message := StringPayload{message: "foo"}
	signer := NewSigner(message)

	signer.Sign([]byte("spam"))

	if err := signer.Verify([]byte("eggs")); err == nil {
		t.Errorf("Verification with wrong key succeeded")
	}

	if err := signer.Verify([]byte("spam")); err != nil {
		t.Errorf("Verification with correct key failed")
	}

}
