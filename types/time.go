/*
 * Copyright (c) 2020 Mikhail Knyazhev <markus621@gmail.com>.
 * All rights reserved. Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package types

import (
	"database/sql"
	"database/sql/driver"
	"strings"
	"time"
)

const timeFormat = "2006-01-02 15:04:05"

type TimeAt sql.NullTime

func (t *TimeAt) Scan(value interface{}) error {
	return (*sql.NullTime)(t).Scan(value)
}

func (t TimeAt) Value() (driver.Value, error) {
	if !t.Valid {
		return nil, nil
	}
	return t.Time, nil
}

func (t TimeAt) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return []byte("null"), nil
	}
	return []byte(`"` + t.Time.Format(timeFormat) + `"`), nil
}

func (t *TimeAt) UnmarshalJSON(jsonData []byte) (err error) {
	str := strings.Trim(string(jsonData), `"`)
	if len(str) == 0 {
		return nil
	}

	t.Time, err = time.Parse(timeFormat, str)
	t.Valid = err != nil
	return
}
