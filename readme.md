Why are map operations not defined to be atomic?

Why is len a function and not a method?

Why don't maps allow slices as keys?

What's the difference between new and make?

In short: new allocates memory, make initializes the slice, map, and channel types.

The method set of any other named type T consists of all methods with receiver type T.
The method set of the corresponding pointer type *T is the set of all methods with receiver *T or T (that is, it also contains the method set of T).


return with defer

取决于该变量是不是事先已经存在.
如果不是事先存在, 则会受影响


It's easy to address the situation, though. Use the blank identifier to let unused things persist while you're developing.
use goimports -w

go get -v golang.org/x/tools/cmd/goimports
goimports -w nil_ex.go 


import "unused"

// This declaration marks the import as used by referencing an
// item from the package.
var _ = unused.Item  // TODO: Delete before committing!

func main() {
    debugData := debug.Profile()
    _ = debugData // Used only during debugging.
    ....
}


slice vs array
An array type definition specifies a length and an element type
An array's size is fixed; its length is part of its type 

A slice literal is declared just like an array literal, except you leave out the element count:



One option is to constantly poll the source and check if the channel is closed or not, but that isn’t very efficient.
Go provides the range keyword which when used with a channel will wait on the channel until it is closed. 


The algorithm with which the select keyword works on blocks can be approximated to this:
* check each of the case blocks
* if any one of them is sending or receiving, execute the code block corresponding to it
* if more than one is sending or receiving, randomly pick any of them and execute its code block
* if none of them are ready, wait
* if there is a default block, and none of the other case blocks are ready, then execute the default block (I’m not very sure about this, but from coding experiments, it appears that default gets last priority)


The multi-valued assignment form of the receive operator reports whether a received value was sent before the channel was closed.

The <- operator associates with the leftmost chan possible:
chan<- chan int    // same as chan<- (chan int)
chan<- <-chan int  // same as chan<- (<-chan int)
<-chan <-chan int  // same as <-chan (<-chan int)
chan (<-chan int)


Execution of a "select" statement proceeds in several steps:

1. For all the cases in the statement, the channel operands of receive operations and the channel and right-hand-side expressions of send statements are evaluated exactly once, in source order, upon entering the "select" statement. The result is a set of channels to receive from or send to, and the
corresponding values to send. Any side effects in that evaluation will occur irrespective of which (if any) communication operation is selected to proceed. Expressions on the left-hand side of a RecvStmt with a short variable declaration or assignment are not yet evaluated.
2. If one or more of the communications can proceed, a single one that can proceed is chosen via a uniform pseudo-random selection. Otherwise, if there is a default case, that case is chosen. If there is no default case, the "select" statement blocks until at least one of the communications can
proceed.
3. Unless the selected case is the default case, the respective communication operation is executed.
4. If the selected case is a RecvStmt with a short variable declaration or an assignment, the left-hand side expressions are evaluated and the received value (or values) are assigned.
5. The statement list of the selected case is executed.

Since communication on nil channels can never proceed, a select with only nil channels and no default case blocks forever.


case i3, ok := (<-c3):  // same as: i3, ok := <-c3
	if ok {
		print("received ", i3, " from c3\n")
	} else {
		print("c3 is closed\n")
	}

Sends on a closed channel panic, so it's important to ensure all sends are done before calling close. 


because a receive operation on a closed channel can always proceed immediately, yielding the element type's zero value.


Here are the guidelines for pipeline construction:

1. stages close their outbound channels when all the send operations are done.
2. stages keep receiving values from inbound channels until those channels are closed or the senders are unblocked.

Pipelines unblock senders either by ensuring there's enough buffer for all the values that are sent or by explicitly signalling senders when the receiver may abandon the channel.


use done through whole pipeline for broadcast
	for single goroutine, 
	
		defer close(output)
		select {
			case <-done: return
			case xxx: // real case
		}

use sync.WaitGroup as sync tool when spawning multiple goroutines

	var wg sync.WaitGroup
	wg.Add(1) // Add(cnt_routines)
	for i := 0; i < cnt; i++ {
		go func() {
			defer wg.Done()
			select {
				case <-done: return
				case xxx: // real case
			}
		}()
	}	
	
	go func() {
		wg.Wait()
		close(output)
	}()
	

被关闭的信道会禁止数据流入, 是只读的。我们仍然可以从关闭的信道中取出数据，但是不能再写入数据了。
