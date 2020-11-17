# Concurrency

Concurrency is the execution of multiple units/processes in parallel.
One of Go's main advantages is the first class support of concurrency.

Concurrency in Go is handled by lightweight threads called _goroutines_.
Initializing a goroutine is just as easy as prepending the keyword **go** in front of a method.

## Mutex Example

```go
package main

import (
    "fmt"
    "sync"
    "strings"
)

// our custom map structure which includes
// a mutex
type SafeMap struct {
    data map[byte]int
    m *sync.RWMutex
}

func main() {
    // initialize the types we need
    safemap := SafeMap{
        data: make(map[byte]int),
        m: &sync.RWMutex{},
    }
    myString := "This is an example string"

    // remove all whitespace characters in the
    // string
    myString = strings.ReplaceAll(myString, " ", "")

    // create a wait group to wait for
    // all concurrent processes to finish
    wg := &sync.WaitGroup{}

    for i := 0; i < len(myString); i++ {
        // add one job count
        wg.Add(1)
        // process via goroutine
        go func(i int, safemap SafeMap, wg *sync.WaitGroup) {
            // release one job count at the end of this
            // function
            defer wg.Done()
            // lock safemap (write lock)
            safemap.m.Lock()
            if _, ok := safemap.data[myString[i]]; !ok {
                // add the new character into our map
                // since it doesn't exist
                safemap.data[myString[i]] = 1
                // unlock safemap
                safemap.m.Unlock()
                return
            }
            // character already exists in our map
            // increase its count
            safemap.data[myString[i]]++
            // unlock safemap
            safemap.m.Unlock()
        }(i, safemap, wg)
    }

    // wait for all concurrent processes to finish
    wg.Wait()

    // print our results
    for k, v := range safemap.data {
        fmt.Printf("%s: %d\n", string(k), v)
    }
}
```

## Channel Examples

Channels are usually great for non-shared memory resource processes. So for example, if your program is
waiting on multiple concurrent processes or is depending on them for next processing steps, channels should
be a viable solution option.

```go
package main

import (
    "fmt"
    "time"
    "math/rand"
)

func main() {

    // initialize some channels
    channelOne := make(chan int, 1)
    channelTwo := make(chan int, 1)

    // goroutine 1
    go func() {
        time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
        channelOne <- 1
    }()

    // goroutine 2
    go func() {
        time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
        channelTwo <- 2
    }()

    // wait until we receive a value from either goroutines
    select {
    case num := <-channelOne:
        fmt.Printf("Num: %d", num)
    case num := <-channelTwo:
        fmt.Printf("Num: %d", num)
    }

}
```
