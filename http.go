package nutshttp

import (
	"github.com/xujiajun/nutsdb"
)

func Enable(db *nutsdb.DB) error {
	s := NewNutsHTTPServer(db)

	return s.Run(":8080")
}
