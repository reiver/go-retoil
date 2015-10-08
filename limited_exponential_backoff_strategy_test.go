
package retoil


import (
	"testing"

	"github.com/reiver/go-toil/toiltest"

	"sync"
)


func TestRetoilWithLimitedExponentialBackoffStrategy(t *testing.T) {

	toiler := toiltest.NewRecorder()


	numToiled := 0
	var toiledWaitGroup sync.WaitGroup
	toiledWaitGroup.Add(1)
	toiler.ToilFunc(func(){
		numToiled++
		toiledWaitGroup.Done()
	})


	numPanicked := 0
	var panickedWaitGroup sync.WaitGroup
	toiler.PanickedNoticeFunc(func(panicValue interface{}){
		numPanicked++
		panickedWaitGroup.Done()
	})


	retoiler := New(toiler, LimitedExponentialBackoffStrategy(2, 0)) // NOTE that the 2 means we can only panic() this toiler 2 times before panic()s get through!

	if expected, actual := 0, numToiled; expected != actual {
		t.Errorf("Expected number of times toiled to be %d, but actually was %d.", expected, actual)
	}

	if expected, actual := 0, numPanicked; expected != actual {
		t.Errorf("Expected number of times panicked to be %d, but actually was %d.", expected, actual)
	}



	var finalPanicValue interface{} = nil
	go func() {
		defer func() {
			if expected, actual := 3, numToiled; expected != actual {
				t.Errorf("Expected number of times toiled to be %d, but actually was %d.", expected, actual)
			}

			if panicValue := recover(); nil != panicValue {
				if expected, actual := panicValue, finalPanicValue; expected != actual {
					t.Errorf("Expected caught panic value to be [%v], but actually was [%v].", expected, actual)
				}
			} else {
				t.Errorf("This should NOT get to this part of the code either!!")
			}
		}()

		retoiler.Toil()
	}()



	toiledWaitGroup.Wait()

	if expected, actual := 1, numToiled; expected != actual {
		t.Errorf("Expected number of times toiled to be %d, but actually was %d.", expected, actual)
	}

	if expected, actual := 0, numPanicked; expected != actual {
		t.Errorf("Expected number of times panicked to be %d, but actually was %d.", expected, actual)
	}



	toiledWaitGroup.Add(1)
	panickedWaitGroup.Add(1)
	panicValue1 := "Panic panic, panic!!!! [1]"
	toiler.Panic(panicValue1)
	toiledWaitGroup.Wait()
	panickedWaitGroup.Wait()

	if expected, actual := 2, numToiled; expected != actual {
		t.Errorf("Expected number of times toiled to be %d, but actually was %d.", expected, actual)
	}

	if expected, actual := 1, numPanicked; expected != actual {
		t.Errorf("Expected number of times panicked to be %d, but actually was %d.", expected, actual)
	}



	toiledWaitGroup.Add(1)
	panickedWaitGroup.Add(1)
	panicValue2 := "Panic panic, panic!!!! [2]"
	toiler.Panic(panicValue2)
	toiledWaitGroup.Wait()
	panickedWaitGroup.Wait()

	if expected, actual := 3, numToiled; expected != actual {
		t.Errorf("Expected number of times toiled to be %d, but actually was %d.", expected, actual)
	}

	if expected, actual := 2, numPanicked; expected != actual {
		t.Errorf("Expected number of times panicked to be %d, but actually was %d.", expected, actual)
	}



	panickedWaitGroup.Add(1)
	panicValue3 := "Panic panic, panic!!!! [3]"
	finalPanicValue = panicValue3 // <----------------- NOTE we set the finalPanicValue
	toiler.Panic(panicValue3)
	panickedWaitGroup.Wait()

	//                     V---------- NOTE that sayed as 3.
	if expected, actual := 3, numToiled; expected != actual {
		t.Errorf("Expected number of times toiled to be %d, but actually was %d.", expected, actual)
	}

	if expected, actual := 3, numPanicked; expected != actual {
		t.Errorf("Expected number of times panicked to be %d, but actually was %d.", expected, actual)
	}
}
