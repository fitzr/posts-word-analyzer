package task

import (
    "../reader"
    "../writer"
    "github.com/kljensen/snowball"
    "log"
)

type stemPair struct {
    word string
    stem string
}

var (
    channelSize = 100

    wordChannel chan string
    stemChannel chan stemPair
    finished chan bool
)

func MapStem(r reader.Reader, w writer.Writer) {

    log.Println("start")

    initialize()

    go read(r)
    go stem()
    go write(w)

    <- finished

    log.Println("finished")
}

func initialize() {
    wordChannel = make(chan string, channelSize)
    stemChannel = make(chan stemPair, channelSize)
    finished = make(chan bool, 1)
}

func read(r reader.Reader) {
    progress, end := logger("read")
    defer end()
    defer close(wordChannel)

    for {
        word := r.Read()
        if word == "" {
            break
        }
        wordChannel <- word
        progress()
    }
}

func stem() {
    progress, end := logger("stem")
    defer end()
    defer close(stemChannel)

    for word := range wordChannel {
        stem, _ := snowball.Stem(word, "english", true)
        stemChannel <- stemPair {word, stem}
        progress()
    }
}

func write(w writer.Writer) {
    progress, end := logger("write")
    defer end()
    defer func() { finished <- true }()

    for pair := range stemChannel {
        w.Write(pair.word, pair.stem)
        progress()
    }
}