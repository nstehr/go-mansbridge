package wire

type CacheEncoder interface {
	Encode(wireMsg WireMessage) ([]byte, error)
	Decode(data []byte, numBytes int) (WireMessage, error)
}
