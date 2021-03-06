package storage

import "github.com/theupdateframework/notary/tuf/data"

// MetaUpdate packages up the fields required to update a TUF record
type MetaUpdate struct {
	Role    data.RoleName
	Version int
	Data    []byte
}
