package wire

import (
	"bytes"
	"encoding/gob"
)

type GobCacheEncoder struct{}

func (g GobCacheEncoder) Encode(wireMsg WireMessage) ([]byte, error) {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(wireMsg); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (g GobCacheEncoder) Decode(data []byte, numBytes int) (WireMessage, error) {
	msg := WireMessage{}
	if err := gob.NewDecoder(bytes.NewReader(data[:numBytes])).Decode(&msg); err != nil {
		return msg, err
	}
	return msg, nil
}
