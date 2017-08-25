package codecs

import (
	"github.com/spartanlogs/spartan/codecs"
	"github.com/spartanlogs/spartan/event"
	"github.com/spartanlogs/spartan/utils"
)

// The PlainCodec reads plaintext with no delimiting between events
type PlainCodec struct {
	codecs.BaseCodec
}

func init() {
	codecs.Register("plain", newPlainCodec)
}

func newPlainCodec(options utils.InterfaceMap) (codecs.Codec, error) {
	return &PlainCodec{}, nil
}

// Encode event as a simple message.
func (c *PlainCodec) Encode(e *event.Event) []byte {
	return []byte(e.String())
}

// Decode creates a new event with message set to data.
func (c *PlainCodec) Decode(data []byte) (*event.Event, error) {
	return event.New(string(data)), nil
}
