package db

import (
    "testing"
    "database/sql"
    "log"
)

const (
    dataSourceName = "test_user@tcp(127.0.0.1:13306)/stack_test"
)

func TestWriteStem(t *testing.T) {
    defer tearDown()

    sut := OpenStemWriter(dataSourceName)
    word := "has"
    stem := "have"

    sut.Write(word, stem)


    // query result
    conn, err := sql.Open("mysql", dataSourceName)
    checkErr(err, "open connection failed (verify)")
    defer conn.Close()

    rows, err := conn.Query("SELECT word, stem FROM word_stem")
    checkErr(err, "query failed")
    defer rows.Close()

    var actualWord, actualStem string
    if rows.Next() {
        err := rows.Scan(&actualWord, &actualStem)
        checkErr(err, "scan failed")
    }

    // verify
    if word != actualWord {
        t.Errorf("\nexpected: %v\nactual: %v", word, actualWord)
    }
    if stem != actualStem {
        t.Errorf("\nexpected: %v\nactual: %v", word, actualStem)
    }
}

func tearDown() {
    conn, err := sql.Open("mysql", dataSourceName)
    checkErr(err, "open connection failed (tearDown)")
    defer conn.Close()

    _, err = conn.Exec("DROP TABLE word_stem")
    checkErr(err, "drop table failed")
}

func checkErr(err error, msg string) {
    if err != nil {
        log.Fatal(msg, " : ", err)
    }
}
