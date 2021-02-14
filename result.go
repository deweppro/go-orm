package orm

//Result model
type Result struct {
	Err  error
	Rows int64
}

//Resulter interface
type Resulter interface {
	Error() error
	RowsAffected() int64
}

//Err ...
func (r *Result) Error() error {
	return r.Err
}

//RowsAffected ...
func (r *Result) RowsAffected() int64 {
	return r.Rows
}
