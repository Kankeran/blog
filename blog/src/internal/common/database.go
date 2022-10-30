package common

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/lib/pq"
)

type Database struct {
	host, port, user, pass, dbname, sslmode string
	db                                      *sql.DB
	once                                    *sync.Once
}

func NewDatabase(host, port, user, pass, dbname, sslmode string, once *sync.Once) *Database {
	return &Database{host: host, port: port, user: user, pass: pass, dbname: dbname, sslmode: sslmode, once: once}
}

func (d *Database) Open() *sql.DB {
	d.once.Do(func() {
		psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			d.host, d.port, d.user, d.pass, d.dbname, d.sslmode)
		var err error
		d.db, err = sql.Open("postgres", psqlconn)
		if err != nil {
			panic(err)
		}
	})

	return d.db
}
