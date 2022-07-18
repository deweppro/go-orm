package orm

import "github.com/deweppro/go-errors"

var (
	//ErrInvalidModelPool if sync pool has invalid model type
	ErrInvalidModelPool = errors.New("invalid internal model pool")
)
