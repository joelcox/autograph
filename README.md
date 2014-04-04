Autograph
=========

A package for signing and verifying Go structs.

Work in progress; except breaking changes!

Example
-------

```go
// Implement the Serializer interface
func (s StringPayload) Serialize() []byte {
    return []byte(s.message)
}

// Create a payload and pass it to the signer, which calls Serialize
// on the struct 
payload := StringPayload{message: "foo"}
signer := NewSigner(payload)

// Sign the payload with the key
signer.Sign([]byte("spam"))

// Verify if the payload has been tempered with
if err := signer.Verify([]byte("eggs")); err != nil {
	t.Errorf("Signature verification failed!")
}
```

License
-------

MIT