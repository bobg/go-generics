package set

import (
	"encoding/json/jsontext"

	"github.com/bobg/seqs"
)

func (s Of[T]) MarshalJSONTo(enc *jsontext.Encoder) error {
	return seqs.JSONMarshalEncode(enc, s.All())
}

func (s *Of[T]) UnmarshalJSONFrom(dec *jsontext.Decoder) error {
	seq, err := seqs.JSONUnmarshalDecode[T](dec)
	if *err != nil {
		return *err
	}

	if *s == nil {
		*s = make(Of[T])
	}

	s.AddSeq(seq)
	return *err
}
