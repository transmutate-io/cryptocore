package types

import (
	"encoding/json"
	"time"
)

type UnixTime int

func NewUnixTime(t time.Time) UnixTime           { return UnixTime(t.UTC().Unix()) }
func (ut UnixTime) Time() time.Time              { return time.Unix(int64(ut), 0).UTC() }
func (ut UnixTime) String() string               { return ut.Time().Format(time.RFC3339) }
func (ut UnixTime) MarshalJSON() ([]byte, error) { return json.Marshal(int(ut)) }

func (ut *UnixTime) UnmarshalJSON(b []byte) error {
	var t int
	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}
	*ut = UnixTime(t)
	return nil
}
