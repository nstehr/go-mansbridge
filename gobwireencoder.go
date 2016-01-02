package mansbridge

import (
	"bytes"
	"encoding/gob"
)

type GobWireEncoder struct{}

func (g GobWireEncoder) Encode(wireMsg WireMessage) ([]byte, error) {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(wireMsg); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (g GobWireEncoder) Decode(data []byte, numBytes int) (WireMessage, error) {
	msg := WireMessage{}
	if err := gob.NewDecoder(bytes.NewReader(data[:numBytes])).Decode(&msg); err != nil {
		return msg, err
	}
	return msg, nil
}

func (g GobWireEncoder) RegisterType(v interface{}) {
	gob.Register(v)
}
