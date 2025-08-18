package impl

import (
	"log"

	"github.com/Denio1337/go-wallet-service/internal/storage/contract"
	"github.com/Denio1337/go-wallet-service/internal/storage/impl/postgres"
)

var Impl contract.Storage

func init() {
	var err error
	Impl, err = postgres.New()
	if err != nil {
		log.Fatal("Can't create connection to database")
	}
}
