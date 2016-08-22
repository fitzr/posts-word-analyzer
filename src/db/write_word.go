package db

import (
    "log"
    "database/sql"
    "../writer"
    _ "github.com/go-sql-driver/mysql"
)

const (
    createStemTable = `
CREATE TABLE IF NOT EXISTS word_stem (
  word VARCHAR(255) NOT NULL PRIMARY KEY,
  stem VARCHAR(255) NOT NULL
)`
    createAttrTable = `
CREATE TABLE IF NOT EXISTS word_attr (
  word VARCHAR(255) NOT NULL PRIMARY KEY,
  frequency DOUBLE NOT NULL,
  part_of_speech TEXT NOT NULL
)`
    insertStem = "INSERT INTO word_stem (word, stem) VALUES (?, ?)"
    insertAttr = "INSERT INTO word_attr (word, frequency, part_of_speech) VALUES (?, ?, ?)"
)

type WordWriter interface {
    writer.Writer
    Close() error
}

type wordWriterConn struct {
    *sql.DB
    createTableSql string
    insertSql string
}

func OpenStemWriter(dataSourceName string) WordWriter {
    return openWriter(dataSourceName, createStemTable, insertStem)
}

func OpenAttrWriter(dataSourceName string) WordWriter {
    return openWriter(dataSourceName, createAttrTable, insertAttr)
}

func openWriter(dataSourceName, createTableSql, insertSql string) WordWriter {
    db, err := sql.Open("mysql", dataSourceName)
    if err != nil {
        log.Fatal("open db failed : ", err)
    }

    conn := &wordWriterConn{db, createTableSql, insertSql}
    conn.createTable()
    return conn
}

func (c *wordWriterConn) createTable() {
    _, err := c.Exec(c.createTableSql)
    if err != nil {
        log.Fatal("create table failed : ", err)
    }
}

func (c *wordWriterConn) Write(args ...interface{}) {
    _, err := c.Exec(c.insertSql, args...)
    if err != nil {
        log.Fatal("write failed : ", err)
    }
}