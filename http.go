package nutshttp

import (
	"github.com/xujiajun/nutsdb"
)

var EnableAuth bool

func Enable(db *nutsdb.DB) error {
	s := NewNutsHTTPServer(db)

	return s.Run(":8080")
}
