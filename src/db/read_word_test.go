package db

import (
    "testing"
    "database/sql"
)

func TestReadWordFromWordCount(t *testing.T) {

    setUpWordCount()

    expected1 := "art"
    expected2 := "go"
    expected3 := ""

    sut := OpenWordReaderFromWordCount(testDataSourceName)
    defer sut.Close()

    actual1 := sut.Read()
    if actual1 != expected1 {
        t.Errorf("\nexpected: %v\nactual: %v", expected1, actual1)
    }

    actual2 := sut.Read()
    if actual2 != expected2 {
        t.Errorf("\nexpected: %v\nactual: %v", expected2, actual2)
    }

    actual3 := sut.Read()
    if actual3 != expected3 {
        t.Errorf("\nexpected: %v\nactual: %v", expected3, actual3)
    }
}

func TestReadWord(t *testing.T) {

    setUpWordCountStem()

    expected1 := "go"
    expected2 := ""

    sut := OpenWordReaderFromWordCountStem(testDataSourceName, 1, 1)
    defer sut.Close()

    actual1 := sut.Read()
    if actual1 != expected1 {
        t.Errorf("\nexpected: %v\nactual: %v", expected1, actual1)
    }

    actual2 := sut.Read()
    if actual2 != expected2 {
        t.Errorf("\nexpected: %v\nactual: %v", expected2, actual2)
    }
}

func setUpWordCount() {
    conn, err := sql.Open("mysql", testDataSourceName)
    checkErr(err, "open connection failed (readWordSetUp)")
    defer conn.Close()

    _, err = conn.Exec("CREATE TABLE IF NOT EXISTS word_count (word VARCHAR(3072) NOT NULL PRIMARY KEY, count INT NOT NULL)")
    checkErr(err, "create table failed")

    _, err = conn.Exec("TRUNCATE word_count")
    checkErr(err, "trancate failed")

    _, err = conn.Exec("INSERT INTO word_count (word, count) VALUES ('go', 1234),('art', 3210),('bar', 999)")
    checkErr(err, "insert failed")
}

func setUpWordCountStem() {
    conn, err := sql.Open("mysql", testDataSourceName)
    checkErr(err, "open connection failed (readWordSetUp)")
    defer conn.Close()

    _, err = conn.Exec("CREATE TABLE IF NOT EXISTS word_count_stem " +
            "(word VARCHAR(3072) NOT NULL PRIMARY KEY, count INT NOT NULL, stem VARCHAR(255) NOT NULL, words TEXT NOT NULL)")
    checkErr(err, "create table failed")

    _, err = conn.Exec("TRUNCATE word_count_stem")
    checkErr(err, "trancate failed")

    _, err = conn.Exec("INSERT INTO word_count_stem (word, count, stem, words) VALUES ('go',1234,'go','go'),('art',3210,'art','art'),('bar',999,'bar','bar')")
    checkErr(err, "insert failed")
}