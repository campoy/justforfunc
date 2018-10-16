package main

import (
	"github.com/sirupsen/logrus"
	sql "gopkg.in/src-d/go-mysql-server.v0"
	"gopkg.in/src-d/go-mysql-server.v0/server"
	"gopkg.in/src-d/go-vitess.v0/mysql"
)

func main() {
	e := sql.NewDefault()
	if err := e.Init(); err != nil {
		logrus.Fatalf("could not initialize server: %v", err)
	}

	cfg := server.Config{
		Protocol: "tcp",
		Address:  "0.0.0.0:3306",
		Auth:     new(mysql.AuthServerNone),
	}
	s, err := server.NewDefaultServer(cfg, e)
	if err != nil {
		logrus.Fatalf("could not create default server: %v", err)
	}

	logrus.Infof("server started on %s", cfg.Address)
	if err := s.Start(); err != nil {
		logrus.Fatalf("server failed: %v", err)
	}
}
