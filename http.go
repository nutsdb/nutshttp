package nutshttp

import (
	"github.com/spf13/viper"
	"github.com/xujiajun/nutsdb"
	"log"
)

func Enable(db *nutsdb.DB) error {
	server, err := NewNutsHTTPServer(db)
	if err != nil {
		log.Fatalln(err)
	}

	port := viper.GetString("port")
	return server.Run(":" + port)
}
