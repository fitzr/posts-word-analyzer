package db

import (
    "log"
    "database/sql"
    "../writer"
    _ "github.com/go-sql-driver/mysql"
)

const (
    sqlCreateTable = `
CREATE TABLE IF NOT EXISTS word_stem (
  word VARCHAR(255) NOT NULL PRIMARY KEY,
  stem VARCHAR(255) NOT NULL
)`
    sqlInsert = "INSERT INTO word_stem (word, stem) VALUES (?, ?)"
)

type StemWriter interface {
    writer.Writer
    Close() error
}

type stemWriterConn struct {
    *sql.DB
}

func OpenStemWriter(dataSourceName string) StemWriter {
    db, err := sql.Open("mysql", dataSourceName)
    if err != nil {
        log.Fatal("open db failed : ", err)
    }

    conn := &stemWriterConn{db}
    conn.createTableIfNotExists()
    return conn
}

func (c *stemWriterConn) createTableIfNotExists() {
    _, err := c.Exec(sqlCreateTable)
    if err != nil {
        log.Fatal("create table failed : ", err)
    }
}

func (c *stemWriterConn) WriteStem(word, stem string) {
    _, err := c.Exec(sqlInsert, word, stem)
    if err != nil {
        log.Fatal("write stem failed : ", err)
    }
}