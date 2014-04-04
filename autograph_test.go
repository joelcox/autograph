package autograph

import (
    "testing"
)

func TestStringPayloadSerializer(t *testing.T) {

    message := StringPayload{message: "foo"}
    message.Sign([]byte("spam"))

    if err := message.Verify([]byte("eggs")); err == nil {
        t.Errorf("Verification with wrong key succeeded")
    }

    if err := message.Verify([]byte("spam")); err != nil {
        t.Errorf("Verification with correct key failed")
    }

}