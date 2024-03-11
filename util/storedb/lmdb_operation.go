package storedb

import (
	"github.com/bmatsuo/lmdb-go/lmdb"

	"os/user"
	loog "log"
)

var Env *lmdb.Env
var Dbi lmdb.DBI
var Store bool

func OpenDBI() {
	current, err := user.Current()
	if err != nil {
		loog.Fatal("Get current user:", err)
	}

	home := current.HomeDir
	loog.Println("home:", home)

	Env, err = lmdb.NewEnv()
	if err != nil {
		loog.Fatal("Get lmdb ENV:", err)
	}

	err = Env.SetMaxDBs(1)
	if err != nil {
		loog.Fatal("Set Max db num:", err)
	}

	err = Env.SetMapSize(1 << 26)
	if err != nil {
		loog.Fatal("Set Max file storage", err)
	}

	err = Env.Open(home + "/.catspa-explorer", 0, 0644)
	if err != nil {
		loog.Fatal("Open Env:", err)
	}

	err = Env.Update(func(txn *lmdb.Txn) (err error) {
		Dbi, err = txn.CreateDBI("tx-block")
		return err
	})

	if err != nil {
		loog.Fatal("create db:", err)
	}

	if err != nil {
		loog.Println(err)
	}
}
