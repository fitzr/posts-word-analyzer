package db

import (
    "database/sql"
    "log"
    "../reader"
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
    querySql = "SELECT word FROM word_count WHERE count >= 1000"
)

func OpenWordReader(dataSourceName string) WordReader {
    db, err := sql.Open("mysql", dataSourceName)
    if err != nil {
        log.Fatal("open db failed : ", err)
    }

    rows, err := db.Query(querySql)
    if err != nil {
        db.Close()
        log.Fatal("query word failed : ", err)
    }

    conn := &wordReaderConn{db:db, rows:rows}
    return conn
}

func (c *wordReaderConn) ReadWord() string {
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
