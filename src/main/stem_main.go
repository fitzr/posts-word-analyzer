package main

import (
    "../db"
    "../task"
    "os"
    "log"
)

func main() {

    // args
    if len(os.Args) < 2 {
        log.Fatal("required arguments : command db_source")
    }
    dataSource := os.Args[1]

    // reader
    reader := db.OpenWordReader(dataSource)
    defer reader.Close()

    // writer
    writer := db.OpenStemWriter(dataSource)
    defer writer.Close()

    // execute
    task.MapStem(reader, writer)
}