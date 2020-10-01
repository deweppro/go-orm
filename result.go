/*
 * Copyright (c) 2020 Mikhail Knyazhev <markus621@gmail.com>.
 * All rights reserved. Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package orm

type Result struct {
	Err  error
	Rows int64
}

func (_result *Result) Error() error        { return _result.Err }
func (_result *Result) RowsAffected() int64 { return _result.Rows }
