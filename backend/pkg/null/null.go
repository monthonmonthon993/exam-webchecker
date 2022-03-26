package null

import "encoding/json"

type NullString struct {
	String string
	Valid  bool
}

type NullInt64 struct {
	Int64 int64
	Valid bool
}

func NewString(s string) NullString {
	return NullString{String: s, Valid: true}
}

func NewInt64(i int64) NullInt64 {
	return NullInt64{Int64: i, Valid: true}
}

// MarshalJSON for NullInt64
func (ni *NullInt64) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ni.Int64)
}

// UnmarshalJSON for NullInt64
func (ni *NullInt64) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		ni.Valid = false
		return nil
	} else {
		err := json.Unmarshal(b, &ni.Int64)
		ni.Valid = (err == nil)
		return err
	}
}

// MarshalJSON for NullString
func (ns *NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.String)
}

// UnmarshalJSON for NullString
func (ns *NullString) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		ns.Valid = false
		return nil
	} else {
		err := json.Unmarshal(b, &ns.String)
		ns.Valid = (err == nil)
		return err
	}

}
