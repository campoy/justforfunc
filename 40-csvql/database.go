package main

import "gopkg.in/src-d/go-mysql-server.v0/sql"

type database struct{}

func (d *database) Name() string { return "csvql" }

func (d *database) Tables() map[string]sql.Table {
	return nil
}
