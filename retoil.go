package retoil


import (
	"github.com/reiver/go-toil"
)


type internalRetoil struct {
	toiler      toil.Toiler
	strategizer Strategizer
}


// New returns an initialized retoiler (which is also a toil.Toiler),
// based on the toiler and strategizer passed as parameters.
func New(toiler toil.Toiler, strategizer Strategizer) toil.Toiler {
	retoiler := internalRetoil{
		toiler:toiler,
		strategizer:strategizer,
	}

	return &retoiler
}


func (retoiler *internalRetoil) Toil() {
	// Retoil!
	//
	// NOTE: This must come first so that it gets called last!!!!!
	defer retoiler.retoil()

	// Notify!
	//
	// NOTE: This must come last so that it gets called first!!!!!
	defer retoiler.notify()

	// Make the toiler toil. (I.e., do work.)
	//
	// This method call is expected to be blocking!
	retoiler.toiler.Toil()
}


func (retoiler *internalRetoil) notify() {

	if panicValue := recover(); nil != panicValue {

		// If we got to this point in the code, then the toiler's Toil()
		// method has panic()ed (rather than returning gracefully).
		//
		// At this point we see if the toiler supports us telling it that its
		// Toil() method panic()ed.
		//
		// We do this by trying to cast it to another type of interface.
		// Specifically, the panickedNotifiableToiler interface.
		//
		// This can be useful for adding in logging, tracking, etc.
		//
		// We do the actual call to the toiler's PanickedNotice() method
		// in a goroutine, since we don't want it to block or panic() here!
		//
		// NOTE THAT THIS IS A POTENTIAL SOURCE OF A RESOURCE LEAK!!!!!!
		if notifiableToiler, ok := retoiler.toiler.(panickedNotifiableToiler); ok {
			go func(notifiableToiler panickedNotifiableToiler){
				notifiableToiler.PanickedNotice(panicValue)
			}(notifiableToiler)
		}

		// Now we re-panic() on the panic()ed value.
		//
		// We do this because notify doesn't actually want to recover,
		// but instead wants to just notify that a panic() happened.
		//
		// THIS IS IMPORTANT!
		panic(panicValue)

	} else {
		// If we got to this point in the code, then the toiler's Toil()
		// method has gracefully returned (rather than panic()ing).
		//
		// At this point we see if the toiler supports us telling it that its
		// Toil() method return (gracefully).
		//
		// We do this by trying to cast it to another type of interface.
		// Specifically, the returnedNotifiableToiler interface.
		//
		// This can be useful for adding in logging, tracking, etc.
		//
		// We do the actual call to the toilerd's ReturnedNotice() method
		// in a goroutine, since we don't want it to block or panic() here!
		//
		// NOTE THAT THIS IS A POTENTIAL SOURCE OF A RESOURCE LEAK!!!!!!
		if notifiableToiler, ok := retoiler.toiler.(returnedNotifiableToiler); ok {
			go func(notifiableToiler returnedNotifiableToiler){
				notifiableToiler.ReturnedNotice()
			}(notifiableToiler)
		}
	}
}


func (retoiler *internalRetoil) retoil() {

	if panicValue := recover(); nil != panicValue {

		// If we got to this point in the code, then the toiler's Toil()
		// method has panic()ed (rather than returning gracefully).
		//
		// Now we check with this retoiler's strategizer to see if we may
		// run the toiler's Toil method again (i.e., retoil).
		//
		// If it says "yes", then we invoke the toiler's Toil method again.
		//
		// If it says "no", then with panic() (using the same panic() value
		// we recovered).
		//
		// NOTE that MayRecoveredRetoil() can do things like sleep, if it
		// wants to.
		if retoiler.strategizer.MayPanickedRetoil(panicValue) {
			// NOTE THAT THIS IS A POTENTIAL SOURCE OF A RESOURCE LEAK!!!!!!
			if notifiableToiler, ok := retoiler.toiler.(recoveredNotifiableToiler); ok {
				go func(notifiableToiler recoveredNotifiableToiler){
					notifiableToiler.RecoveredNotice(panicValue)
				}(notifiableToiler)
			}
			retoiler.Toil()
		} else {
			panic(panicValue)
		}
	} else {
		// If we got to this point in the code, then the toiler's Toil()
		// method has gracefully returned (rather than panic()ing).
		//
		// Now we check with this retoiler's strategizer to see if we may
		// run the toiler's Toil method again (i.e., retoil).
		//
		// If it says "yes", then we invoke the toiler's Toil method again.
		//
		// NOTE that MayTerminateRetoil() can do things like sleep, if it
		// wants to.
		if retoiler.strategizer.MayReturnedRetoil() {
			retoiler.Toil()
		}
	}
}
