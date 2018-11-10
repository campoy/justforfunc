package csvql

import (
	"encoding/csv"
	"io"
	"io/ioutil"
	"os"
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
		t, err := newTable(filepath.Join(path, name))
		if err != nil {
			return nil, errors.Wrapf(err, "could not create table from %s", name)
		}
		tables[strings.TrimSuffix(name, ".csv")] = t
	}

	return &Database{tables}, nil
}

func (d *Database) Name() string                 { return "csvql" }
func (d *Database) Tables() map[string]sql.Table { return d.tables }

type table struct {
	name   string
	schema sql.Schema
	path   string
}

func newTable(path string) (sql.Table, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrapf(err, "could not read %s", path)
	}
	defer f.Close()

	r := csv.NewReader(f)
	headers, err := r.Read()
	if err != nil {
		return nil, errors.Wrapf(err, "could not read headers in %s", path)
	}

	tableName := strings.TrimSuffix(filepath.Base(path), ".csv")

	var schema []*sql.Column
	for _, header := range headers {
		schema = append(schema, &sql.Column{
			Name:   header,
			Type:   sql.Text,
			Source: tableName,
		})
	}

	return &table{name: tableName, schema: schema, path: path}, nil
}

func (t *table) Name() string       { return t.name }
func (t *table) String() string     { return t.name }
func (t *table) Schema() sql.Schema { return t.schema }

type partitionIter struct{ done bool }

func (p *partitionIter) Close() error { return nil }

type partition struct{}

func (partition) Key() []byte { return []byte{'@'} }

func (p *partitionIter) Next() (sql.Partition, error) {
	if p.done {
		return nil, io.EOF
	}
	p.done = true
	return &partition{}, nil
}

func (t *table) Partitions(ctx *sql.Context) (sql.PartitionIter, error) {
	return &partitionIter{}, nil
}

type rowIter struct {
	*csv.Reader
	io.Closer
}

func (r *rowIter) Next() (sql.Row, error) {
	cols, err := r.Read()
	if err == io.EOF {
		return nil, err
	} else if err != nil {
		return nil, errors.Wrap(err, "could not read row")
	}
	row := make(sql.Row, len(cols))
	for i, col := range cols {
		row[i] = col
	}
	return row, err
}

func (t *table) PartitionRows(ctx *sql.Context, p sql.Partition) (sql.RowIter, error) {
	f, err := os.Open(t.path)
	if err != nil {
		return nil, errors.Wrapf(err, "could not open %s", t.path)
	}

	r := csv.NewReader(f)
	r.Read()
	return &rowIter{r, f}, nil
}
