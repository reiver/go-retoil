
package retoil


import (
	"testing"

	"github.com/reiver/go-toil/toiltest"

	"sync"
)


func TestNew(t *testing.T) {

	toiler := toiltest.NewRecorder()

	retoiler := New(toiler, NeverStrategy())
	if nil == retoiler {
		t.Errorf("After trying to create a new retoiler, received nil: %v", retoiler)
	}
}


func TestPanickedNotice(t *testing.T) {

	toiler := toiltest.NewRecorder()

	numPanicked := 0

	var waitGroup sync.WaitGroup
	var caughtPanicValue interface{}
	waitGroup.Add(1)
	toiler.PanickedNoticeFunc(func(panicValue interface{}){
		numPanicked++
		caughtPanicValue = panicValue
		waitGroup.Done()
	})

	retoiler := New(toiler, DelayedLimitedStrategy(2, 0))

	if expected, actual := 0, numPanicked; expected != actual {
		t.Errorf("Expected number of times panicked to be %d, but actually was %d.", expected, actual)
	}

	go retoiler.Toil()

	panicValue := "Panic panic, panic!!!!"
	toiler.Panic(panicValue)
	waitGroup.Wait()

	if expected, actual := 1, numPanicked; expected != actual {
		t.Errorf("Expected number of times panicked to be %d, but actually was %d.", expected, actual)
	}
	if expected, actual := panicValue, caughtPanicValue; expected != actual {
		t.Errorf("Expected caught panic value to be [%v], but actually was [%v].", expected, actual)
	}
}



func TestReturnedNotice(t *testing.T) {

	toiler := toiltest.NewRecorder()

	numReturned := 0

	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	toiler.ReturnedNoticeFunc(func(){
		numReturned++
		waitGroup.Done()
	})

	retoiler := New(toiler, DelayedLimitedStrategy(2, 0))

	if expected, actual := 0, numReturned; expected != actual {
		t.Errorf("Expected number of times returned to be %d, but actually was %d.", expected, actual)
	}

	go retoiler.Toil()

	toiler.Terminate()
	waitGroup.Wait()

	if expected, actual := 1, numReturned; expected != actual {
		t.Errorf("Expected number of times returned to be %d, but actually was %d.", expected, actual)
	}
}


func TestToil(t *testing.T) {

	toiler := toiltest.NewRecorder()

	numToiled := 0

	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	toiler.ToilFunc(func(){
		numToiled++
		waitGroup.Done()
	})

	retoiler := New(toiler, DelayedLimitedStrategy(2, 0)) // NOTE that the 2 means we can only panic() this toiler 2 times!

	if expected, actual := 0, numToiled; expected != actual {
		t.Errorf("Expected number of times toiled to be %d, but actually was %d.", expected, actual)
	}

	go retoiler.Toil()

	waitGroup.Wait()

	if expected, actual := 1, numToiled; expected != actual {
		t.Errorf("Expected number of times toiled to be %d, but actually was %d.", expected, actual)
	}



	waitGroup.Add(1)
	panicValue1 := "Panic panic, panic!!!! [1]"
	toiler.Panic(panicValue1)
	waitGroup.Wait()

	if expected, actual := 2, numToiled; expected != actual {
		t.Errorf("Expected number of times toiled to be %d, but actually was %d.", expected, actual)
	}



	waitGroup.Add(1)
	panicValue2 := "Panic panic, panic!!!! [2]"
	toiler.Panic(panicValue2)
	waitGroup.Wait()

	if expected, actual := 3, numToiled; expected != actual {
		t.Errorf("Expected number of times toiled to be %d, but actually was %d.", expected, actual)
	}
}
