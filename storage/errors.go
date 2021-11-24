package storage

import (
	"errors"
)

var ErrStoreUriEmpty = errors.New("store uri is empty")
var ErrStoreDbEmpty = errors.New("store database is empty")
