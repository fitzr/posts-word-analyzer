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

var stemChannelSize = 100

type mapStemTask struct {
    wordChannel chan string
    stemChannel chan stemPair
    finished chan bool
}

func MapStem(r reader.Reader, w writer.Writer) {

    log.Println("start")

    t := newMapStemTask()

    go t.read(r)
    go t.stem()
    go t.write(w)

    t.waitToFinish()

    log.Println("finished")
}

func newMapStemTask() mapStemTask {
    return mapStemTask {
        wordChannel: make(chan string, stemChannelSize),
        stemChannel: make(chan stemPair, stemChannelSize),
        finished: make(chan bool, 1),
    }
}

func (t *mapStemTask) read(r reader.Reader) {
    progress, end := logger("read")
    defer end()
    defer close(t.wordChannel)

    for {
        word := r.Read()
        if word == "" {
            break
        }
        t.wordChannel <- word
        progress()
    }
}

func (t *mapStemTask) stem() {
    progress, end := logger("stem")
    defer end()
    defer close(t.stemChannel)

    for word := range t.wordChannel {
        stem, _ := snowball.Stem(word, "english", true)
        t.stemChannel <- stemPair {word, stem}
        progress()
    }
}

func (t *mapStemTask) write(w writer.Writer) {
    progress, end := logger("write")
    defer end()
    defer func() { t.finished <- true }()

    for pair := range t.stemChannel {
        w.Write(pair.word, pair.stem)
        progress()
    }
}

func (t *mapStemTask) waitToFinish() {
    <- t.finished
}