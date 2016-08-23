package task

import "log"

var logInterval = 10

func logger(msg string) (func(), func()) {
    i := 0
    return func () {
        i++
        if i % logInterval  == 0 {
            log.Print(msg + " : ", i)
        }
    }, func () {
        log.Print(msg + " : ", i, " (finished)")
    }
}
