package mysqlrepository

import "errors"

var (
	ErrNotFound = errors.New("record not found")
)
