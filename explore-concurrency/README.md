# The Go Memory Model

The Go Memory model is a specification that has a set of conditions. These conditions ensure visibility, i.e., a write to a memory location from one Go routine is visible to the other Go routine that reads the same memory location.

Programs that modify data being simultaneously accessed by multiple goroutines must serialize such access.

To serialize access, protect the data with channel operations or other synchronization primitives such as those in the sync and sync/atomic packages.

A data race is defined as a write to a memory location happening concurrently with another read or write to that same location, unless all the accesses involved are atomic data accesses as provided by the sync/atomic package. 

To avoid a data race, proper synchronization techniques have to be used.


## Happens Before

Within a single goroutine, reads and writes must behave as if they executed in the order specified by the program. That is, compilers and processors may reorder the reads and writes executed within a single goroutine only when the reordering does not change the behavior within that goroutine as defined by the language specification. 

the read r will see the most recent write w when, synchronization events happens-before conditions that ensure reads observe the desired writes.

write w 

    a = 2
    mutex.Lock()
    b = 3
    mutex.Unlock()

read r

    mutex.Lock()
    localvarB = b
    mutex.Unlock()
    localvarA = a // a will read the value 2

i.e., write w to varible a happend before synchronization event i.e write to b. Therefore after the synchronization event of reading b, ensures that latest value a is read. This is beacuse no write operations has happend aftrwer w and beore read r. 

### Initialization

- Program initialization runs in a single goroutine and new goroutines created during initialization do not start running until initialization ends.

- If a package p imports package q, the completion of q's init functions happens before the start of any of p's.

- The start of the function main.main happens after all init functions have finished.

- The execution of any goroutines created during init functions happens after all init functions have finished.

### Once 
A single call of f() from once.Do(f) happens (returns) before any call of once.Do(f) returns.

- sync.Once uses a mutex and an internal flag to ensure that the function passed to Do is executed only once, even if multiple goroutines call Do concurrently.

- The first goroutine to acquire the mutex will execute the function.
All other goroutines will wait until the first goroutine completes the function execution.

- Once the function has been executed, subsequent calls to Do will return immediately without executing the function again.


## Go routine

The go routine follows a fork-join thread model. The child (callee) thread runs independently from the parent (caller) thread and joins the parent later at some point in time when the child completes its execution. if the parent thread exits before the child joins, the child thread will be aborted.<br>

Note: Go doesn't have an explicit join method. But it can be achieved using sync.WaitGroup methods

### References: 
- https://www.cs.cmu.edu/afs/cs.cmu.edu/academic/class/15440-f11/go/doc/go_mem.html
- https://go.dev/ref/mem
