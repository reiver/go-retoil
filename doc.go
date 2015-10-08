/*
Package retoil provides simple functionality for restarting toilers (i.e., workers).

A toiler that has a Toil() method, that does work, blocks (i.e., doesn't return) until
the work is done, and panic()s if there is a problem it cannot or doesn't want to deal
with.

Usage

To use, create one or more types that implement the toil.Toiler interface. For example:

	type awesomeToiler struct{}
	
	func newAwesomeToiler() {
	
		toiler := awesomeToiler{}
	
		return &toiler
	}
	
	func (toiler *awesomeToiler) Toil() {
		//@TODO: Do work here.
		//
		// And this blocks (i.e., not not return)
		// until the work is done.
		//
		// It also panic()s if it encounters a problem
		// it cannot or doesn't want to deal with.
	}

Then create a retoiler that wraps that toiler. (Also choosing a retoil strategy, when doing that.)

	toiler := newAwesomeToiler()
	
	strategizer := DelayedLimitedStrategy(16, 5 * time.Second)
	//strategizer := LimitedExponentialBackoffStrategy(16, 5 * time.Second)
	
	retoiler := retoil.New(toiler, strategizer)

Observers

A toiler's Toil method can finish in one of two ways. Either it will return gracefully, or
it will panic().

The retoiler is OK with either.

But also, the retoiler provides the toiler with a convenient way of being notified
of each case.

If a toiler also has a ReturnedNotice() method, then the retoiler will call the toiler's
ReturnedNotice() method when the toiler's Toil() method has returned gracefully. For example:

	type awesomeToiler struct{}
	
	func newAwesomeToiler() {
	
		toiler := awesomeToiler{}
	
		return &toiler
	}
	
	func (toiler *awesomeToiler) Toil() {
		//@TODO: Do work here.
	}
	
	func (toiler *awesomeToiler) ReturnedNotice() {
		//@TODO: Do something with this notification.
	}

If a toiler also has a PanickedNotice() method, then the retoiler will call the toiler's
PanickedNotice() method when the toiler's Toil() method has panic()ed. For example:

	type awesomeToiler struct{}
	
	func newAwesomeToiler() {
	
		toiler := awesomeToiler{}
	
		return &toiler
	}
	
	func (toiler *awesomeToiler) Toil() {
		//@TODO: Do work here.
	}
	
	func (toiler *awesomeToiler) PanickedNotice() {
		//@TODO: Do something with this notification.
	}

If a toiler also has a RecoveredNotice() method, then the retoiler will call the toiler's
RecoveredNotice() method when the toiler's Toil() method has restarted after a panic(). For example:

	type awesomeToiler struct{}
	
	func newAwesomeToiler() {
	
		toiler := awesomeToiler{}
	
		return &toiler
	}
	
	func (toiler *awesomeToiler) Toil() {
		//@TODO: Do work here.
	}
	
	func (toiler *awesomeToiler) RecoveredNotice() {
		//@TODO: Do something with this notification.
	}

And of course, a toiler can take advantage of both of these notifications and have
both a ReturnedNotice(), PanickedNotice() and RecoveredNotice() method. For example:

	type awesomeToiler struct{}
	
	func newAwesomeToiler() {
	
		toiler := awesomeToiler{}
	
		return &toiler
	}
	
	func (toiler *awesomeToiler) Toil() {
		//@TODO: Do work here.
	}
	
	func (toiler *awesomeToiler) ReturnedNotice() {
		//@TODO: Do something with this notification.
	}
	
	func (toiler *awesomeToiler) PanickedNotice() {
		//@TODO: Do something with this notification.
	}
	
	func (toiler *awesomeToiler) RecoveredNotice() {
		//@TODO: Do something with this notification.
	}

*/
package retoil
