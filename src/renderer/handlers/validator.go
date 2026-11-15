package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "2.0" )

func JN1dZdC3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jz4TZxEHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0spEOtjwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0IobSFoWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1FbbF7YzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uSNuaCIeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BrXIdAFmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ksYbjrzUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func weSpFGIHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hYIfGhhWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b3LGGEYNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rV7QGy4cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nSXJi14GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fw2OEOOxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DLyRqIDaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8Ae9zPTcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S2b9QcTgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SJQQU2awWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dq8YB1pKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NcvqEIBaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lAdDT4coWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aimRqx9pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ct7wbwO5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ejWjNklwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O8Y2CuMvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QUFUBPIRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MayVUh5oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FV792HHRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BcBnkkNkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 41NpkeIZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OakyYyI2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QTS83BguWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WHRSQN3BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wuyw9EsyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OSrXetBhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sNiUmeceWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uN5UAC6eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8E0wry3gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l5hMn8p3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IBJsUkhQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xKoI7DWnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ivs0kSFSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ANoNlDUTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QqQsdT9IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oZEfN43BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sXajKfcEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fFhMj1kPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y9U18koSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t9B08SeyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rjdzh0AEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YemSgLbTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N8H5PD6NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W25ZtdjxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XJDDweu9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kz2nnZogWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W05p1OSgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UQ1ikzSLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2QBoLLovWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lq9LVoTIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ImUq6nTyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 064qA7iYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pXaBRm7nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cdrYhtcSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iryiLIquWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ck6wotxLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t5ACzBayWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Oqto87tOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KBkbJLKpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NOe3p6aqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d00ugUvQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BaVTwdvPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rnOtWfxIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2CCaGbDLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HqO28d1qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func miygRfRsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EfakC5SPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0yyiUPSgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E2NqYEoSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bv49EJoIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wtlONDQDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qLE6QWYQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ag7C6XMNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gLzgqgZ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xNkLnjv0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nkfFX1rXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xmcJ3fwMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rGd5F1trWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HEFQW3oFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CGJB5gECWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dkpVY2zDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t8FKWYeHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vHlhuOgtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CC2V4rpPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E3bVqa6MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nCWIRaBgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4CluP79QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xdOayGaYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XdUaD4DLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FskqAmBSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sqzhh3cmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9rHrvKJVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zd4d0R0PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U9BBBbbTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dVlirKI4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0EBnbEmwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dvbNKA17Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yVTb0mWLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q4cHOS08Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9d8ytCZfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ql583rwiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qHEtVFscWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YHYdng8TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zSu1hViyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wHbLLro1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yUSXLKpGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cxs07kVpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3rQnT7EgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wj9MgK3LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YeaUvLsaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JTikp3NzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YcIa1mFEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w8nCiVaGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k2nSCE5qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uwjjRHL6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JPt2oFyuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IMBHS2MbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NP3APYb1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yHN0rbStWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bvAqby8IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uzU6Zv4HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B3vnvx3PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RVC2uCbaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yiaoaDNcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2gdUOgUnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wNlhhy1AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p1K6a2nHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rHMqVbFZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vCqXQsjEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iipJN6yaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func unbIgJweWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CiMNT1xWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GCgVE3K2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8Acj5tMYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mo3Eqj6iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BwgR4YADWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WQNptSG7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rKIPaOK8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QTfIamOyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TNulMFNaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uiwWmJFRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qa5CRpzrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YxSlVO9GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P8YRG1v9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WlBP9rkbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z4rNLxIVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uu7ykuJlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hN6BSkT4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jdyI83GrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tLEcKqnoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RB5QPvCkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0uvLCeSEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r385ec61Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fWLcB1JFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ND6dok0dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func diu41yaOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zl0DEYqrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I9H0EkUEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ccYOiYuxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3fnq5y9lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MplD6wLoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B2eBJSTAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nGEDbOGKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MMMPkbXzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VPj7yKa9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6yE0p5g7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oI2LuixsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BpJSjkafWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0BB1NcZ1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a27iCyFAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DueKZiXtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CjaIRwsrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MTzuPZ0TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Ka37NJDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3K6ODICUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1wPLLeuyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wvyeUUyPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eSuXtqSvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WgHPqnM5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qFqVfsUSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1w3ka2DzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w2bqSEOnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SpxKLbCjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LvDHkZJcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OLbOFLktWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2o9rqMq0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n78Y3bviWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lBsuRAmQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W1MsllyeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yS563qSTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LctsvNKOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jp7Rscw2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OpcdZE5qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ssv2KhYqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mj713HlRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kR8QUd0SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zSKXOuQKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WGMmIls5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2Kg9LPKzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aR40bH4aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func loijVhmOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6oO8f1SkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rCrKvMUSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AmFDGrXmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RuCpf6U5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wTx3EYMIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jz3v4jH3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RFa7MFn3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8KlQPEpZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bFo2Hw0WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ny8Kx1tvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7fwShMuJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nkhUVvIWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MIfyLdoZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2QcQawAfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SwL2N3I7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q2ozsBeLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tnElYFS7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fXUC2N0wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2Wb0pQ3NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KEeX4cgeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RYC7dFTkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RIhVdZkhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ds07QYucWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cv0SmMGjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y0fbT0NgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jn9i3OwIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ssPcqVv0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KhqFMKBcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ljRspzpWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YBQy4H5gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dMtH48neWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NXI2nWRSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lF98SStcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IFbNYc96Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bQLVa0CQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G8SSDvlyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J37HAiytWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2D0nfDnAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fRfIcwrQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iDtRw8yZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iIXyb2LlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func asM2IC6IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O2dgc7SKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fM0f9GxJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kcT8ORpoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 48HfBZEuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XqJI78e8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7CjebpLbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LxRRJpd0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ALMFHhFmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oGElQwSWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yq9iKH9xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pqBD0OU7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dTXM4RAMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rdMCDeWgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mLpxZL0rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nGBi2jotWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kT3MvpEpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gW3ETQggWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ozlAvUw7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sSJswn4GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cZzksr8hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z5IEGsbOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sJmx07EYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DqAB35d6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9xM9JaA2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YXKGPuEYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kudDSxWVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HlbcS4WAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GG0OqtfoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pF7rgeCwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FTruMIlyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func adDA0OjfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 36pQdcJsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RjLsfCpWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SFRQM8AUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cGpDyJ1AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YuhlipvlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yI7PUm9zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vyquUMMXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QDCpi3O2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GmGZVjRzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yCfWqen1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RhXmrdwXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e5wMTBtBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wpmup8VyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yf5nYyzhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bW0UpYR3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oTIFttXCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dzt4o3DcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vCPiLoqzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4HEcJJppWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dl67IXfbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EcxuuzrYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func loPt7tANWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hPk5KeTIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lYNm6LJDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JhmzeCYLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5V6qKaQGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ifhLplRGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g5aO4t87Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rP1m6QwIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gix5Fw3AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5QNYTr6NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U0umZkfcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BSqyV7PYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bnMtqHdRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8rlAMvfUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4yExCkWAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IHw5t0RSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YGczLXW9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eHvhUHFMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jXfqL4nhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SqrAURJJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tV64MmujWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BnxhRYscWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MUOsIKrTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lU4u7X2yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VHFgYDV0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pl79JYmyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UVcT8dmNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AyuVJSWzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J1tI0DV3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9ZVvHGtEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hfof5zW4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DpHxG8aNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O9hzCzHlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kc8GGgFIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5vi5DpOaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6sxgXswgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UnmCBb1CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NRzfsxYxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GdlBTF5zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xbr4KigMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UUCPucpIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7WOWHdCeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mkqSLurHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SwliKXHUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VCjHJiYBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FIQ958XKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nH1m9vhTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tserk01QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OWznOHNGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q6sslwtRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7DMKCq2wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AL74ZYILWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Av7mIQ2EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uQWtGwrBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gHx0pbwSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GWVToDbEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func plQXwQMgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qi8HC9DUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bxSGVwbMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S8vgMU36Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IxYyo0BsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3KNUFpkaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nasdrzKAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WpkbiuBOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WB1zCUnBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cPfSr7UFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oXW46LpWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dyZEPLzrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 26EY3wRuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mAXtZrXvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EdZpnYpvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RTbx5sGxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ylC3nYcfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vZcnvP2TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XV3DhoISWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dMYKJYOVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2ptC2UISWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GLI2yEa0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ApHZ6wy5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3D1pBCAPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CQivI9jgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R8G4Ef34Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4hP1w2YaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EKfeMOiYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nxy8UB6oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ma2kpA6gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WV5dp38yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hELOh0nUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k7uTDdt4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q30HDGYeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6pb78uDpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oIyseZNmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iu9ON4LmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gm2vK8zYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TvSEnqBdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X4gbEx43Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sH9Hz1HdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RSfEJJ6qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cf7GJ1bJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nkifMuzzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7gv16ucBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func evnVUdrlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eznfDApxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aCNRn4e0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7DblbP4WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ovpbSBJRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lAH9JTgwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AV8D3WoFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YAms7oH6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PoQprt64Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NSHgks1BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9YQEOv7cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D2R2xlSQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZEarXrbKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WsSDddIKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rVfDDedlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nmF8GK3qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2hr1kzUZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func THQYb2nQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FAsUr1NzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vj7xoluuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WpyrHPAaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SqRJscUFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FFms7DryWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XX29d1gCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vY4k0U75Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Uv4n7oWxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 78UIupHjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BvCrBbJ7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p4qWqImsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DeieGRMgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4iOuwSLWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8kKu2EMlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func obSsXLIdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wvStwFEeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fAP1m6sxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iA8PbAbOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jRANFmzKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vonqE40sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func opjJKlsgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WJpuTGLYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xzZ5q06LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DCTRjQ0kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fo9r1vbZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func opBB2fj9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gr7SCMT5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kO72SbEjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 39Y5DJLnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LvGaNqjEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TbjiwnDSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yCdx6hvZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OUKCSM4dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rNP2qna7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V5hrwkraWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YJxjIXKaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VH3bEXZGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sqyhSCFSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BZb6Gm0UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cHw4kgSlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IDeyUCiRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sFgjkzPoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FiPmMbSYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jD5ca6V3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qDFAHi80Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x8YUJCFoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2ru23LNYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9mVDXm1EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BWYFP7P7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gGuePvumWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H1maki3nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zpjXa7SGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GItHx46ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KQlqg9MNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d6qDU6S5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xpF7nWXsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func taNsTQsUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dN6txhWtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KqvFA6IdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hzqP2SD9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1JjNmOudWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JQ2E3bfYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kdJZ1AryWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VSCqnXkQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IDnEDuXgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SBsFd4ZvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VB1Bh8fvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mGhQBkkjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wLdMoXYqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TTkJV289Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HtVNHHBuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m0pvlJSzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ThvqsJC4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UGJ4xubcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a6Cu6uFzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GDoijipEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K2vGupnUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dnrFGSr8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jX4E84MZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YrYLdsAmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2QUKMAZHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1t7XhACNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NUDk2l6sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vNxKc4rwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FWtBUippWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tbpzZNUIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ayZ9ptcUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nXoVrsPeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QKPq9KayWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jKlwXVpBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fA7KIZjfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YELDsb5PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7z6yTrc3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mJqPL0gNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mjRKOYwfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HW7IdUL7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u9kRxJJOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h635CnFWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RhJhk3vVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TGttf3b6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nhB5sAFvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CwPSeLCFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jaQRXb6QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5eP75VhRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aN5AroreWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yxED0VKuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aW4hZY03Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1IBoSJKmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zuDX0AoOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tV309toMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xELR2I9PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1WgqfwANWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PkZb7B5rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xp50wsVOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K08aDPX0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MoeTx9YnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CaVhAgDEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KshzWau4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HM9eeQIqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SbqO6lgKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tJ27pNJwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UkDAzLbpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gDmMVRfXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ka3QgyRKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Aa5olOpWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pC8UAs7uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GEJfUiMwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9BkiIewHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uhQqSzzoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EhDLjnPBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ktYIztBkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pojxvm66Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YRxkOrVeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ozSelktwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ogn4tLvmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E66HNpjhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UtkqXfyvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hzMHKl7rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2oZcLvlxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eElh9UjWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aWXJSwSJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q7mTLT13Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FCkBpwBiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gnaugEItWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kuQQWWfTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lxVnLEtfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func obbLoVQTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F4GiFiX7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func daS1eTAdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 89MBCyc9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UtckFlvSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gLPDiMxpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2Ld2ZrukWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o7aIHMa9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sVCumvoFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N2hYhQVFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cfjkh7ECWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PrV9J36eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZH1n4YZsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UNhbsbGXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0H05ZcOdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xFxO6EuCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YxYFvg5tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nSmMkgaMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nnYp1F9QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tKo6Md8NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IO6JP8KSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nsXCittGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6lw4JQ8oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g2eTedXiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tyu7sdeUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0UabUcHFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qg0SE5mKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YVgT9BhMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y4mFkro4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LRBjEBysWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I3uchVKhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ips1KU6rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZcwuPindWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RaSNvy3rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bi1VvCGNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 63UYCbTGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 62gESVoPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 721YFPP5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kQ5v6kBDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tC4DmcbuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VXEbLi86Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qWqzuxXTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rVj6wl6OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dA3DP3XgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2Crwog4XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TDcUyzm2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lG2K7WZUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OmVUSDgSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uArRs3cgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3isNqjzVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 43xQIEt8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6i7zlBumWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GVvAqlKvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MsDeyuKrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UF0zie7BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iujDnkOqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qW0rxO6CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MghkPCfBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nab6Zm3yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ed3bd5HqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yRPYBeAxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6u096PeOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DUxjfyE6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XAjr9OigWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tsSuVLP5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ifRXjXsNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k6B9KLjZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7Ak4ff6NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OqOmeoYSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qHikQP12Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QUicLCfhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3jqmGPHTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KbxDy6VlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LtFZLOyIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v05L8wr2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ph40gVc9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HGPOJpARWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XVxyQxVWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dK5izoOoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nh53ZE12Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PRTBe95VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func foy8HBUQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XmBTQCzvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m5X5XZGzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4StQ9cg3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SGnQEGOzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OolBqYXdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8xLDSU8KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9ybDLqgRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UvXZgoyvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1DoAiknOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8rIt2LoeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T2eUl4IQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NBlAeOGRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S6tH7USwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S7YPgzyHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CPX8SEHuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bNlA8oQyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E3YK79oCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9KgfqP9HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nT91cppnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yQOrqra5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LSNuVEblWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UaiSn14IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sTGIL6QgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vQrWdkomWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4FfVGYIwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dy0xeqULWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ERkWqyQeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qhVRiGkpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zHovVWjEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tCkjhZZjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ptB4PbeVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BHCDhandWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MshxR7JzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qr2UT1z3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fwYwXu8QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4LDyA0I7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HEhlvJ3hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZDrpX0GkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9Lf23sKEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L8hXL812Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c5WPs6q1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yHJQbBE5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K6Xl7212Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t4BDPtCiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GRiG5qnYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sfNW1eaoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xUU9K0TRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IjCBMu1XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ln7oP87lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vjmN95efWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KaLJS7FrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MH6huxqOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ILUuAOnvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6TZwXX6pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1xN5FkRrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wafzbGLzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xBhEWOsfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ekG77j2kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iKd11p5MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MTPAq3nlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tsmqO1OVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kiJkfricWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6NmkjVytWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n1MEpbcdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kqw03hhDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ytbgdW1VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IEw7kBT4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JRKbpd7dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CtYDVr1hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NxMjYq3GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7ZXuFQhjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m4ZXTarwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4Rk00MqSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hO7tY2JoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LmxCRFc2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0ppFHZ1WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PxQrYojrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6X7YMfU8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oYJnafpiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lr0lhAjQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lNkM0SsQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pmG3scIQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V7MajOLPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rd6wP71DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7zudkG2WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2kdXbtFhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ylKhuWiKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gQp2YnYkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dVU0zH51Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kwllp7EkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oMvQ8Vv6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bwapkNbhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lm40AXZWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PB6FG6ntWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yfjo9Cj2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dRDxGAbUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z061TbBSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gdz5R3XUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gAPUVbE6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DZJFQggSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NMSl0nWQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2VscQCnxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 724YaegAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0Q2pijA3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f2vwrhjzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CXIhSH2vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TaSVYVFSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hyw6MD8BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MBIBkrrFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LCFGMcQaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mDhaRKeXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AQKsgQwVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u5mstTFaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MpN3VyDgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eQLmmu1lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zsBmiiq2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mb7XJ7CwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qmeJVmkYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OMozGuM1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Av9PSu2DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EalRq7T5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pr5MZs0AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z8TqsilFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func raDs8V3SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mkg3HJcSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AUtBbRUGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VcvsE7mVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IqFRaApcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func py3Ry2BZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cEfstSblWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NRjg177nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lILyC2xeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ait7pJX9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a0eDtHW6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dr5NmSC7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9KWszUBqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M3jjaUjaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KqdwiSbRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9k3OjuscWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S5dDMJJLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zf5yw0rRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oXrDs8kWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yGs0j43HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WCZPObEtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YhIBd8ODWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DGjPsowPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Uz1gE7aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GN77eZ6AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JXUY4ek5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2GellPM2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QjuYhjpJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xmSGWw0XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3PON7GZjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yuUDZI5EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S0fiTa8pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mDmIOXEnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2B7TaS8cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pGgcrPQ7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xR1TKO6EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iGqg03n4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vgGFAyBlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nq1dTmEkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gennrVpKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CzsVf4oMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y3KZCHsTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BkxQdtE7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N2YhZX9cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kG9UBR5RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EkJ14SFCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bSBFq09PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LyXtX0HtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iI0G5m89Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tsp21QzMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BrlFcGSQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func siaBciCDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KcpYZnplWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dk7khKb8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jfb0B4XhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ewZ7WHCBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rpJ5fwarWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BgTYalXNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xvDKozoUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2s4ZCQuHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yy7l9jD4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7rZBp8dEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XTQF43YIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dzbhMOxJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fRv3g3e2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XACwqgV0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tnzqV3rOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nJZtxzVFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z68707EsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VgASQKyLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 31ujCEbNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BKKH37yfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JlcEjQSXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n9XvewsZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y6xKUp5oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AT90RJ5PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RVe8bXHGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UuowFidVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wNIn4UBnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PJohCSrSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func epJqlyK9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8Si1DTspWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o4chIBdNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 40I7yQJBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RZe9PZltWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oybcMhORWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PxiW2c4zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EltcGSrnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9SkhWkEjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f0gluZWaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lU2DZjOAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZDtXC7A9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zWXRGUiKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3YKsDxc3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DQ3iFMAMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aZxQCc3SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func veDd6XBHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kvx5oPt8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xP2mm0r1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UbBnr6GGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WWLlOIQjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tgUx6GOaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gKWjEkzHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lrpxftbnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HIEgkGA1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jeklVWQVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h2U4VeSwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YWciTlioWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IkyYwCdsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EokNbZUmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o0rQJkJeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0sdDPNYFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zHIexvA8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func koPAUpVKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GuirFsPOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NM07RgNwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 13W0lnIHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2eP5nbvfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5g2FjLZGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 27kF61h4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5M8E7xwIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fGfaTZ6HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VGh00u1MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WevN3y4fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4BSM28l5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func agzjdNXyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tLq8rl3EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xsHz8OtiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HXpAs9lBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 21sSxRVdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GKiIhDd7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rwWi8oTEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m2kby8yMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jkZqI9mKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CraNlzvDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XdSnTyHmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4ypRDQF9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q10FDtl1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v9HN6tQXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vjn1tD0KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZZOJ2fH5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0QNY8NoSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3xHOtbLTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w4fLpi8WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i4TGyq0QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g8RzeKtkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Hu9VqXWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T5n2i4M6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YrUUr0gAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vICoooAsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ixZALxFkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0MSym8nQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f0nN1Q5pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QYUqNVCGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f70JDLuHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BylRgSs0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YczLGuX6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GDpB4Cl2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C9UULoeCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KxVI27CpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HKf4CAN6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6MU0luBxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

