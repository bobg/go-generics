package set

import (
	"encoding/json/jsontext"

	"github.com/bobg/seqs"
)

// MarshalJSONTo implements ["encoding/json/v2".MarshalerTo].
func (s Of[T]) MarshalJSONTo(enc *jsontext.Encoder) error {
	return seqs.JSONMarshalEncode(enc, s.All())
}

// UnmarshalJSONFrom implements ["encoding/json/v2".UnmarshalerFrom].
func (s *Of[T]) UnmarshalJSONFrom(dec *jsontext.Decoder) error {
	seq, errptr := seqs.JSONUnmarshalDecode[T](dec)
	if *s == nil {
		*s = New[T]()
	}
	s.AddSeq(seq)
	return *errptr
}
