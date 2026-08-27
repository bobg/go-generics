//go:build go1.27

package set

import (
	"encoding/json/jsontext"
	"encoding/json/v2"
)

func (s Of[T]) MarshalJSONTo(enc *jsontext.Encoder) error {
	if s == nil {
		return enc.WriteToken(jsontext.Null)
	}

	if err := enc.WriteToken(jsontext.BeginArray); err != nil {
		return err
	}

	for k := range s.All() {
		if err := json.MarshalEncode(enc, k); err != nil {
			return err
		}
	}

	return enc.WriteToken(jsontext.EndArray)
}

func (s *Of[T]) UnmarshalJSONFrom(dec *jsontext.Decoder) error {
	switch dec.PeekKind() {
	case jsontext.KindBeginArray:
	case jsontext.KindNull:
		*s = nil
		_, err := dec.ReadToken()
		return err
	default:
		return &json.SemanticError{JSONKind: dec.PeekKind()}
	}

	if _, err := dec.ReadToken(); err != nil {
		return err
	}

	for dec.PeekKind() != ']' {
		var v T
		if err := json.UnmarshalDecode(dec, &v); err != nil {
			return err
		}

		if *s == nil {
			*s = make(Of[T])
		}

		s.Add(v)
	}

	if _, err := dec.ReadToken(); err != nil {
		return err
	}

	return nil
}
