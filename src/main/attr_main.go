package main

import (
    "../db"
    "../task"
    "../service"
    "os"
    "log"
    "strconv"
)

func main() {

    // args
    if len(os.Args) < 5 {
        log.Fatal("required arguments : command db_source mashape_key limit offset")
    }
    dataSource := os.Args[1]
    mashapeKey := os.Args[2]
    limit, err := strconv.Atoi(os.Args[3])
    if err != nil {
        log.Fatal("limit is not int : ", err)
    }
    offset, err := strconv.Atoi(os.Args[4])
    if err != nil {
        log.Fatal("offset is not int : ", err)
    }

    // setting
    service.MashapeKey = mashapeKey

    // reader
    reader := db.OpenWordReaderFromWordCountStem(dataSource, limit, offset)
    defer reader.Close()

    // writer
    writer := db.OpenAttrWriter(dataSource)
    defer writer.Close()

    // execute
    task.MapAttr(reader, writer)
}