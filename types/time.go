package types

import (
	"database/sql"
	"database/sql/driver"
	"strings"
	"time"
)

const timeFormat = "2006-01-02 15:04:05"

//TimeAt model
type TimeAt sql.NullTime

//Scan implements the Scanner interface.
func (t *TimeAt) Scan(value interface{}) error {
	return (*sql.NullTime)(t).Scan(value)
}

//Value implements the driver Valuer interface.
func (t TimeAt) Value() (driver.Value, error) {
	if !t.Valid {
		return nil, nil
	}
	return t.Time, nil
}

//MarshalJSON implements json encoding interface.
func (t TimeAt) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return []byte("null"), nil
	}
	return []byte(`"` + t.Time.Format(timeFormat) + `"`), nil
}

//UnmarshalJSON implements json decoding interface.
func (t *TimeAt) UnmarshalJSON(jsonData []byte) (err error) {
	str := strings.Trim(string(jsonData), `"`)
	if len(str) == 0 {
		return nil
	}

	t.Time, err = time.Parse(timeFormat, str)
	t.Valid = err != nil
	return
}
