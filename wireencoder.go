package mansbridge

type WireEncoder interface {
	Encode(wireMsg WireMessage) ([]byte, error)
	Decode(data []byte, numBytes int) (WireMessage, error)
}
