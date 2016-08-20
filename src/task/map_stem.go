package task

import (
    "../reader"
    "../writer"
    "../service"
    "log"
)

type stemPair struct {
    word string
    stem string
}

var (
    channelSize = 100
    parallelNum = 5

    wordChannel chan string
    stemChannel chan stemPair
    semaphore chan bool
    finished chan bool
)

func MapStem(r reader.Reader, w writer.Writer) {

    log.Println("start")

    initialize()

    go read(r)
    go get()
    go write(w)

    <- finished

    log.Println("finished")
}

func initialize() {
    wordChannel = make(chan string, channelSize)
    stemChannel = make(chan stemPair, channelSize)
    semaphore = make(chan bool, parallelNum)
    finished = make(chan bool, 1)
}

func read(r reader.Reader) {
    progress, end := logger("read")
    defer end()
    defer close(wordChannel)

    for {
        word := r.ReadWord()
        if word == "" {
            break
        }
        wordChannel <- word
        progress()
    }
}

func get() {
    progress, end := logger("get")
    defer end()
    defer close(semaphore)

    for word := range wordChannel {
        semaphore <- true
        go stem(word)
        progress()
    }

    semaphore <- false
    go stem("")
}

func stem(word string) {
    if word != "" {
        stem := service.GetStem(word)
        stemChannel <- stemPair{word, stem}
    }
    if ! <- semaphore {
        close(stemChannel)
    }
}

func write(w writer.Writer) {
    progress, end := logger("write")
    defer end()
    defer func() { finished <- true }()

    for pair := range stemChannel {
        w.WriteStem(pair.word, pair.stem)
        progress()
    }
}