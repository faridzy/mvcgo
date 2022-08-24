/*
 * MVCGo
 * Copyright (c) 2021.
 * author Muhammad Farid H
 */

package config

import (
	"mvcgo/systems/handler"
	"sync"

	"github.com/go-pg/pg/extra/pgdebug"
	"github.com/go-pg/pg/v10"
)

type Database struct {
	*pg.DB
}

var db *Database
var once sync.Once

func Db() *Database {
	once.Do(func() {
		var n int
		configuration := New()
		host := configuration.Get("DB_HOST")
		name := configuration.Get("DB_DATABASE")
		username := configuration.Get("DB_USERNAME")
		password := configuration.Get("DB_PASSWORD")

		conn := pg.Connect(&pg.Options{
			Addr:         host,
			User:         username,
			Password:     password,
			Database:     name,
			PoolSize:     20,
			MinIdleConns: 200,
		})
		_, err := conn.QueryOne(pg.Scan(&n), "SELECT 1")

		handler.PanicIfNeeded(err)
		conn.AddQueryHook(pgdebug.DebugHook{
			// Print all queries.
			Verbose: true,
		})

		db = &Database{conn}
	})
	return db

}
