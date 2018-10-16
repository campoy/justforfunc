package csvql

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"gopkg.in/src-d/go-mysql-server.v0/sql"
)

type Database struct {
	tables map[string]sql.Table
}

func NewDatabase(path string) (*Database, error) {
	fis, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, errors.Wrapf(err, "could not read %s", path)
	}

	tables := make(map[string]sql.Table)
	for _, fi := range fis {
		name := strings.ToLower(fi.Name())
		if fi.IsDir() || filepath.Ext(name) != ".csv" {
			continue
		}
		tables[strings.TrimSuffix(name, ".csv")] = &table{}
	}

	return &Database{tables}, nil
}

func (d *Database) Name() string                 { return "csvql" }
func (d *Database) Tables() map[string]sql.Table { return d.tables }

type table struct {
	name   string
	schema sql.Schema
}

func (t *table) Name() string       { return t.name }
func (t *table) String() string     { return t.name }
func (t *table) Schema() sql.Schema { return t.schema }
func (t *table) Partitions(ctx *sql.Context) (sql.PartitionIter, error) {
	return nil, errors.New("Partitions not implemented")
}
func (t *table) PartitionRows(ctx *sql.Context, p sql.Partition) (sql.RowIter, error) {
	return nil, errors.New("Partitions not implemented")
}
