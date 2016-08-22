package db

import (
    "database/sql"
    "log"
    "../reader"
    _ "github.com/go-sql-driver/mysql"
)

type WordReader interface {
    reader.Reader
    Close()
}

type wordReaderConn struct {
    db *sql.DB
    rows *sql.Rows
}

const (
    selectFromWordCount = "SELECT word FROM word_count WHERE count >= 1000"
    selectFromWordCountStem = "SELECT word FROM word_count_stem ORDER BY count DESC LIMIT ? OFFSET ?"
)

func OpenWordReaderFromWordCount(dataSourceName string) WordReader {
    return openReader(dataSourceName, selectFromWordCount)
}

func OpenWordReaderFromWordCountStem(dataSourceName string, limit, offset int) WordReader {
    return openReader(dataSourceName, selectFromWordCountStem, limit, offset)
}

func openReader(dataSourceName, querySql string, args ...interface{}) WordReader {
    db, err := sql.Open("mysql", dataSourceName)
    if err != nil {
        log.Fatal("open db failed : ", err)
    }

    rows, err := db.Query(querySql, args...)
    if err != nil {
        db.Close()
        log.Fatal("query word failed : ", err)
    }

    conn := &wordReaderConn{db:db, rows:rows}
    return conn
}

func (c *wordReaderConn) Read() string {
    if c.rows.Next() {
        var word string
        c.rows.Scan(&word)
        return word
    } else if c.rows.Err() != nil {
        log.Fatal("read word failed", c.rows.Err())
    }
    return ""
}

func (c *wordReaderConn) Close() {
    c.rows.Close()
    c.db.Close()
}
