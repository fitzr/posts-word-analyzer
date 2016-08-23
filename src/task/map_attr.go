package task

import (
    "../reader"
    "../writer"
    "../service"
    "log"
)

var (
    attrChannelSize = 100
    attrParallelNum = 5
)

type mapAttrTask struct {
    wordChannel chan string
    attrChannel chan attr
    semaphore chan bool
    finished chan bool
}

type attr struct {
    word string
    frequency float64
    partOfSpeech string
}

func MapAttr(r reader.Reader, w writer.Writer) {

    log.Println("start")

    t := newMapAttrTask()

    go t.read(r)
    go t.getAttr()
    go t.write(w)

    t.waitToFinish()

    log.Println("finished")
}

func newMapAttrTask() mapAttrTask {
    return mapAttrTask {
        wordChannel: make(chan string, attrChannelSize),
        attrChannel: make(chan attr, attrChannelSize),
        semaphore: make(chan bool, attrParallelNum),
        finished: make(chan bool, 1),
    }
}

func (t *mapAttrTask) read(r reader.Reader) {
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

func (t *mapAttrTask) getAttr() {
    progress, end := logger("get")
    defer end()
    defer close(t.semaphore)

    for word := range t.wordChannel {
        t.semaphore <- true
        go t.getEachAttr(word)
        progress()
    }

    t.semaphore <- false
    go t.getEachAttr("")
}

func (t *mapAttrTask) getEachAttr(word string) {
    if word != "" {
        frequency, partOfSpeech := service.GetAttr(word)
        t.attrChannel <- attr{word, frequency, partOfSpeech}
    }
    if ! <- t.semaphore {
        close(t.attrChannel)
    }
}

func (t *mapAttrTask) write(w writer.Writer) {
    progress, end := logger("write")
    defer end()
    defer func() { t.finished <- true }()

    for attr := range t.attrChannel {
        w.Write(attr.word, attr.frequency, attr.partOfSpeech)
        progress()
    }
}

func (t *mapAttrTask) waitToFinish() {
    <- t.finished
}